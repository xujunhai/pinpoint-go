package profiler

type AgentStatus uint8

const (
	INITIALIZING AgentStatus = iota
	RUNNING
	STOPPED
)
