package main

import (
	"fmt"
	"net"
)

func main() {
	// Connect to the TCP server
	conn, err := net.Dial("tcp", "localhost:6969")
	if err != nil {
		fmt.Printf("Error connecting to server: %v\n", err)
		return
	}
	defer conn.Close() // Always close the connection when done

	// Send data to the server
	// message := "*3\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"
	message := "PING"
	// message := "*2\r\n+OK\r\n+doke\r\n"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Printf("Error sending data: %v\n", err)
		return
	}

	// Read response from server
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	fmt.Printf("Server response: %s\n", string(buffer[:n]))
}
