// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: service.proto

package trace

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0a, 0x53, 0x70, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x53, 0x74,
	0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x43, 0x6d, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0x40, 0x0a, 0x04, 0x53, 0x70, 0x61, 0x6e, 0x12, 0x38, 0x0a, 0x08, 0x53,
	0x65, 0x6e, 0x64, 0x53, 0x70, 0x61, 0x6e, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x53, 0x70,
	0x61, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x28, 0x01, 0x32, 0x65, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x31,
	0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x12, 0x29, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x09, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x50, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x32, 0xb5, 0x01, 0x0a,
	0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x12, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x53, 0x71, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x10, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x53, 0x71, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74,
	0x61, 0x1a, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00,
	0x12, 0x35, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x70, 0x69, 0x4d, 0x65,
	0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x41, 0x70, 0x69,
	0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x15, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74,
	0x61, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x00, 0x32, 0x45, 0x0a, 0x04, 0x53, 0x74, 0x61, 0x74, 0x12, 0x3d, 0x0a, 0x0d,
	0x53, 0x65, 0x6e, 0x64, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x12, 0x10, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x53, 0x74, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x32, 0x9a, 0x03, 0x0a, 0x16,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0d, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x43, 0x6d,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x43,
	0x6d, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12,
	0x3d, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x14,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x43, 0x6d, 0x64, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x5a,
	0x0a, 0x1e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x43, 0x6d, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12, 0x50, 0x0a, 0x17, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x43, 0x6d, 0x64, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x44, 0x75, 0x6d, 0x70, 0x52,
	0x65, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x1c,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x68, 0x72,
	0x65, 0x61, 0x64, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x20, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x43, 0x6d, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x44, 0x75, 0x6d, 0x70, 0x52, 0x65, 0x73, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x3c, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e,
	0x6e, 0x61, 0x76, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x70, 0x69, 0x6e, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x42, 0x0c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x07, 0x2e,
	0x3b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_proto_goTypes = []interface{}{
	(*PSpanMessage)(nil),                 // 0: v1.PSpanMessage
	(*PAgentInfo)(nil),                   // 1: v1.PAgentInfo
	(*PPing)(nil),                        // 2: v1.PPing
	(*PSqlMetaData)(nil),                 // 3: v1.PSqlMetaData
	(*PApiMetaData)(nil),                 // 4: v1.PApiMetaData
	(*PStringMetaData)(nil),              // 5: v1.PStringMetaData
	(*PStatMessage)(nil),                 // 6: v1.PStatMessage
	(*PCmdMessage)(nil),                  // 7: v1.PCmdMessage
	(*PCmdEchoResponse)(nil),             // 8: v1.PCmdEchoResponse
	(*PCmdActiveThreadCountRes)(nil),     // 9: v1.PCmdActiveThreadCountRes
	(*PCmdActiveThreadDumpRes)(nil),      // 10: v1.PCmdActiveThreadDumpRes
	(*PCmdActiveThreadLightDumpRes)(nil), // 11: v1.PCmdActiveThreadLightDumpRes
	(*empty.Empty)(nil),                  // 12: google.protobuf.Empty
	(*PResult)(nil),                      // 13: v1.PResult
	(*PCmdRequest)(nil),                  // 14: v1.PCmdRequest
}
var file_service_proto_depIdxs = []int32{
	0,  // 0: v1.Span.SendSpan:input_type -> v1.PSpanMessage
	1,  // 1: v1.Agent.RequestAgentInfo:input_type -> v1.PAgentInfo
	2,  // 2: v1.Agent.PingSession:input_type -> v1.PPing
	3,  // 3: v1.Metadata.RequestSqlMetaData:input_type -> v1.PSqlMetaData
	4,  // 4: v1.Metadata.RequestApiMetaData:input_type -> v1.PApiMetaData
	5,  // 5: v1.Metadata.RequestStringMetaData:input_type -> v1.PStringMetaData
	6,  // 6: v1.Stat.SendAgentStat:input_type -> v1.PStatMessage
	7,  // 7: v1.ProfilerCommandService.HandleCommand:input_type -> v1.PCmdMessage
	8,  // 8: v1.ProfilerCommandService.CommandEcho:input_type -> v1.PCmdEchoResponse
	9,  // 9: v1.ProfilerCommandService.CommandStreamActiveThreadCount:input_type -> v1.PCmdActiveThreadCountRes
	10, // 10: v1.ProfilerCommandService.CommandActiveThreadDump:input_type -> v1.PCmdActiveThreadDumpRes
	11, // 11: v1.ProfilerCommandService.CommandActiveThreadLightDump:input_type -> v1.PCmdActiveThreadLightDumpRes
	12, // 12: v1.Span.SendSpan:output_type -> google.protobuf.Empty
	13, // 13: v1.Agent.RequestAgentInfo:output_type -> v1.PResult
	2,  // 14: v1.Agent.PingSession:output_type -> v1.PPing
	13, // 15: v1.Metadata.RequestSqlMetaData:output_type -> v1.PResult
	13, // 16: v1.Metadata.RequestApiMetaData:output_type -> v1.PResult
	13, // 17: v1.Metadata.RequestStringMetaData:output_type -> v1.PResult
	12, // 18: v1.Stat.SendAgentStat:output_type -> google.protobuf.Empty
	14, // 19: v1.ProfilerCommandService.HandleCommand:output_type -> v1.PCmdRequest
	12, // 20: v1.ProfilerCommandService.CommandEcho:output_type -> google.protobuf.Empty
	12, // 21: v1.ProfilerCommandService.CommandStreamActiveThreadCount:output_type -> google.protobuf.Empty
	12, // 22: v1.ProfilerCommandService.CommandActiveThreadDump:output_type -> google.protobuf.Empty
	12, // 23: v1.ProfilerCommandService.CommandActiveThreadLightDump:output_type -> google.protobuf.Empty
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_Span_proto_init()
	file_Stat_proto_init()
	file_Cmd_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   5,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SpanClient is the client API for Span service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SpanClient interface {
	SendSpan(ctx context.Context, opts ...grpc.CallOption) (Span_SendSpanClient, error)
}

type spanClient struct {
	cc grpc.ClientConnInterface
}

func NewSpanClient(cc grpc.ClientConnInterface) SpanClient {
	return &spanClient{cc}
}

func (c *spanClient) SendSpan(ctx context.Context, opts ...grpc.CallOption) (Span_SendSpanClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Span_serviceDesc.Streams[0], "/v1.Span/SendSpan", opts...)
	if err != nil {
		return nil, err
	}
	x := &spanSendSpanClient{stream}
	return x, nil
}

