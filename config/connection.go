package config

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/redis-starter-go/command/parser"
	"github.com/codecrafters-io/redis-starter-go/command/handler"
)

func HandleConnection(conn net.Conn) {
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

		// Decode the data
		command, argument := parser.Decode(buffer[:n])

		// Handle the command
		response := handler.HandleCommand(command, argument)

		// Encode the response
		encodedResponse := parser.Encode(response)

		// Write the response
		conn.Write([]byte(encodedResponse))
	}
}
