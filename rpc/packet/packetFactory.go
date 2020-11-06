package packet

import (
	"errors"
	"io"
	"net"
	"pinpoint-go/rpc/packet/stream"
)

func DecodePacketFactory(packetType int16, conn net.Conn, reader io.Reader) (Packet, error) {

	var pack Packet
	var err error
	switch packetType {
	case ApplicationSend:
		pack = NewSendPacket()
	case ApplicationTraceSend:
		pack = NewTraceSendPacket()
	case ApplicationTraceSendAck:
		pack = NewTraceSendAckPacket()
	case ApplicationRequest:
		pack = NewRequestPacket()
	case ApplicationResponse:
		pack = NewResponsePacket()
	case stream.ApplicationStreamCreate:
		pack = stream.NewStreamCreatePacket()
	case stream.ApplicationStreamCreateSuccess:
		pack = stream.NewStreamCreateSuccessPacket()
	case stream.ApplicationStreamCreateFail:
		pack = stream.NewStreamCreateFailPacket()
	case stream.ApplicationStreamClose:
		pack = stream.NewStreamClosePacket()
	case stream.ApplicationStreamPing:
		pack = stream.NewStreamPingPacket()
	case stream.ApplicationStreamPong:
		pack = stream.NewStreamPongPacket()
	case stream.ApplicationStreamResponse:
		pack = stream.NewStreamResponsePacket()
	case ControlClientClose:
		pack = NewClientClosePacket()
	case ControlServerClose:
		pack = NewServerClosePacket()
	case ControlHandshake:
		pack = NewControlHandshakePacket()
	case ControlHandshakeResponse:
		pack = NewControlHandshakeResponsePacket()
	case ControlPong:
		pack = NewPongPacket()
	case ControlPingSimple:
		pack = NewPingSimplePacket()
	case ControlPingPayload:
		pack = NewPingPayloadPacket()

	default:
		err = errors.New("unknown packet type")
	}

	if pack != nil {
		err = pack.Decode(conn, reader)
	}

	return pack, err
}
