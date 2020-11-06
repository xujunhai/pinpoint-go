package packet

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"net"
)

type TraceSendPacket struct {
	Type    int16
	TraceID int
	Length  int
	Payload []byte
}

func NewTraceSendPacket() *TraceSendPacket {
	return &TraceSendPacket{
		Type: ApplicationTraceSend,
	}
}

func (c *TraceSendPacket) Decode(conn net.Conn, reader io.Reader) error {
	buf := make([]byte, 8)

	if _, err := io.ReadFull(reader, buf); err != nil {
		log.Println("decode application request error")
		return err
	}

	c.TraceID = int(binary.BigEndian.Uint32(buf[0:4]))
	c.Length = int(binary.BigEndian.Uint32(buf[4:8]))

	if c.Length <= 0 {
		return nil
	}

	c.Payload = make([]byte, c.Length)

	if _, err := io.ReadFull(reader, c.Payload); err != nil {
		return err
	}

	return nil
}

// Encode ...
func (c *TraceSendPacket) Encode() ([]byte, error) {
	body := make([]byte, 10)
	binary.BigEndian.PutUint16(body[0:2], uint16(c.Type))
	binary.BigEndian.PutUint32(body[2:6], uint32(c.TraceID))
	binary.BigEndian.PutUint32(body[6:10], uint32(len(c.Payload)))

	bys := bytes.NewBuffer(body)
	bys.Write(c.Payload)

	return bys.Bytes(), nil
}

// GetPacketType ...
func (c *TraceSendPacket) GetPacketType() int16 {
	return c.Type
}

// GetPayload ...
func (c *TraceSendPacket) GetPayload() []byte {
	return c.Payload
}

// GetRequestID ...
func (c *TraceSendPacket) GetRequestID() int {
	return 0
}
