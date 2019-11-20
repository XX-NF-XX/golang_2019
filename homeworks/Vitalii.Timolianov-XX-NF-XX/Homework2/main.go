package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const sourceURL string = "https://gist.githubusercontent.com/XX-NF-XX/267c2b85144d4de18c5b917414557230/raw/38c96e44b4ba453f2e7052be908ca6a5a29183a0/sample.json"
const sourceFilePath string = "./sample.json"

// NameReader -
// source can be a local file or a url to JSON file
// read this file, parse and return "key" field value
type NameReader interface {
	Read(source string, key string) string
}

// JSON - json representation
type JSON map[string]interface{}

// Storages - map of JSON storages
type Storages map[string]JSON

func (s Storages) checkError(err error) {
	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	}
}

func (s Storages) unmarshalJSON(data []byte) JSON {
	var jsonData JSON
	err := json.Unmarshal(data, &jsonData)
	s.checkError(err)

	return jsonData
}

func (s Storages) createURLStorage(url string) {
	response, httpErr := http.Get(url)
	s.checkError(httpErr)

	defer response.Body.Close()
	data, readErr := ioutil.ReadAll(response.Body)
	s.checkError(readErr)

	s[url] = s.unmarshalJSON(data)
}

func (s Storages) createFileStorage(path string) {
	data, err := ioutil.ReadFile(sourceFilePath)
	s.checkError(err)

	s[path] = s.unmarshalJSON(data)
}

func (s Storages) createStorage(source string) {
	_, ok := s[source]
	if ok {
		return
	}

	if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
		s.createURLStorage(source)
	} else {
		s.createFileStorage(source)
	}
}

func (s Storages) Read(source string, key string) string {
	s.createStorage(source)

	jsonData := s[source]
	switch value := jsonData[key].(type) {
	case string:
		return value
	default:
		fmt.Println("Sorry! Can read only strings from JSON!")
	}

	return ""
}

func main() {
	storages := make(Storages)

	name := storages.Read(sourceURL, "name")
	fmt.Println("name from URL:", name)

	name = storages.Read(sourceFilePath, "name")
	fmt.Println("name from file:", name)
}
