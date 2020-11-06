package monitor

import (
	"pinpoint-go/config"
	"pinpoint-go/profiler/sender"
	"pinpoint-go/proto/thrift"
	"pinpoint-go/proto/thrift/gen-go/pinpoint"
	"time"
)

const (
	AgentStatBatchNum = 6
)

type AgentStatMonitor struct {
	client    *sender.UdpDataSender
	agentStat []*pinpoint.TAgentStat
}

func NewAgentStatMonitor(addr string) *AgentStatMonitor {
	return &AgentStatMonitor{
		client:    sender.NewUdpDataSender(addr),
		agentStat: make([]*pinpoint.TAgentStat, 0, AgentStatBatchNum),
	}
}

func (p *AgentStatMonitor) Start() {
	quitC := make(chan struct{})
	defer func() {
		close(quitC)
	}()

	go p.collectAgentStat(quitC)
	p.client.Start()
}

func (p *AgentStatMonitor) collectAgentStat(quitC <-chan struct{}) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			agentStat := pinpoint.NewTAgentStat()
			startTime := config.StartTime()
			agentId := config.AgentId()
			agentStat.AgentId = &agentId
			agentStat.StartTimestamp = &startTime

			p.agentStat = append(p.agentStat, agentStat)

			if len(p.agentStat) == AgentStatBatchNum {
				agentStatBatch := pinpoint.NewTAgentStatBatch()
				agentStatBatch.AgentId = config.AgentId()
				agentStatBatch.StartTimestamp = config.StartTime()
				agentStatBatch.AgentStats = p.agentStat[0:AgentStatBatchNum]

				payload := thrift.Serialize(agentStatBatch)
				p.client.SendPacket(payload)
				p.agentStat = p.agentStat[:0]
			}

		case <-quitC:
			return
		}
	}
}
