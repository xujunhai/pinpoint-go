package profiler

import (
	"log"
	"os"
	"pinpoint-go/common"
	"pinpoint-go/common/util"
	"pinpoint-go/config"
	"pinpoint-go/profiler/metadata"
	"pinpoint-go/profiler/sender"
	"pinpoint-go/proto/thrift"
	"pinpoint-go/proto/thrift/gen-go/pinpoint"
	"pinpoint-go/rpc/packet"
	"time"
)

const (
	DefaultAgentInfoRefreshIntervalMs = 24 * 60 * 60 * 1000
	DefaultAgentInfoSendIntervalMs    = 3 * 1000
	DefaultMaxTryCountPerAttempt      = 3
)

type AgentInfoSender struct {
	AgentInfoFactory  metadata.AgentInfoFactory
	RefreshIntervalMs int64
	SendIntervalMs    int64
	MaxTryPerAttempt  int
	dataSender        sender.DataSender

	//TODO refactor
	TcpDataSender *sender.TcpDataSender
}

func NewAgentInfoSender(tcpDataSender *sender.TcpDataSender /*agentInfoFactory util.AgentInfoFactory,dataSender sender.DataSender*/) *AgentInfoSender {
	return &AgentInfoSender{
		//AgentInfoFactory:agentInfoFactory,
		//dataSender: dataSender,
		RefreshIntervalMs: DefaultAgentInfoRefreshIntervalMs,
		SendIntervalMs:    DefaultAgentInfoSendIntervalMs,
		MaxTryPerAttempt:  DefaultMaxTryCountPerAttempt,
		TcpDataSender:     tcpDataSender,
	}
}

func (p *AgentInfoSender) Start() {
	//do nothing 创建task任务 定时发送
}

func (p *AgentInfoSender) SendAgentInfo(quitC <-chan struct{}) {
	//agentInfo := p.AgentInfoFactory.CreateAgentInfo()
	agentInfo := pinpoint.NewTAgentInfo()
	hostName, err := os.Hostname()
	if err != nil {
		log.Println("get hostname error ", err)
		hostName = "Unknown"
	}

	agentInfo.Hostname = hostName
	agentInfo.IP = util.GetHostIP()
	agentInfo.Ports = ""
	agentInfo.AgentId = config.AgentId()
	agentInfo.ApplicationName = config.AppName()
	agentInfo.ServiceType = int16(config.ServerType())
	agentInfo.Pid = int32(os.Getpid())
	agentInfo.AgentVersion = common.PinpointAgentVersion
	agentInfo.VmVersion = ""
	agentInfo.StartTimestamp = config.StartTime()

	payload := thrift.Serialize(agentInfo)

	request := packet.NewRequestPacket()
	request.Payload = payload

	ticker := time.NewTicker(DefaultAgentInfoRefreshIntervalMs)
	defer ticker.Stop()

	firstTicker := time.After(1 * time.Second)

	for {
		select {
		case <-firstTicker:
			p.TcpDataSender.SendPacket(request)
		case <-ticker.C:
			p.TcpDataSender.SendPacket(request)
		case <-quitC:
			return
		}
	}

}
