package connection

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/redis-starter-go/command/handler"
	"github.com/codecrafters-io/redis-starter-go/command/helper"
	"github.com/codecrafters-io/redis-starter-go/command/parser"
	"github.com/codecrafters-io/redis-starter-go/config"
	"github.com/codecrafters-io/redis-starter-go/store"
)

func HandleConnection(conn net.Conn, store *store.Store, replicaConfig *config.ReplicaConfig) {
	defer conn.Close()
	fmt.Println("Connection from ", conn.RemoteAddr().String())

	/*
		Here I will be keeping track of slave replica connecting to the master
		So that I can propagate commands with those replicas
	*/
	var replica *ReplicaConnection
	if replicaConfig.Role == "slave" {
		replica = addReplica(conn)
		defer removeReplica(replica.SessionID) // For cleaning up the Replicas
	}

	for {
		// Read data from the connection
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Failed to read data: ", err.Error())
			break
		}

		// Handle PING check
		if string(buffer[:n]) == "+PING\r\n" {
			handler.HandleCommand("PING", nil, store, replicaConfig)
			break
		}

		// Decode the data
		command, argument, err := parser.Decode(buffer[:n])
		if err != nil {
			fmt.Println("Failed to decode data: ", err.Error())
			break
		}

		// Handle the command
		response := handler.HandleCommand(command, argument, store, replicaConfig)

		// Encode the response
		encodedResponse := parser.Encode(response)

		// Write the response
		conn.Write([]byte(encodedResponse))

		/*
			Check if we have to send an RDB file as well
			The RDB file will be send after PSYNC command
		*/
		if command == "PSYNC" {
			config.SendRDBFile(conn)
		}

		// Propagate the command to the slave replicas
		if helper.IsWriteCommand(command) {
			PropagateCommandToReplica(buffer[:n])
		}
	}
}
