package config

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
        Role: role,
    }
}