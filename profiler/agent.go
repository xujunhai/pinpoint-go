package profiler

import (
	"pinpoint-go/bootstrap"
	"pinpoint-go/bootstrap/config"
)

type Agent struct {
	profilerConfig     *config.ProfilerConfig    //配置项
	//applicationContext module.ApplicationContext //应用监控上下文
	agentStatus        AgentStatus               //agent状态
}

func NewAgent(agentOption bootstrap.AgentOption) *Agent {
	return &Agent{
		profilerConfig:     agentOption.ProfilerConfig,
		//applicationContext: module.NewApplicationContext(agentOption.ProfilerConfig),
		agentStatus:        INITIALIZING,
	}
}

func (a *Agent) ChangeStatus(status AgentStatus) {
	a.agentStatus = status
}

func (a *Agent) Start() {
	if a.agentStatus == INITIALIZING {
		a.ChangeStatus(RUNNING)
	}
	//a.applicationContext.Start()
}

func (a *Agent) Stop() {
	if a.agentStatus == RUNNING {
		a.ChangeStatus(STOPPED)
	}
	//a.applicationContext.Close()
}
