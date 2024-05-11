package helper

import "strings"

func IsWriteCommand(command string) bool {
	lowerCommand := strings.ToLower(command)
	switch lowerCommand {
	case "set", "del", "append", "incr", "decr":
		return true
	default:
		return false
	}
}
