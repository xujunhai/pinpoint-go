package config

const (
	DefaultIp                                        = "127.0.0.1"
	DefaultDataSenderPinpointClientWriteTimeout      = 3 * 1000
	DefaultDataSenderPinpointClientRequestTimeout    = 3 * 1000
	DefaultDataSenderPinpointClientReconnectInterval = 3 * 1000

	DefaultDataSenderPinpointClientPingInterval      = 60 * 1000 * 5
	DefaultDataSenderPinpointClientHandshakeInterval = 60 * 1000 * 1
)

type ThriftTransportConfig struct {
	collectorSpanServerIp   string
	collectorSpanServerPort int //9996

	collectorStatServerIp   string
	collectorStatServerPort int //9995

	collectorTcpServerIp   string
	collectorTcpServerPort int //9994

	spanDataSenderWriteQueueSize       int    //1024 * 5
	spanDataSenderSocketSendBufferSize int    //1024 * 64 * 16
	spanDataSenderSocketTimeout        int    //1000 * 3
	spanDataSenderChunkSize            int    //1024 * 16
	spanDataSenderTransportType        string //udp
	spanDataSenderSocketType           string //oio

	statDataSenderWriteQueueSize       int    //1024 * 5
	statDataSenderSocketSendBufferSize int    // 1024 * 64 * 16
	statDataSenderSocketTimeout        int    //1000 * 3
	statDataSenderChunkSize            int    //1024 * 16
	statDataSenderTransportType        string // "UDP"
	statDataSenderSocketType           string // "OIO"

	tcpDataSenderPinpointClientWriteTimeout      int64
	tcpDataSenderPinpointClientRequestTimeout    int64
	tcpDataSenderPinpointClientReconnectInterval int64
	tcpDataSenderPinpointClientPingInterval      int64
	tcpDataSenderPinpointClientHandshakeInterval int64
}
