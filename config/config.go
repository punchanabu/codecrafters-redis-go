package config

import (
	"encoding/base64"
	"fmt"
	"net"
	"strings"

	"github.com/google/uuid"
)

type ReplicaConfig struct {
	Role                       string
	ConnectedSlaves            int
	MasterReplID               string
	MasterReplOffset           int
	SecondReplOffset           int
	ReplBacklogActive          int
	ReplBacklogSize            int
	ReplBacklogFirstByteOffset int
	ReplBacklogHistLen         int
}

func NewReplicaConfig(role string) *ReplicaConfig {
	return &ReplicaConfig{
		Role:             role,
		MasterReplID:     strings.ReplaceAll(uuid.New().String(), "-", ""),
		MasterReplOffset: 0,
	}
}

/*
RDB files will contains the current state of the replica
For now, I will hardcoded the RDB file in the function
*/
func SendRDBFile(conn net.Conn) {
	// Example base64 string (you would use your actual string)
	base64RDB := "UkVESVMwMDEx+glyZWRpcy12ZXIFNy4yLjD6CnJlZGlzLWJpdHPAQPoFY3RpbWXCbQi8ZfoIdXNlZC1tZW3CsMQQAPoIYW9mLWJhc2XAAP/wbjv+wP9aog=="

	// Decode the base64 string to binary
	rdbContents, err := base64.StdEncoding.DecodeString(base64RDB)
	if err != nil {
		fmt.Println("Failed to decode RDB content: ", err)
		return
	}

	// Create the header indicating the length of the RDB content
	rdbLength := len(rdbContents)
	rdbHeader := fmt.Sprintf("$%d\r\n", rdbLength)

	// Ensure that the header and the content are sent in one go to avoid misinterpretation
	fullMessage := rdbHeader + string(rdbContents)
	conn.Write([]byte(fullMessage))
}
