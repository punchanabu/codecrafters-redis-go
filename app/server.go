package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/redis-starter-go/config"
	"github.com/codecrafters-io/redis-starter-go/connection"
	"github.com/codecrafters-io/redis-starter-go/store"
)

func main() {

	// Configure listening to port
	var portFlag string
	var replicaOf string // To hold the replica of flag data if any

	flag.StringVar(&portFlag, "port", "6379", "The port on which the Redis service will bind")
	flag.StringVar(&replicaOf, "replicaof", "", "Specify master host and port if this instance is a replica")
	flag.Parse()

	// Determine role based on the replicaOf flag
	role := "master"
	if replicaOf != "" {
		role = "slave"
	}

	// Initialize the Replica Config with the determined role
	replicaConfig := config.NewReplicaConfig(role)

	// Listen on all interfaces on port 6379
	listener, err := net.Listen("tcp", "0.0.0.0:"+portFlag)
	if err != nil {
		fmt.Println("Failed to bind to port 6379:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server listening on port 6379")

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
