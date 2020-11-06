package metadata

type AgentInformation struct {
	agentId         string
	applicationName string
	isContainer     bool
	startTime       int64
	pid             int
	machineName     string
	hostIp          string
	serverType      int16
	agentVersion    string
}

func (a *AgentInformation) AgentVersion() string {
	return a.agentVersion
}

func (a *AgentInformation) SetAgentVersion(agentVersion string) {
	a.agentVersion = agentVersion
}

func (a *AgentInformation) HostIp() string {
	return a.hostIp
}

func (a *AgentInformation) SetHostIp(hostIp string) {
	a.hostIp = hostIp
}

func (a *AgentInformation) MachineName() string {
	return a.machineName
}

func (a *AgentInformation) SetMachineName(machineName string) {
	a.machineName = machineName
}

func (a *AgentInformation) Pid() int {
	return a.pid
}

func (a *AgentInformation) SetPid(pid int) {
	a.pid = pid
}

func (a *AgentInformation) StartTime() int64 {
	return a.startTime
}

func (a *AgentInformation) SetStartTime(startTime int64) {
	a.startTime = startTime
}

func (a *AgentInformation) ApplicationName() string {
	return a.applicationName
}

func (a *AgentInformation) SetApplicationName(applicationName string) {
	a.applicationName = applicationName
}

func (a *AgentInformation) AgentId() string {
	return a.agentId
}

func (a *AgentInformation) SetAgentId(agentId string) {
	a.agentId = agentId
}

func NewAgentInformation(agentId, applicationName string, isContainer bool, startTime int64, pid int, machineName string, hostIp string, agentVersion string) *AgentInformation {
	return &AgentInformation{
		agentId:         agentId,
		applicationName: applicationName,
		isContainer:     isContainer,
		startTime:       startTime,
		pid:             pid,
		machineName:     machineName,
		hostIp:          hostIp,
		agentVersion:    agentVersion,
	}
}