type Span_SendSpanClient interface {
	Send(*PSpanMessage) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type spanSendSpanClient struct {
	grpc.ClientStream
}

func (x *spanSendSpanClient) Send(m *PSpanMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *spanSendSpanClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SpanServer is the server API for Span service.
type SpanServer interface {
	SendSpan(Span_SendSpanServer) error
}

// UnimplementedSpanServer can be embedded to have forward compatible implementations.
type UnimplementedSpanServer struct {
}

func (*UnimplementedSpanServer) SendSpan(Span_SendSpanServer) error {
	return status.Errorf(codes.Unimplemented, "method SendSpan not implemented")
}

func RegisterSpanServer(s *grpc.Server, srv SpanServer) {
	s.RegisterService(&_Span_serviceDesc, srv)
}

func _Span_SendSpan_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SpanServer).SendSpan(&spanSendSpanServer{stream})
}

type Span_SendSpanServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*PSpanMessage, error)
	grpc.ServerStream
}

type spanSendSpanServer struct {
	grpc.ServerStream
}

func (x *spanSendSpanServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *spanSendSpanServer) Recv() (*PSpanMessage, error) {
	m := new(PSpanMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Span_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Span",
	HandlerType: (*SpanServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendSpan",
			Handler:       _Span_SendSpan_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentClient interface {
	RequestAgentInfo(ctx context.Context, in *PAgentInfo, opts ...grpc.CallOption) (*PResult, error)
	PingSession(ctx context.Context, opts ...grpc.CallOption) (Agent_PingSessionClient, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) RequestAgentInfo(ctx context.Context, in *PAgentInfo, opts ...grpc.CallOption) (*PResult, error) {
	out := new(PResult)
	err := c.cc.Invoke(ctx, "/v1.Agent/RequestAgentInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) PingSession(ctx context.Context, opts ...grpc.CallOption) (Agent_PingSessionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Agent_serviceDesc.Streams[0], "/v1.Agent/PingSession", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentPingSessionClient{stream}
	return x, nil
}

type Agent_PingSessionClient interface {
	Send(*PPing) error
	Recv() (*PPing, error)
	grpc.ClientStream
}

type agentPingSessionClient struct {
	grpc.ClientStream
}

func (x *agentPingSessionClient) Send(m *PPing) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentPingSessionClient) Recv() (*PPing, error) {
	m := new(PPing)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentServer is the server API for Agent service.
type AgentServer interface {
	RequestAgentInfo(context.Context, *PAgentInfo) (*PResult, error)
	PingSession(Agent_PingSessionServer) error
}

// UnimplementedAgentServer can be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (*UnimplementedAgentServer) RequestAgentInfo(context.Context, *PAgentInfo) (*PResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestAgentInfo not implemented")
}
func (*UnimplementedAgentServer) PingSession(Agent_PingSessionServer) error {
	return status.Errorf(codes.Unimplemented, "method PingSession not implemented")
}

func RegisterAgentServer(s *grpc.Server, srv AgentServer) {
	s.RegisterService(&_Agent_serviceDesc, srv)
}

func _Agent_RequestAgentInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PAgentInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).RequestAgentInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Agent/RequestAgentInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).RequestAgentInfo(ctx, req.(*PAgentInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_PingSession_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServer).PingSession(&agentPingSessionServer{stream})
}

type Agent_PingSessionServer interface {
	Send(*PPing) error
	Recv() (*PPing, error)
	grpc.ServerStream
}

type agentPingSessionServer struct {
	grpc.ServerStream
}

func (x *agentPingSessionServer) Send(m *PPing) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentPingSessionServer) Recv() (*PPing, error) {
	m := new(PPing)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Agent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestAgentInfo",
			Handler:    _Agent_RequestAgentInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PingSession",
			Handler:       _Agent_PingSession_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}

// MetadataClient is the client API for Metadata service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetadataClient interface {
	RequestSqlMetaData(ctx context.Context, in *PSqlMetaData, opts ...grpc.CallOption) (*PResult, error)
	RequestApiMetaData(ctx context.Context, in *PApiMetaData, opts ...grpc.CallOption) (*PResult, error)
	RequestStringMetaData(ctx context.Context, in *PStringMetaData, opts ...grpc.CallOption) (*PResult, error)
}

