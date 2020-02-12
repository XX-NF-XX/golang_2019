package main

import (
	"log"
)

func main() {
	loadConfig()
	address := config.getAddress()

	showUI()

	if config.ServerIP == defaultServerIP {
		if err := runServer(address); err != nil {
			log.Fatalf("%v", err)
		}
		return
	}

	if err := connectToServer(address); err != nil {
		log.Fatalf("%v", err)
	}
}
