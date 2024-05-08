package handler

import (
	"strings"

	"github.com/codecrafters-io/redis-starter-go/store"
)

func HandleCommand(command string, argument []string, store *store.Store) string {

	lowerCommand := strings.ToLower(command)

	switch lowerCommand {
	case "ping":
		return handlePing()
	case "echo":
		return handleEcho(argument)
	case "get":
		return handleGet(argument, store)
	case "set":
		return handleSet(argument, store)
	default:
		return "-ERR unknown command"
	}
}

func handlePing() string {
	return "+PONG"
}

func handleEcho(argument []string) string {
	if len(argument) == 0 {
		return "-ERR no argument provided"
	}
	return "+" + argument[0]
}

func handleGet(argument []string, store *store.Store) string {
	if len(argument) == 0 {
		return "-ERR no argument provided"
	}
	value, ok := store.Get(argument[0])
	/*
		If there is no value returns an empty string
		as it will be encoded as a $-1 response.
	*/
	if !ok {
		return ""
	}
	return "+" + value
}

func handleSet(argument []string, store *store.Store) string {
	if len(argument) < 2 {
		return "-ERR not enough arguments"
	}
	store.Set(argument[0], argument[1])
	return "+OK"
}
