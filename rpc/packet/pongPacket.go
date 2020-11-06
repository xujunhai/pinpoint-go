package packet

import (
	"encoding/binary"
	"io"
	"net"
)

type PongPacket struct {
	Type int16
}

func NewPongPacket() *PongPacket {
	return &PongPacket{
		Type: ControlPong,
	}
}

func (c *PongPacket) Decode(conn net.Conn, reader io.Reader) error {

	return nil
}

// Encode ...
func (c *PongPacket) Encode() ([]byte, error) {
	body := make([]byte, 2)
	binary.BigEndian.PutUint16(body[0:2], uint16(c.Type))
	return body, nil
}

// GetPacketType ...
func (c *PongPacket) GetPacketType() int16 {
	return c.Type
}

// GetPayload ...
func (c *PongPacket) GetPayload() []byte {
	return nil
}

// GetRequestID ...
func (c *PongPacket) GetRequestID() int {
	return 0
}
