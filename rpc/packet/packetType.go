package packet

//PackageType
var (
	ApplicationSend         int16 = 1
	ApplicationTraceSend    int16 = 2
	ApplicationTraceSendAck int16 = 3

	ApplicationRequest  int16 = 5
	ApplicationResponse int16 = 6

	ControlClientClose int16 = 100
	ControlServerClose int16 = 110

	// control packet
	ControlHandshake         int16 = 150
	ControlHandshakeResponse int16 = 151

	// keep stay because of performance in case of ping and pong. others removed.
	// CONTROL_PING will be deprecated. caused : Two payload types are used in one control packet.
	// since 1.7.0, use CONTROL_PING_SIMPLE, CONTROL_PING_PAYLOAD
	//@Deprecated
	ControlPing int16 = 200
	ControlPong int16 = 201

	ControlPingSimple  int16 = 210
	ControlPingPayload int16 = 211

	UNKNOWN int16 = 500

	PacketTypeSize int16 = 2
)
