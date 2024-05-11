package connection

import (
	"log"
	"net"

	"github.com/google/uuid"
)

type ReplicaConnection struct {
	Conn      net.Conn
	SessionID string
}

var replicas []*ReplicaConnection

// For adding new replicas (registeration)
func addReplica(conn net.Conn) *ReplicaConnection {
	sessionID := uuid.New().String()
	replica := &ReplicaConnection{
		Conn:      conn,
		SessionID: sessionID,
	}
	replicas = append(replicas, replica)
	log.Println("Added new replica:", sessionID)
	return replica
}

func removeReplica(sessionID string) {
	for i, replica := range replicas {
		if replica.SessionID == sessionID {
			replicas = append(replicas[:i], replicas[i+1:]...)
			log.Println("Removed replica with session ID:", sessionID)
			break
		}
	}
}

/*
I think we dont need to do unneccessary operation just send the raw bytes
and let the replicas process it instead ?
*/
func PropagateCommandToReplica(command []byte) {
	for _, replica := range replicas {
		_, err := replica.Conn.Write(command)
		if err != nil {
			log.Println("Failed to Propagate command to replica:", err)
		}
	}
}