type metadataClient struct {
	cc grpc.ClientConnInterface
}

func NewMetadataClient(cc grpc.ClientConnInterface) MetadataClient {
	return &metadataClient{cc}
}

func (c *metadataClient) RequestSqlMetaData(ctx context.Context, in *PSqlMetaData, opts ...grpc.CallOption) (*PResult, error) {
	out := new(PResult)
	err := c.cc.Invoke(ctx, "/v1.Metadata/RequestSqlMetaData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataClient) RequestApiMetaData(ctx context.Context, in *PApiMetaData, opts ...grpc.CallOption) (*PResult, error) {
	out := new(PResult)
	err := c.cc.Invoke(ctx, "/v1.Metadata/RequestApiMetaData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataClient) RequestStringMetaData(ctx context.Context, in *PStringMetaData, opts ...grpc.CallOption) (*PResult, error) {
	out := new(PResult)
	err := c.cc.Invoke(ctx, "/v1.Metadata/RequestStringMetaData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetadataServer is the server API for Metadata service.
type MetadataServer interface {
	RequestSqlMetaData(context.Context, *PSqlMetaData) (*PResult, error)
	RequestApiMetaData(context.Context, *PApiMetaData) (*PResult, error)
	RequestStringMetaData(context.Context, *PStringMetaData) (*PResult, error)
}

// UnimplementedMetadataServer can be embedded to have forward compatible implementations.
type UnimplementedMetadataServer struct {
}

func (*UnimplementedMetadataServer) RequestSqlMetaData(context.Context, *PSqlMetaData) (*PResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestSqlMetaData not implemented")
}
func (*UnimplementedMetadataServer) RequestApiMetaData(context.Context, *PApiMetaData) (*PResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestApiMetaData not implemented")
}
func (*UnimplementedMetadataServer) RequestStringMetaData(context.Context, *PStringMetaData) (*PResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestStringMetaData not implemented")
}

func RegisterMetadataServer(s *grpc.Server, srv MetadataServer) {
	s.RegisterService(&_Metadata_serviceDesc, srv)
}

func _Metadata_RequestSqlMetaData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PSqlMetaData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).RequestSqlMetaData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Metadata/RequestSqlMetaData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).RequestSqlMetaData(ctx, req.(*PSqlMetaData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metadata_RequestApiMetaData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PApiMetaData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).RequestApiMetaData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Metadata/RequestApiMetaData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).RequestApiMetaData(ctx, req.(*PApiMetaData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metadata_RequestStringMetaData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PStringMetaData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).RequestStringMetaData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Metadata/RequestStringMetaData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).RequestStringMetaData(ctx, req.(*PStringMetaData))
	}
	return interceptor(ctx, in, info, handler)
}

