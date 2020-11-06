package stream

// Status Code
var (
	OK           uint16 = 0
	UnknownError uint16 = 100
	IdError      uint16 = 110
	IdIllegal    uint16 = 111
	IdDuplicated uint16 = 112
	IdNotFound   uint16 = 113

	StateError        uint16 = 120
	StateNotConnected uint16 = 121
	StateClosed       uint16 = 122

	TypeError     uint16 = 130
	TypeUnknown   uint16 = 131
	TypeUnSupport uint16 = 132

	PacketError         uint16 = 140
	PacketUnknown       uint16 = 141
	PacketUnSupport     uint16 = 142
	ConnectionError     uint16 = 150
	ConnectionNotFound  uint16 = 151
	ConnectionTimeout   uint16 = 152
	ConnectionUnSupport uint16 = 153
	RouteError          uint16 = 160
	StatusCode                 = [...]uint16{0, 100, 110, 111, 112, 113, 120, 121, 122, 130, 131, 132, 140, 141, 142, 150, 151, 152, 153, 160}
)

var codeMap map[uint16]StreamCode

func init() {
	codeMap = make(map[uint16]StreamCode)
	for _, v := range StatusCode {
		codeMap[v] = StreamCode{value: v}
	}
}

type StreamCode struct {
	value uint16
}

func IsConnectionError(streamCode StreamCode) bool {
	if ConnectionError == streamCode.value || ConnectionNotFound == streamCode.value || ConnectionTimeout == streamCode.value || ConnectionUnSupport == streamCode.value {
		return true
	}
	return false
}

func GetCode(value uint16) StreamCode {
	if v, ok := codeMap[value]; ok {
		return v
	}

	codeGroup := value - (value % 10)
	if v, ok := codeMap[codeGroup]; ok {
		return v
	}

	return StreamCode{value: UnknownError}
}
