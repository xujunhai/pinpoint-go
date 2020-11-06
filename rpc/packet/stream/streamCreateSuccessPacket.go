package stream

import (
	"encoding/binary"
	"io"
	"log"
	"net"
)

type StreamCreateSuccessPacket struct {
	Type      int16
	ChannelID int
}

func NewStreamCreateSuccessPacket() *StreamCreateSuccessPacket {
	return &StreamCreateSuccessPacket{
		Type: ApplicationStreamCreateSuccess,
	}
}

// Decode ...
func (p *StreamCreateSuccessPacket) Decode(conn net.Conn, reader io.Reader) error {
	buf := make([]byte, 4)
	if _, err := io.ReadFull(reader, buf); err != nil {
		log.Println("decode application request error")
		return err
	}
	p.ChannelID = int(binary.BigEndian.Uint32(buf[:4]))
	return nil
}

// Encode ...
func (p *StreamCreateSuccessPacket) Encode() ([]byte, error) {
	buf := make([]byte, 6)
	//2 + 4
	binary.BigEndian.PutUint16(buf[0:2], uint16(p.Type))
	binary.BigEndian.PutUint16(buf[2:4], uint16(p.ChannelID))
	return buf, nil
}

// GetPacketType ...
func (p *StreamCreateSuccessPacket) GetPacketType() int16 {
	return p.Type
}

// GetPayload ...
func (p *StreamCreateSuccessPacket) GetPayload() []byte {
	return nil
}

// GetRequestID ...
func (p *StreamCreateSuccessPacket) GetRequestID() int {
	return 0
}
