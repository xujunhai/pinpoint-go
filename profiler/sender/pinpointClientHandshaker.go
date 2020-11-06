package sender

import (
	"log"
	"os"
	"pinpoint-go/common"
	"pinpoint-go/config"
	"pinpoint-go/rpc/control"
	"pinpoint-go/rpc/packet"
	"time"
)

const (
	HandshakeStateInit     byte = 0
	HandshakeStateStarted  byte = 1
	HandshakeStateFinished byte = 2
	// STATE_INIT -> STATE_STARTED -> STATE_COMPLETED
	// STATE_INIT -> STATE_STARTED -> STATE_ABORTED
	HandshakeMaxCount = 10
	HandshakeInterval = 60 * time.Second
)

type HandShake struct {
	tcpDataSender     *TcpDataSender
	status            byte
	handshakeCount    int
	handshakeResponse *packet.HandshakeResponseCode
	handshakePacket   packet.Packet
	sendTicker        *time.Ticker
}

func NewHandShake(tcpDataSender *TcpDataSender) *HandShake {
	return &HandShake{
		tcpDataSender:  tcpDataSender,
		status:         HandshakeStateInit,
		handshakeCount: 0,
	}
}

func (p *HandShake) createHandShakePacket() packet.Packet {
	hostName, err := os.Hostname()
	if err != nil {
		log.Println("get hostname error ", err)
		hostName = "Unknown"
	}

	var handshakeData = make(map[string]interface{})
	handshakeData["socketId"] = p.tcpDataSender.GetSocketID()
	handshakeData["hostName"] = hostName
	handshakeData["supportServer"] = true
	handshakeData["ip"] = "0.0.0.0"
	handshakeData["agentId"] = config.AgentId()
	handshakeData["applicationName"] = config.AppName()
	handshakeData["serviceType"] = config.ServerType()
	handshakeData["pid"] = int32(os.Getegid())
	handshakeData["version"] = common.PinpointAgentVersion
	handshakeData["startTimestamp"] = config.StartTime()

	encoder := control.NewEncoder()
	payload := encoder.Encode(handshakeData)

	pack := packet.NewControlHandshakePacket()
	pack.RequestID = 0
	pack.Payload = payload

	return pack
}

func (p *HandShake) Start(quitC <-chan struct{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("handshake panic: ", err)
		}
	}()

	p.status = HandshakeStateStarted
	p.handshakePacket = p.createHandShakePacket()

	p.sendTicker = time.NewTicker(HandshakeInterval)
	defer p.sendTicker.Stop()

	p.tcpDataSender.SendPacket(p.handshakePacket)
	p.handshakeCount++

	for {
		select {
		case <-p.sendTicker.C:
			if p.handshakeCount == HandshakeMaxCount {
				p.status = HandshakeStateFinished
				return
			}

			if p.status == HandshakeStateStarted {
				p.tcpDataSender.SendPacket(p.handshakePacket)
			}
			p.handshakeCount++

		case <-quitC:
			return
		}
	}
}

func (p *HandShake) HandlerResponse(pack *packet.ControlHandshakeResponsePacket) {
	decoder := control.NewDecoder(pack.GetPayload())
	responseData := decoder.Decode()

	p.status = HandshakeStateFinished

	var code int32 = -1
	var subCode int32 = -1

	msgData, ok := responseData.(map[string]interface{})
	if !ok {
		p.handshakeResponse = packet.HandshakeUnknownCode
		return
	}

	vcode, ok := msgData[packet.HsCode]
	if !ok {
		p.handshakeResponse = packet.HandshakeUnknownCode
		return
	}

	code, ok = vcode.(int32)
	if !ok {
		p.handshakeResponse = packet.HandshakeUnknownCode
		return
	}

	vcode, ok = msgData[packet.HsSubCode]
	if !ok {
		p.handshakeResponse = packet.HandshakeUnknownCode
		return
	}

	subCode, ok = vcode.(int32)
	if !ok {
		p.handshakeResponse = packet.HandshakeUnknownCode
		return
	}

	p.handshakeResponse = packet.GetHandShakeCode(code, subCode)

	if p.handshakeResponse == packet.HandshakeSuccess || p.handshakeResponse == packet.HandshakeAlreadyKnown {
		p.tcpDataSender.SockState.ToRunSimplex()
	} else if p.handshakeResponse == packet.HandshakeDuplexCommunication ||
		p.handshakeResponse == packet.HandshakeAlreadyDuplexCommunication {
		p.tcpDataSender.SockState.ToRunDuplex()
	} else if p.handshakeResponse == packet.HandshakeSimplexCommunication ||
		p.handshakeResponse == packet.HandshakeAlreadySimplexCommunication {
		p.tcpDataSender.SockState.ToRunSimplex()
	}

	p.stopSendTicker()
}

func (p *HandShake) stopSendTicker() {
	if p.sendTicker != nil {
		p.sendTicker.Stop()
	}
}
