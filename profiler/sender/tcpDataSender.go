package sender

import (
	"bufio"
	"encoding/binary"
	"io"
	"log"
	"net"
	"pinpoint-go/config"
	"pinpoint-go/rpc/common"
	"pinpoint-go/rpc/packet"
	"sync/atomic"
	"time"
)

//tcp udp相关定义
const (
	TcpReadBufferLen  int = 1024 * 64
	TcpWriteBufferLen int = 1024 * 64
	TcpPacketCount        = 1024 * 10
	PingInterval          = 5 * time.Minute
)

type TcpDataSender struct {
	conn      net.Conn
	SockState *common.SocketState
	sendChan  chan packet.Packet
	pingID    int32
	socketID  int32
	requestID int
	handShake *HandShake
	isRestart bool
}

func NewTcpDataSender() *TcpDataSender {
	return &TcpDataSender{isRestart: true}
}

func (p *TcpDataSender) Init(quitC <-chan struct{}) {
	var err error
	p.handShake = NewHandShake(p)
	p.SockState = common.NewSocketState()
	p.sendChan = make(chan packet.Packet, TcpPacketCount)
	atomic.AddInt32(&p.socketID, 1)

	defer func() {
		if err := recover(); err != nil {
			//先用系统的日志
			log.Println("tcpclient init recover: ", err)
		}

		if p.isRestart {
			time.AfterFunc(5*time.Second, func() {
				go p.Init(quitC)
			})
		}
	}()

	defer func() {
		log.Println("its closed???")
		if p.conn != nil {
			_ = p.conn.Close()
		}
	}()

	for {
		p.conn, err = net.Dial("tcp", config.InfoAddr())
		if err != nil {
			log.Println("dial ", config.InfoAddr(), " error：", err)
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}

	tcpConn := p.conn.(*net.TCPConn)
	err = tcpConn.SetReadBuffer(TcpReadBufferLen)
	if err != nil {
		log.Println("set read buffer error: ", err)
		return
	}

	err = tcpConn.SetWriteBuffer(TcpWriteBufferLen)
	if err != nil {
		log.Println("set write buffer error: ", err)
		return
	}

	p.SockState.ToConnected()
	log.Println("connect success ", config.InfoAddr())
	p.SockState.ToRunWithoutHandShake()

	go p.writePacket(quitC) //从channel读取数据流
	go p.doPing(quitC)      //定时任务执行ping
	go p.handShake.Start(quitC)

	p.handleRead()
}

func (p *TcpDataSender) GetSocketID() int32 {
	return atomic.LoadInt32(&p.socketID)
}

func (p *TcpDataSender) SendPacket(pack packet.Packet) {
	if p.sendChan == nil {
		return
	}
	p.sendChan <- pack
}

func (p *TcpDataSender) writePacket(quitC <-chan struct{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("write packet error: ", err)
		}
	}()

	for {
		select {
		case pack, ok := <-p.sendChan:
			if !ok {
				break
			}

			if pack.GetPacketType() == packet.ApplicationRequest {
				if v, ok := pack.(*packet.RequestPacket); ok {
					v.RequestID = p.requestID
					p.requestID++
				}
			}

			payload, err := pack.Encode()
			if err != nil {
				break
			}
			_, err = p.conn.Write(payload)
			if err != nil {
				log.Println("error: ", err)
			}

		case <-quitC:
			return
		}
	}
}

func (p *TcpDataSender) doPing(quitC <-chan struct{}) {
	firstTick := time.After(300 * time.Millisecond)
	tick := time.NewTicker(PingInterval)
	defer tick.Stop()

	ping := func() {
		pingPacket := packet.NewPingPayloadPacket()
		pingPacket.StateCode = p.SockState.GetCurrentState()
		pingPacket.StateVersion = 0
		pingPacket.PingID = int(atomic.AddInt32(&p.pingID, 1))
		p.SendPacket(pingPacket)
	}

	for {
		select {
		case <-firstTick:
			ping()
		case <-tick.C:
			ping()

		case <-quitC:
			return
		}
	}
}

func (p *TcpDataSender) handleRead() {
	var err error
	reader := bufio.NewReaderSize(p.conn, TcpReadBufferLen)
	typeBuf := make([]byte, 2)
	//
	for {
		if _, err = io.ReadFull(reader, typeBuf); err != nil {
			log.Println("io.ReadFull", err)
			return
		}
		packetType := int16(binary.BigEndian.Uint16(typeBuf[0:2]))

		pack, err := packet.DecodePacketFactory(packetType, p.conn, reader)

		if err != nil {
			return
		}

		switch packetType {
		case packet.ControlHandshakeResponse:
			p.handShake.HandlerResponse(pack.(*packet.ControlHandshakeResponsePacket))

		case packet.ControlPingPayload:
			fallthrough
		case packet.ControlPingSimple:
			fallthrough
		case packet.ControlPing:
			pack := packet.NewPongPacket()
			p.SendPacket(pack)
		}
	}
}

func (p *TcpDataSender) Close() {
	p.isRestart = false

	if p.conn != nil {
		_ = p.conn.Close()
	}

	if p.SockState != nil {
		p.SockState.ToClosedByClient()
	}
}
