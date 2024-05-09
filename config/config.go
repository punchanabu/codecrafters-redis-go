package config

import (
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