var _Metadata_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Metadata",
	HandlerType: (*MetadataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestSqlMetaData",
			Handler:    _Metadata_RequestSqlMetaData_Handler,
		},
		{
			MethodName: "RequestApiMetaData",
			Handler:    _Metadata_RequestApiMetaData_Handler,
		},
		{
			MethodName: "RequestStringMetaData",
			Handler:    _Metadata_RequestStringMetaData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

// StatClient is the client API for Stat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StatClient interface {
	SendAgentStat(ctx context.Context, opts ...grpc.CallOption) (Stat_SendAgentStatClient, error)
}

type statClient struct {
	cc grpc.ClientConnInterface
}

func NewStatClient(cc grpc.ClientConnInterface) StatClient {
	return &statClient{cc}
}

func (c *statClient) SendAgentStat(ctx context.Context, opts ...grpc.CallOption) (Stat_SendAgentStatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Stat_serviceDesc.Streams[0], "/v1.Stat/SendAgentStat", opts...)
	if err != nil {
		return nil, err
	}
	x := &statSendAgentStatClient{stream}
	return x, nil
}

type Stat_SendAgentStatClient interface {
	Send(*PStatMessage) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type statSendAgentStatClient struct {
	grpc.ClientStream
}

func (x *statSendAgentStatClient) Send(m *PStatMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *statSendAgentStatClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StatServer is the server API for Stat service.
type StatServer interface {
	SendAgentStat(Stat_SendAgentStatServer) error
}

// UnimplementedStatServer can be embedded to have forward compatible implementations.
type UnimplementedStatServer struct {
}

func (*UnimplementedStatServer) SendAgentStat(Stat_SendAgentStatServer) error {
	return status.Errorf(codes.Unimplemented, "method SendAgentStat not implemented")
}

func RegisterStatServer(s *grpc.Server, srv StatServer) {
	s.RegisterService(&_Stat_serviceDesc, srv)
}

func _Stat_SendAgentStat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StatServer).SendAgentStat(&statSendAgentStatServer{stream})
}

type Stat_SendAgentStatServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*PStatMessage, error)
	grpc.ServerStream
}

type statSendAgentStatServer struct {
	grpc.ServerStream
}

