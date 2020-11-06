package bootstrap

import "pinpoint-go/bootstrap/config"

type AgentOption struct {
	agentId         string
	applicationName string
	ProfilerConfig  *config.ProfilerConfig
}

func NewAgentOption(agentId, applicationName string, config *config.ProfilerConfig) AgentOption {
	return AgentOption{
		agentId:         agentId,
		applicationName: applicationName,
		ProfilerConfig:  config,
	}
}
