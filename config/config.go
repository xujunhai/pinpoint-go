package config

import "pinpoint-go/common/util"

const (
	UnknownAppName = "UnknownAppName"
	UnknownAgentId = "UnknownAgentId"
)

//指定grpc|thrift
type Config struct {
	AgentID         string
	ApplicationName string
	Pinpoint        struct {
		InfoAddr string //tcp agent info
		StatAddr string //udp agent stat
		SpanAddr string //udp span
	}
	Grpc       bool //使用grpc传输协议
	Thrift     bool //使用thrift传输协议
	Sampler    int  //span采样比
	ServerType int16
}

var globalCfg = NewDefaultConfig()

var agentStartTime int64

func init() {
	agentStartTime = util.GetNowMSec()
}

// InitConfig loads general configuration
func InitConfig(conf *Config) {
	globalCfg = conf
}

// NewDefaultConfig creates a new default config entity.
func NewDefaultConfig() *Config {
	agentId := util.GetHostIP() + util.GetRandomChar(4)
	return &Config{
		AgentID:         agentId,
		ApplicationName: UnknownAppName,
		Pinpoint: struct {
			InfoAddr string
			StatAddr string
			SpanAddr string
		}{
			InfoAddr: "127.0.0.1:9994",
			StatAddr: "127.0.0.1:9995",
			SpanAddr: "127.0.0.1:9996",
		},
		Grpc:       false,
		Thrift:     true,
		Sampler:    20,   //20%
		ServerType: 1502, //Go应用
	}
}

func AgentId() string {
	return globalCfg.AgentID
}

func AppName() string {
	return globalCfg.ApplicationName
}

func InfoAddr() string {
	return globalCfg.Pinpoint.InfoAddr
}

func StatAddr() string {
	return globalCfg.Pinpoint.StatAddr
}

func SpanAddr() string {
	return globalCfg.Pinpoint.SpanAddr
}

func ServerType() int16 {
	return globalCfg.ServerType
}

func StartTime() int64 {
	return agentStartTime
}
