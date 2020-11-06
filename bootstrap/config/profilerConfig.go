package config

const (
	DefaultAgentStatCollectionIntervalMs = 5 * 1000
	DefaultNumAgentStatBatchSend         = 6
	DefaultTransportModule               = THRIFT
	DefaultAgentInfoSendRetryInterval    = 5 * 60 * 1000
)

type ProfilerConfig struct {
	transportModule       TransportModule //thrfit or grpc
	thriftTransportConfig ThriftTransportConfig

	// Sampling
	samplingEnable             bool //true
	samplingRate               int  //1
	samplingNewThroughput      int  //0;
	samplingContinueThroughput int  // 0;

	agentInfoSendRetryInterval int64 //DefaultAgentInfoSendRetryInterval
	applicationServerType      string

	proxyHttpHeaderEnable bool //true
	httpStatusCodeErrors  HttpStatusCodeErrors
}

func (p ProfilerConfig) TransportModule() TransportModule {
	return p.transportModule
}
