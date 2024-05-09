package connection

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func InitHandshake(masterHost string, masterPort string, slavePort string) {

	// Initiate the connection with the master
	conn, err := net.Dial("tcp", masterHost+":"+masterPort)
	if err != nil {
		fmt.Println("Failed to connect to master: ", err.Error())
		return
	}
	defer conn.Close()

	ok, err := sendPing(conn)
	if err != nil {
		fmt.Println("Failed to send PING to master: ", err.Error())
		return
	}

	if !ok {
		fmt.Println("Master did not respond with PONG")
		return
	}

	if ok, err := sendReplConfig(conn, "listening-port", slavePort); !ok {
		fmt.Println("Failed to send REPLCONF for listening port:", err)
		return
	}

	if ok, err := sendReplConfig(conn, "capa", "psync2"); !ok {
		fmt.Println("Failed to send REPLCONF for capabilities:", err)
		return
	}
}

func sendPing(conn net.Conn) (bool, error) {
	// Send the PING command to the Redis server
	_, err := conn.Write([]byte("*1\r\n$4\r\nPING\r\n"))
	if err != nil {
		return false, fmt.Errorf("failed to send PING: %w", err)
	}

	// Prepare to read the response using bufio.Reader
	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false, fmt.Errorf("failed to read response: %w", err)
	}

	// Trim the response and check if it's "+PONG"
	response = strings.TrimSpace(response)
	if response == "+PONG" {
		fmt.Println("Received:", response)
		return true, nil // Successfully received the correct response
	} else {
		return false, fmt.Errorf("unexpected response: %s", response)
	}
}

func sendReplConfig(conn net.Conn, command, value string) (bool, error) {
	message := fmt.Sprintf("*3\r\n$8\r\nREPLCONF\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n", len(command), command, len(value), value)
	_, err := conn.Write([]byte(message))
	if err != nil {
		return false, fmt.Errorf("failed to send REPLCONF: %w", err)
	}

	// Prepare to read the response using bufio.Reader
	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false, fmt.Errorf("failed to read response: %w", err)
	}

	// Trim the response and check if it's "+OK"
	response = strings.TrimSpace(response)
	if response == "+OK" {
		fmt.Println("Received:", response)
		return true, nil // Successfully received the correct response
	} else {
		return false, fmt.Errorf("unexpected response: %s", response)
	}
}
