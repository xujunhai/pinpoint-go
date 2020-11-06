package sender

import (
	"pinpoint-go/proto/thrift"
	"pinpoint-go/proto/thrift/gen-go/trace"
)

type AgentSpanClient struct {
	client *UdpDataSender
}

func NewAgentSpanClient(addr string) *AgentSpanClient {
	return &AgentSpanClient{
		client: NewUdpDataSender(addr),
	}
}

func (p *AgentSpanClient) Start() {
	p.client.Start()
}

func (p *AgentSpanClient) SendTrace(tspan *trace.TSpan) {
	payload := thrift.Serialize(tspan)
	p.client.SendPacket(payload)
}
