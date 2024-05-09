package main

import (
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/redis-starter-go/config"
	"github.com/codecrafters-io/redis-starter-go/connection"
	"github.com/codecrafters-io/redis-starter-go/store"
)

func main() {

	var portFlag = "6379" // default port
	var replicaOf = ""
	var role = "master"

	// parse command-line arguments
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--port":
			if i+1 < len(os.Args) {
				portFlag = os.Args[i+1]
				i++ // Skip next argument since it's the value for --port
			}
		case "--replicaof":
			if i+2 < len(os.Args) {
				replicaOf = os.Args[i+1] + ":" + os.Args[i+2]
				role = "slave"
				i += 2 // Skip next two arguments as they are host and port
			}
		}
	}

	// Initialize the Replica Config with the determined role
	replicaConfig := config.NewReplicaConfig(role)

	// Temporarily print the replica information if the role is slave
	if role == "slave" {
		connection.InitHandshake(replicaOf)
	}

	// Listen on all interfaces on port 6379
	listener, err := net.Listen("tcp", "0.0.0.0:"+portFlag)
	if err != nil {
		fmt.Println("Failed to bind to port 6379:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server listening on port " + portFlag)

	// Initialize the Store for Get and Set commands
	redisStore := store.New()
	// Start the cleanup routine for store.
	redisStore.CleanUpExpiredKey()

	// Accept connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		// Handle net connection in a new goroutine
		go connection.HandleConnection(conn, redisStore, replicaConfig)
	}
}
