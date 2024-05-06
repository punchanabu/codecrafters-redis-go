package main

import (
	"fmt"

	"net"
	"os"
)

func main() {

	// Listen on all interfaces on port 6379
	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server listening on port 6379")

	// Accept connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		
		// Handle net connection in a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Connection from ", conn.RemoteAddr().String())

	for {
		// Read data from the connection
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Failed to read data: ", err.Error())
			break
		}
		
		fmt.Println("Received: ", string(buffer[:n]))

		// Write response to the connection
		response := "+PONG\r\n"
		conn.Write([]byte(response))
		
	}
}
