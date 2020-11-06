package metadata

type AgentInfoFactory struct {
	agentInformation *AgentInformation
}

func (af *AgentInfoFactory) AgentInformation() *AgentInformation {
	return af.agentInformation
}

func (af *AgentInfoFactory) SetAgentInformation(agentInformation *AgentInformation) {
	af.agentInformation = agentInformation
}

func NewAgentInfoFactory(agentInformation *AgentInformation) *AgentInfoFactory {
	return &AgentInfoFactory{agentInformation: agentInformation}
}

func (af AgentInfoFactory) CreateAgentInfo() AgentInfo {
	return AgentInfo{}
}
