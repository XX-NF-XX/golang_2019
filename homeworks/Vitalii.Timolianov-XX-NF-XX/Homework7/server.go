package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func readFromConnection(conn net.Conn) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read message: %v", err)
		}

		fmt.Printf("<<< %v", string(message))

		reader := bufio.NewReader(os.Stdin)
		fmt.Printf(">>> ")

		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to from STDIN: %v", err)
		}

		fmt.Fprintf(conn, text)
	}
}

func acceptConnection(listener net.Listener) {
	log.Printf("Waiting for a connection...\n")

	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf("Failed to accept connection: %v", err)
	}

	log.Printf("Connection acquired from %v", conn.RemoteAddr())

	readFromConnection(conn)
}

func runServer(address string) error {
	log.Printf("Start server at %v...\n", address)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	log.Print("Server started.\n")

	acceptConnection(listener)
	return nil
}
