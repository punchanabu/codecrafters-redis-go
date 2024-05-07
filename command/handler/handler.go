package handler

import "strings"

func HandleCommand(command string, argument []string) string {

	lowerCommand := strings.ToLower(command)

	switch lowerCommand {
	case "ping":
		return handlePing()
	case "echo":
		return handleEcho(argument)
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
