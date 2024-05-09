package connection

import (
	"fmt"
	"net"
)

func InitHandshake(masterAddress string) {
	conn, err := net.Dial("tcp", masterAddress)
	if err != nil {
		fmt.Println("Failed to connect to master: ", err.Error())
		return
	}
	defer conn.Close()

	err = sendPing(conn)
	if err != nil {
		fmt.Println("Failed to send PING to master: ", err.Error())
		return
	}
}

func sendPing(conn net.Conn) error {
	_, err := conn.Write([]byte("*1\r\n$4\r\nPING\r\n"))
	if err != nil {
		return err
	}
	return nil
}
