package packet

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"net"
)

type ControlHandshakeResponsePacket struct {
	Type      int16
	RequestID int
	Length    int
	Payload   []byte
}

func NewControlHandshakeResponsePacket() *ControlHandshakeResponsePacket {
	return &ControlHandshakeResponsePacket{
		Type: ControlHandshakeResponse,
	}
}

// Decode ...
func (c *ControlHandshakeResponsePacket) Decode(conn net.Conn, reader io.Reader) error {
	buf := make([]byte, 8)
	if _, err := io.ReadFull(reader, buf); err != nil {
		log.Println("control handshake decode error ", err)
		return err
	}
	c.RequestID = int(binary.BigEndian.Uint32(buf[:4]))
	c.Length = int(binary.BigEndian.Uint32(buf[4:8]))
	c.Payload = make([]byte, c.Length)

	if _, err := io.ReadFull(reader, c.Payload); err != nil {
		return err
	}

	//log.Println(string(c.Payload))
	return nil
}

// Encode ...
func (c *ControlHandshakeResponsePacket) Encode() ([]byte, error) {
	buf := make([]byte, 10)
	//2 + 4 + 4
	binary.BigEndian.PutUint16(buf[0:2], uint16(c.Type))
	binary.BigEndian.PutUint32(buf[2:6], uint32(c.RequestID))
	binary.BigEndian.PutUint32(buf[6:10], uint32(len(c.Payload)))
	bys := bytes.NewBuffer(buf)
	bys.Write(c.Payload)

	return bys.Bytes(), nil
}

// GetPacketType ...
func (c *ControlHandshakeResponsePacket) GetPacketType() int16 {
	return c.Type
}

// GetPayload ...
func (c *ControlHandshakeResponsePacket) GetPayload() []byte {
	return c.Payload
}

func (c *ControlHandshakeResponsePacket) GetRequestID() int {
	return c.RequestID
}
