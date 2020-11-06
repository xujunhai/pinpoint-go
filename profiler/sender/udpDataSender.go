package sender

import (
	"log"
	"net"
	"time"
)

const (
	UdpPacketCount = 1024
)

type UdpDataSender struct {
	conn      net.Conn
	addr      string
	sendChan  chan []byte
	isRestart bool
}

func NewUdpDataSender(addr string) *UdpDataSender {
	return &UdpDataSender{addr: addr, isRestart: true}
}

func (p *UdpDataSender) Start() {
	var err error

	p.sendChan = make(chan []byte, UdpPacketCount)

	defer func() {
		if err := recover(); err != nil {
			log.Println("udp client error: ", err)
		}

		if p.conn != nil {
			_ = p.conn.Close()
		}

		if p.isRestart {
			time.AfterFunc(5*time.Second, func() {
				go p.Start()
			})
		}
	}()

	for {
		log.Println("net.Dial Udp Addr", p.addr)
		p.conn, err = net.Dial("udp", p.addr)
		if err != nil {
			log.Println("dial ", p.addr, " error：", err)
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}

	for {
		select {
		case payload, ok := <-p.sendChan:
			if !ok {
				break
			}

			_, err = p.conn.Write(payload)
			if err != nil {
				log.Println("UdpDataSender SendPacket: ", err)
				return
			}
			log.Println("UdpDataSender SendPacket Success: ", payload)
		}
	}
}

func (p *UdpDataSender) SendPacket(pack []byte) {
	if p.sendChan == nil {
		return
	}
	p.sendChan <- pack
}

func (p *UdpDataSender) Close() {
	p.isRestart = false

	if p.conn != nil {
		p.conn.Close()
	}
}
