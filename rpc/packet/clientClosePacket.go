package packet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"log"
	"net"
)

type ClientClosePacket struct {
	Type    int16
	Length  int
	Payload []byte
}

func NewClientClosePacket() *ClientClosePacket {
	return &ClientClosePacket{
		Type: ControlClientClose,
	}
}

// Decode ...
func (c *ClientClosePacket) Decode(conn net.Conn, reader io.Reader) error {
	buf := make([]byte, 4)
	if _, err := io.ReadFull(reader, buf); err != nil {
		log.Println("control client close: ", err)
		return err
	}
	c.Length = int(binary.BigEndian.Uint32(buf[:4]))

	if c.Length >= 10*1024 {
		return errors.New("关闭包太大了")
	}

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
func (c *ClientClosePacket) Encode() ([]byte, error) {
	body := make([]byte, 6)
	binary.BigEndian.PutUint16(body[0:2], uint16(c.Type))
	binary.BigEndian.PutUint32(body[2:6], uint32(len(c.Payload)))
	bys := bytes.NewBuffer(body)
	bys.Write(c.Payload)

	return bys.Bytes(), nil
}

// GetPacketType ...
func (c *ClientClosePacket) GetPacketType() int16 {
	return c.Type
}

// GetPayload ...
func (c *ClientClosePacket) GetPayload() []byte {
	return c.Payload
}

// GetRequestID ...
func (c *ClientClosePacket) GetRequestID() int {
	return 0
}
