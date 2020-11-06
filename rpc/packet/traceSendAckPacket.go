package packet

import (
	"encoding/binary"
	"io"
	"log"
	"net"
)

type TraceSendAckPacket struct {
	Type    int16
	TraceID int
}

func NewTraceSendAckPacket() *TraceSendAckPacket {
	return &TraceSendAckPacket{
		Type: ApplicationTraceSendAck,
	}
}

func (c *TraceSendAckPacket) Decode(conn net.Conn, reader io.Reader) error {
	buf := make([]byte, 4)

	if _, err := io.ReadFull(reader, buf); err != nil {
		log.Println("decode application request error")
		return err
	}

	c.TraceID = int(binary.BigEndian.Uint32(buf[0:4]))
	return nil
}

// Encode ...
func (c *TraceSendAckPacket) Encode() ([]byte, error) {
	body := make([]byte, 6)
	binary.BigEndian.PutUint16(body[0:2], uint16(c.Type))
	binary.BigEndian.PutUint32(body[2:6], uint32(c.TraceID))

	return body, nil
}

// GetPacketType ...
func (c *TraceSendAckPacket) GetPacketType() int16 {
	return c.Type
}

// GetPayload ...
func (c *TraceSendAckPacket) GetPayload() []byte {
	return nil
}

// GetRequestID ...
func (c *TraceSendAckPacket) GetRequestID() int {
	return 0
}
