package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Listen for connection
	listener, err := net.Listen("tcp", "127.0.0.1:6969")
	defer listener.Close()
	fmt.Printf("Server is listening!\n")

	// Handle error(s)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	for {
		connection, err := listener.Accept()

		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			return
		}

		// Send message
		connection.Write([]byte("You have joined the chat!\n"))

		// Read recieved bytes as string
		username, err := bufio.NewReader(connection).ReadString('\n')

		// Handle error(s)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			return
		}

		// Print recieved bytes
		fmt.Printf("%s has joined the chat\n", string(username[:len(username)-1]))

		// Handle connection
		go handle(username, connection)
	}
}

func handle(username string, connection net.Conn) {
	for {
		// Read recieved bytes as string
		output, err := bufio.NewReader(connection).ReadString('\n')

		// Handle error(s)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			return
		}

		// Print recieved bytes
		fmt.Printf("%s: %s", string(username[:len(username)-1]), output)

		// Send message back
		connection.Write([]byte(username[:len(username)-1] + ": " + output))
	}
}
