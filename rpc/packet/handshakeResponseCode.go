package packet

type HandshakeResponseCode struct {
	Code        int32
	SubCode     int32
	CodeMessage string
}

var (
	HsCode    = "code"
	HsSubCode = "subCode"
	HsCluster = "cluster"
)

var (
	HandshakeSuccess              = &HandshakeResponseCode{Code: 0, SubCode: 0, CodeMessage: "Success."}
	HandshakeSimplexCommunication = &HandshakeResponseCode{Code: 0, SubCode: 1, CodeMessage: "Simplex Connection successfully established."}
	HandshakeDuplexCommunication  = &HandshakeResponseCode{Code: 0, SubCode: 2, CodeMessage: "Duplex Connection successfully established."}

	HandshakeAlreadyKnown                = &HandshakeResponseCode{1, 0, "Already Known."}
	HandshakeAlreadySimplexCommunication = &HandshakeResponseCode{1, 1, "Already Simplex Connection established."}
	HandshakeAlreadyDuplexCommunication  = &HandshakeResponseCode{1, 2, "Already Duplex Connection established."}

	HandshakePropertyError = &HandshakeResponseCode{2, 0, "Property error."}

	HandshakeProtocolError = &HandshakeResponseCode{3, 0, "Illegal protocol error."}
	HandshakeUnknownError  = &HandshakeResponseCode{4, 0, "Unknown Error."}
	HandshakeUnknownCode   = &HandshakeResponseCode{-1, -1, "Unknown Code."}
	HandshakeSet           = []*HandshakeResponseCode{
		HandshakeSuccess,
		HandshakeSimplexCommunication,
		HandshakeDuplexCommunication,
		HandshakeAlreadyKnown,
		HandshakeAlreadySimplexCommunication,
		HandshakeAlreadyDuplexCommunication,
		HandshakePropertyError,
		HandshakeProtocolError,
		HandshakeUnknownError,
		HandshakeUnknownCode,
	}
)

func (hr *HandshakeResponseCode) getCode() int32 {
	return hr.Code
}

func (hr *HandshakeResponseCode) getSubCode() int32 {
	return hr.SubCode
}

func GetHandShakeCode(code int32, subCode int32) *HandshakeResponseCode {
	for _, v := range HandshakeSet {
		if code == v.getCode() && subCode == v.getSubCode() {
			return v
		}
	}

	for _, v := range HandshakeSet {
		if code == v.getCode() && 0 == v.getSubCode() {
			return v
		}
	}

	return HandshakeUnknownCode
}
