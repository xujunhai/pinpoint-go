package packet

import (
	"io"
	"net"
)

type Packet interface {
	Decode(conn net.Conn, reader io.Reader) error
	Encode() ([]byte, error)
	GetPacketType() int16
	GetPayload() []byte
	GetRequestID() int
}
