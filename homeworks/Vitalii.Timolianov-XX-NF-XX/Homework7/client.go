package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func connectToServer(address string) error {
	log.Printf("Connect to server at %v...\n", address)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Printf("Cannot connect to server %v", address)
		return err
	}

	log.Printf("Connected to server successfully.\n")

	for {
		reader := bufio.NewReader(os.Stdin)
		log.Print("Text to send: ")

		text, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		fmt.Fprintf(conn, text)

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return err
		}

		log.Print("Message from server: " + message)
	}
}
