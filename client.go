package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    // Connect using TCP and to server IP/domain
    connection, err := net.Dial("tcp", "127.0.0.1:6969")

    // Handle error(s)
    if err != nil {
        fmt.Printf("Error: %s\n", err.Error())
        return
    }

    fmt.Printf("Client has connected!\n")
    // Read input
    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("Username: ")

    username, err := reader.ReadString('\n')

    // Handle error(s)
    if err != nil {
        fmt.Printf("Error: %s\n", err.Error())
        return
    }
    // Send message
    fmt.Fprintf(connection, fmt.Sprintf("%s has joined the chat", username))

    for {
        // Read recieved bytes as a string
        output, err := bufio.NewReader(connection).ReadString('\n')

        // Handle error(s)
        if err != nil {
            fmt.Printf("Error: %s\n", err.Error())
            return
        }

        // Print recieved bytes
        fmt.Printf("%s", string(output))

        // Read input
        reader := bufio.NewReader(os.Stdin)
        fmt.Printf("Send message: ")

        input, err := reader.ReadString('\n')

        // Handle error(s)
        if err != nil {
            fmt.Printf("Error: %s\n", err.Error())
            return
        }

        // Send message
        fmt.Fprintf(connection, input+"\n")
    }
}