func (x *statSendAgentStatServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *statSendAgentStatServer) Recv() (*PStatMessage, error) {
	m := new(PStatMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Stat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Stat",
	HandlerType: (*StatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendAgentStat",
			Handler:       _Stat_SendAgentStat_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}

// ProfilerCommandServiceClient is the client API for ProfilerCommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProfilerCommandServiceClient interface {
	HandleCommand(ctx context.Context, opts ...grpc.CallOption) (ProfilerCommandService_HandleCommandClient, error)
	CommandEcho(ctx context.Context, in *PCmdEchoResponse, opts ...grpc.CallOption) (*empty.Empty, error)
	CommandStreamActiveThreadCount(ctx context.Context, opts ...grpc.CallOption) (ProfilerCommandService_CommandStreamActiveThreadCountClient, error)
	CommandActiveThreadDump(ctx context.Context, in *PCmdActiveThreadDumpRes, opts ...grpc.CallOption) (*empty.Empty, error)
	CommandActiveThreadLightDump(ctx context.Context, in *PCmdActiveThreadLightDumpRes, opts ...grpc.CallOption) (*empty.Empty, error)
}

type profilerCommandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfilerCommandServiceClient(cc grpc.ClientConnInterface) ProfilerCommandServiceClient {
	return &profilerCommandServiceClient{cc}
}

func (c *profilerCommandServiceClient) HandleCommand(ctx context.Context, opts ...grpc.CallOption) (ProfilerCommandService_HandleCommandClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProfilerCommandService_serviceDesc.Streams[0], "/v1.ProfilerCommandService/HandleCommand", opts...)
	if err != nil {
		return nil, err
	}
	x := &profilerCommandServiceHandleCommandClient{stream}
	return x, nil
}

type ProfilerCommandService_HandleCommandClient interface {
	Send(*PCmdMessage) error
	Recv() (*PCmdRequest, error)
	grpc.ClientStream
}

type profilerCommandServiceHandleCommandClient struct {
	grpc.ClientStream
}

func (x *profilerCommandServiceHandleCommandClient) Send(m *PCmdMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *profilerCommandServiceHandleCommandClient) Recv() (*PCmdRequest, error) {
	m := new(PCmdRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *profilerCommandServiceClient) CommandEcho(ctx context.Context, in *PCmdEchoResponse, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v1.ProfilerCommandService/CommandEcho", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilerCommandServiceClient) CommandStreamActiveThreadCount(ctx context.Context, opts ...grpc.CallOption) (ProfilerCommandService_CommandStreamActiveThreadCountClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProfilerCommandService_serviceDesc.Streams[1], "/v1.ProfilerCommandService/CommandStreamActiveThreadCount", opts...)
	if err != nil {
		return nil, err
	}
	x := &profilerCommandServiceCommandStreamActiveThreadCountClient{stream}
	return x, nil
}

type ProfilerCommandService_CommandStreamActiveThreadCountClient interface {
	Send(*PCmdActiveThreadCountRes) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type profilerCommandServiceCommandStreamActiveThreadCountClient struct {
	grpc.ClientStream
}

func (x *profilerCommandServiceCommandStreamActiveThreadCountClient) Send(m *PCmdActiveThreadCountRes) error {
	return x.ClientStream.SendMsg(m)
}

func (x *profilerCommandServiceCommandStreamActiveThreadCountClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *profilerCommandServiceClient) CommandActiveThreadDump(ctx context.Context, in *PCmdActiveThreadDumpRes, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v1.ProfilerCommandService/CommandActiveThreadDump", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilerCommandServiceClient) CommandActiveThreadLightDump(ctx context.Context, in *PCmdActiveThreadLightDumpRes, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v1.ProfilerCommandService/CommandActiveThreadLightDump", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfilerCommandServiceServer is the server API for ProfilerCommandService service.
type ProfilerCommandServiceServer interface {
	HandleCommand(ProfilerCommandService_HandleCommandServer) error
	CommandEcho(context.Context, *PCmdEchoResponse) (*empty.Empty, error)
	CommandStreamActiveThreadCount(ProfilerCommandService_CommandStreamActiveThreadCountServer) error
	CommandActiveThreadDump(context.Context, *PCmdActiveThreadDumpRes) (*empty.Empty, error)
	CommandActiveThreadLightDump(context.Context, *PCmdActiveThreadLightDumpRes) (*empty.Empty, error)
}

// UnimplementedProfilerCommandServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProfilerCommandServiceServer struct {
}

func (*UnimplementedProfilerCommandServiceServer) HandleCommand(ProfilerCommandService_HandleCommandServer) error {
	return status.Errorf(codes.Unimplemented, "method HandleCommand not implemented")
}
func (*UnimplementedProfilerCommandServiceServer) CommandEcho(context.Context, *PCmdEchoResponse) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommandEcho not implemented")
}
func (*UnimplementedProfilerCommandServiceServer) CommandStreamActiveThreadCount(ProfilerCommandService_CommandStreamActiveThreadCountServer) error {
	return status.Errorf(codes.Unimplemented, "method CommandStreamActiveThreadCount not implemented")
}
func (*UnimplementedProfilerCommandServiceServer) CommandActiveThreadDump(context.Context, *PCmdActiveThreadDumpRes) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommandActiveThreadDump not implemented")
}
func (*UnimplementedProfilerCommandServiceServer) CommandActiveThreadLightDump(context.Context, *PCmdActiveThreadLightDumpRes) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommandActiveThreadLightDump not implemented")
}

