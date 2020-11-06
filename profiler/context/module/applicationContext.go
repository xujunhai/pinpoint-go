package module

//
//import (
//	"pinpoint-go/bootstrap/config"
//	"pinpoint-go/profiler/monitor"
//	"pinpoint-go/profiler/sender"
//)
//
//type ApplicationContext struct {
//	ProfilerConfig   *config.ProfilerConfig
//	AgentInfoSender  *sender.AgentInfoSender
//	AgentStatMonitor *monitor.AgentStatMonitor
//	//traceContext trace.Context
//}
//
////两项监控
//func NewApplicationContext(profilerConfig *config.ProfilerConfig) ApplicationContext {
//
//	var dataSender sender.DataSender
//	switch profilerConfig.TransportModule() {
//	case config.THRIFT:
//		dataSender = sender.NewTcpDataSender()
//	case config.GRPC:
//	}
//
//	return ApplicationContext{
//		ProfilerConfig:   profilerConfig,
//		AgentInfoSender:  sender.NewAgentInfoSender(),
//		AgentStatMonitor: monitor.NewAgentStatMonitor(""),
//	}
//}
//
//func (ac ApplicationContext) Start() {
//	ac.AgentInfoSender.Start()
//	ac.AgentStatMonitor.Start()
//}
//
//func (ac ApplicationContext) Close() {
//}
