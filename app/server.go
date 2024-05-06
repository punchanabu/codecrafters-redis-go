package main

import (
	"fmt"
	"net"
	"os"
	"github.com/codecrafters-io/redis-starter-go/config"
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
		go config.HandleConnection(conn)
	}
}
