package packet

import (
	"encoding/binary"
	"io"
	"net"
)

type PingSimplePacket struct {
	Type int16
}

func NewPingSimplePacket() *PingSimplePacket {
	return &PingSimplePacket{
		Type: ControlPingSimple,
	}
}

// Decode ...
func (c *PingSimplePacket) Decode(conn net.Conn, reader io.Reader) error {
	return nil
}

// Encode ...
func (c *PingSimplePacket) Encode() ([]byte, error) {
	body := make([]byte, 2)
	binary.BigEndian.PutUint16(body[0:2], uint16(c.Type))
	return body, nil
}

// GetPacketType ...
func (c *PingSimplePacket) GetPacketType() int16 {
	return c.Type
}

// GetPayload ...
func (c *PingSimplePacket) GetPayload() []byte {
	return nil
}

func (c *PingSimplePacket) GetRequestID() int {
	return 0
}
