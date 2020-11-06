package config

type TransportModule uint8

const (
	THRIFT TransportModule = 0
	GRPC   TransportModule = 1
)