func RegisterProfilerCommandServiceServer(s *grpc.Server, srv ProfilerCommandServiceServer) {
	s.RegisterService(&_ProfilerCommandService_serviceDesc, srv)
}

func _ProfilerCommandService_HandleCommand_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProfilerCommandServiceServer).HandleCommand(&profilerCommandServiceHandleCommandServer{stream})
}

type ProfilerCommandService_HandleCommandServer interface {
	Send(*PCmdRequest) error
	Recv() (*PCmdMessage, error)
	grpc.ServerStream
}

type profilerCommandServiceHandleCommandServer struct {
	grpc.ServerStream
}

func (x *profilerCommandServiceHandleCommandServer) Send(m *PCmdRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *profilerCommandServiceHandleCommandServer) Recv() (*PCmdMessage, error) {
	m := new(PCmdMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProfilerCommandService_CommandEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PCmdEchoResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilerCommandServiceServer).CommandEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ProfilerCommandService/CommandEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilerCommandServiceServer).CommandEcho(ctx, req.(*PCmdEchoResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfilerCommandService_CommandStreamActiveThreadCount_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProfilerCommandServiceServer).CommandStreamActiveThreadCount(&profilerCommandServiceCommandStreamActiveThreadCountServer{stream})
}

type ProfilerCommandService_CommandStreamActiveThreadCountServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*PCmdActiveThreadCountRes, error)
	grpc.ServerStream
}

type profilerCommandServiceCommandStreamActiveThreadCountServer struct {
	grpc.ServerStream
}

func (x *profilerCommandServiceCommandStreamActiveThreadCountServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *profilerCommandServiceCommandStreamActiveThreadCountServer) Recv() (*PCmdActiveThreadCountRes, error) {
	m := new(PCmdActiveThreadCountRes)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProfilerCommandService_CommandActiveThreadDump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PCmdActiveThreadDumpRes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilerCommandServiceServer).CommandActiveThreadDump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ProfilerCommandService/CommandActiveThreadDump",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilerCommandServiceServer).CommandActiveThreadDump(ctx, req.(*PCmdActiveThreadDumpRes))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfilerCommandService_CommandActiveThreadLightDump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PCmdActiveThreadLightDumpRes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilerCommandServiceServer).CommandActiveThreadLightDump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ProfilerCommandService/CommandActiveThreadLightDump",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilerCommandServiceServer).CommandActiveThreadLightDump(ctx, req.(*PCmdActiveThreadLightDumpRes))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProfilerCommandService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ProfilerCommandService",
	HandlerType: (*ProfilerCommandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommandEcho",
			Handler:    _ProfilerCommandService_CommandEcho_Handler,
		},
		{
			MethodName: "CommandActiveThreadDump",
			Handler:    _ProfilerCommandService_CommandActiveThreadDump_Handler,
		},
		{
			MethodName: "CommandActiveThreadLightDump",
			Handler:    _ProfilerCommandService_CommandActiveThreadLightDump_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HandleCommand",
			Handler:       _ProfilerCommandService_HandleCommand_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "CommandStreamActiveThreadCount",
			Handler:       _ProfilerCommandService_CommandStreamActiveThreadCount_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}
