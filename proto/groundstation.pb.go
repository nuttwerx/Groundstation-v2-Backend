// Code generated by protoc-gen-go. DO NOT EDIT.
// source: groundstation.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	groundstation.proto

It has these top-level messages:
	StreamRequest
	ServerControl
	StartLogCommand
	Command
	Ack
	ServerStatus
	OpenPort
	Pong
	Ping
	DataBundle
	DataPacket
	Params
	Value
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type ServerControl_CommandEnum int32

const (
	ServerControl_LogServiceStart       ServerControl_CommandEnum = 0
	ServerControl_LogServiceStop        ServerControl_CommandEnum = 1
	ServerControl_DataStoreManagerStart ServerControl_CommandEnum = 2
	ServerControl_DataStoreManagerStop  ServerControl_CommandEnum = 3
	ServerControl_BroadcasterStart      ServerControl_CommandEnum = 4
	ServerControl_BroadcasterStop       ServerControl_CommandEnum = 5
)

var ServerControl_CommandEnum_name = map[int32]string{
	0: "LogServiceStart",
	1: "LogServiceStop",
	2: "DataStoreManagerStart",
	3: "DataStoreManagerStop",
	4: "BroadcasterStart",
	5: "BroadcasterStop",
}
var ServerControl_CommandEnum_value = map[string]int32{
	"LogServiceStart":       0,
	"LogServiceStop":        1,
	"DataStoreManagerStart": 2,
	"DataStoreManagerStop":  3,
	"BroadcasterStart":      4,
	"BroadcasterStop":       5,
}

func (x ServerControl_CommandEnum) String() string {
	return proto1.EnumName(ServerControl_CommandEnum_name, int32(x))
}
func (ServerControl_CommandEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type StreamRequest struct {
}

func (m *StreamRequest) Reset()                    { *m = StreamRequest{} }
func (m *StreamRequest) String() string            { return proto1.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()               {}
func (*StreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ServerControl struct {
	Command ServerControl_CommandEnum `protobuf:"varint,1,opt,name=Command,enum=proto.ServerControl_CommandEnum" json:"Command,omitempty"`
}

func (m *ServerControl) Reset()                    { *m = ServerControl{} }
func (m *ServerControl) String() string            { return proto1.CompactTextString(m) }
func (*ServerControl) ProtoMessage()               {}
func (*ServerControl) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ServerControl) GetCommand() ServerControl_CommandEnum {
	if m != nil {
		return m.Command
	}
	return ServerControl_LogServiceStart
}

type StartLogCommand struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *StartLogCommand) Reset()                    { *m = StartLogCommand{} }
func (m *StartLogCommand) String() string            { return proto1.CompactTextString(m) }
func (*StartLogCommand) ProtoMessage()               {}
func (*StartLogCommand) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StartLogCommand) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Command struct {
	Node       string  `protobuf:"bytes,1,opt,name=Node" json:"Node,omitempty"`
	PacketType int32   `protobuf:"varint,2,opt,name=PacketType" json:"PacketType,omitempty"`
	Data       []int32 `protobuf:"varint,3,rep,packed,name=Data" json:"Data,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto1.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Command) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *Command) GetPacketType() int32 {
	if m != nil {
		return m.PacketType
	}
	return 0
}

func (m *Command) GetData() []int32 {
	if m != nil {
		return m.Data
	}
	return nil
}

type Ack struct {
}

func (m *Ack) Reset()                    { *m = Ack{} }
func (m *Ack) String() string            { return proto1.CompactTextString(m) }
func (*Ack) ProtoMessage()               {}
func (*Ack) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ServerStatus struct {
	DataStoreManagerRunning bool        `protobuf:"varint,1,opt,name=DataStoreManagerRunning" json:"DataStoreManagerRunning,omitempty"`
	GRPCServerRunning       bool        `protobuf:"varint,2,opt,name=GRPCServerRunning" json:"GRPCServerRunning,omitempty"`
	BroadcasterRunning      bool        `protobuf:"varint,3,opt,name=BroadcasterRunning" json:"BroadcasterRunning,omitempty"`
	GSLoggerRunning         bool        `protobuf:"varint,4,opt,name=GSLoggerRunning" json:"GSLoggerRunning,omitempty"`
	OpenPorts               []*OpenPort `protobuf:"bytes,5,rep,name=OpenPorts" json:"OpenPorts,omitempty"`
}

func (m *ServerStatus) Reset()                    { *m = ServerStatus{} }
func (m *ServerStatus) String() string            { return proto1.CompactTextString(m) }
func (*ServerStatus) ProtoMessage()               {}
func (*ServerStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ServerStatus) GetDataStoreManagerRunning() bool {
	if m != nil {
		return m.DataStoreManagerRunning
	}
	return false
}

func (m *ServerStatus) GetGRPCServerRunning() bool {
	if m != nil {
		return m.GRPCServerRunning
	}
	return false
}

func (m *ServerStatus) GetBroadcasterRunning() bool {
	if m != nil {
		return m.BroadcasterRunning
	}
	return false
}

func (m *ServerStatus) GetGSLoggerRunning() bool {
	if m != nil {
		return m.GSLoggerRunning
	}
	return false
}

func (m *ServerStatus) GetOpenPorts() []*OpenPort {
	if m != nil {
		return m.OpenPorts
	}
	return nil
}

type OpenPort struct {
	Port    int32 `protobuf:"varint,1,opt,name=Port" json:"Port,omitempty"`
	Serving bool  `protobuf:"varint,2,opt,name=Serving" json:"Serving,omitempty"`
}

func (m *OpenPort) Reset()                    { *m = OpenPort{} }
func (m *OpenPort) String() string            { return proto1.CompactTextString(m) }
func (*OpenPort) ProtoMessage()               {}
func (*OpenPort) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *OpenPort) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *OpenPort) GetServing() bool {
	if m != nil {
		return m.Serving
	}
	return false
}

type Pong struct {
	Status *ServerStatus `protobuf:"bytes,1,opt,name=Status" json:"Status,omitempty"`
}

func (m *Pong) Reset()                    { *m = Pong{} }
func (m *Pong) String() string            { return proto1.CompactTextString(m) }
func (*Pong) ProtoMessage()               {}
func (*Pong) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Pong) GetStatus() *ServerStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type Ping struct {
}

func (m *Ping) Reset()                    { *m = Ping{} }
func (m *Ping) String() string            { return proto1.CompactTextString(m) }
func (*Ping) ProtoMessage()               {}
func (*Ping) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type DataBundle struct {
	Parameters []*Params `protobuf:"bytes,1,rep,name=Parameters" json:"Parameters,omitempty"`
}

func (m *DataBundle) Reset()                    { *m = DataBundle{} }
func (m *DataBundle) String() string            { return proto1.CompactTextString(m) }
func (*DataBundle) ProtoMessage()               {}
func (*DataBundle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DataBundle) GetParameters() []*Params {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type DataPacket struct {
	PacketName string    `protobuf:"bytes,1,opt,name=PacketName" json:"PacketName,omitempty"`
	PacketType int32     `protobuf:"varint,2,opt,name=PacketType" json:"PacketType,omitempty"`
	RxTime     int64     `protobuf:"varint,3,opt,name=RxTime" json:"RxTime,omitempty"`
	Parameters []*Params `protobuf:"bytes,4,rep,name=Parameters" json:"Parameters,omitempty"`
}

func (m *DataPacket) Reset()                    { *m = DataPacket{} }
func (m *DataPacket) String() string            { return proto1.CompactTextString(m) }
func (*DataPacket) ProtoMessage()               {}
func (*DataPacket) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DataPacket) GetPacketName() string {
	if m != nil {
		return m.PacketName
	}
	return ""
}

func (m *DataPacket) GetPacketType() int32 {
	if m != nil {
		return m.PacketType
	}
	return 0
}

func (m *DataPacket) GetRxTime() int64 {
	if m != nil {
		return m.RxTime
	}
	return 0
}

func (m *DataPacket) GetParameters() []*Params {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type Params struct {
	PacketName string `protobuf:"bytes,1,opt,name=PacketName" json:"PacketName,omitempty"`
	ParamName  string `protobuf:"bytes,2,opt,name=ParamName" json:"ParamName,omitempty"`
	RxTime     int64  `protobuf:"varint,3,opt,name=RxTime" json:"RxTime,omitempty"`
	Value      *Value `protobuf:"bytes,4,opt,name=Value" json:"Value,omitempty"`
	Units      string `protobuf:"bytes,5,opt,name=Units" json:"Units,omitempty"`
}

func (m *Params) Reset()                    { *m = Params{} }
func (m *Params) String() string            { return proto1.CompactTextString(m) }
func (*Params) ProtoMessage()               {}
func (*Params) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Params) GetPacketName() string {
	if m != nil {
		return m.PacketName
	}
	return ""
}

func (m *Params) GetParamName() string {
	if m != nil {
		return m.ParamName
	}
	return ""
}

func (m *Params) GetRxTime() int64 {
	if m != nil {
		return m.RxTime
	}
	return 0
}

func (m *Params) GetValue() *Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Params) GetUnits() string {
	if m != nil {
		return m.Units
	}
	return ""
}

type Value struct {
	Index       int32   `protobuf:"varint,1,opt,name=Index" json:"Index,omitempty"`
	Int64Value  int64   `protobuf:"varint,2,opt,name=Int64Value" json:"Int64Value,omitempty"`
	Uint64Value uint64  `protobuf:"varint,3,opt,name=Uint64Value" json:"Uint64Value,omitempty"`
	DoubleValue float64 `protobuf:"fixed64,4,opt,name=DoubleValue" json:"DoubleValue,omitempty"`
}

func (m *Value) Reset()                    { *m = Value{} }
func (m *Value) String() string            { return proto1.CompactTextString(m) }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Value) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Value) GetInt64Value() int64 {
	if m != nil {
		return m.Int64Value
	}
	return 0
}

func (m *Value) GetUint64Value() uint64 {
	if m != nil {
		return m.Uint64Value
	}
	return 0
}

func (m *Value) GetDoubleValue() float64 {
	if m != nil {
		return m.DoubleValue
	}
	return 0
}

func init() {
	proto1.RegisterType((*StreamRequest)(nil), "proto.StreamRequest")
	proto1.RegisterType((*ServerControl)(nil), "proto.ServerControl")
	proto1.RegisterType((*StartLogCommand)(nil), "proto.StartLogCommand")
	proto1.RegisterType((*Command)(nil), "proto.Command")
	proto1.RegisterType((*Ack)(nil), "proto.Ack")
	proto1.RegisterType((*ServerStatus)(nil), "proto.ServerStatus")
	proto1.RegisterType((*OpenPort)(nil), "proto.OpenPort")
	proto1.RegisterType((*Pong)(nil), "proto.Pong")
	proto1.RegisterType((*Ping)(nil), "proto.Ping")
	proto1.RegisterType((*DataBundle)(nil), "proto.DataBundle")
	proto1.RegisterType((*DataPacket)(nil), "proto.DataPacket")
	proto1.RegisterType((*Params)(nil), "proto.Params")
	proto1.RegisterType((*Value)(nil), "proto.Value")
	proto1.RegisterEnum("proto.ServerControl_CommandEnum", ServerControl_CommandEnum_name, ServerControl_CommandEnum_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GroundStationService service

type GroundStationServiceClient interface {
	StreamPackets(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (GroundStationService_StreamPacketsClient, error)
	SendCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Ack, error)
	Ping(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error)
	ControlServer(ctx context.Context, in *ServerControl, opts ...grpc.CallOption) (*Ack, error)
}

type groundStationServiceClient struct {
	cc *grpc.ClientConn
}

func NewGroundStationServiceClient(cc *grpc.ClientConn) GroundStationServiceClient {
	return &groundStationServiceClient{cc}
}

func (c *groundStationServiceClient) StreamPackets(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (GroundStationService_StreamPacketsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GroundStationService_serviceDesc.Streams[0], c.cc, "/proto.GroundStationService/streamPackets", opts...)
	if err != nil {
		return nil, err
	}
	x := &groundStationServiceStreamPacketsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GroundStationService_StreamPacketsClient interface {
	Recv() (*DataBundle, error)
	grpc.ClientStream
}

type groundStationServiceStreamPacketsClient struct {
	grpc.ClientStream
}

func (x *groundStationServiceStreamPacketsClient) Recv() (*DataBundle, error) {
	m := new(DataBundle)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *groundStationServiceClient) SendCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := grpc.Invoke(ctx, "/proto.GroundStationService/sendCommand", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groundStationServiceClient) Ping(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := grpc.Invoke(ctx, "/proto.GroundStationService/ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groundStationServiceClient) ControlServer(ctx context.Context, in *ServerControl, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := grpc.Invoke(ctx, "/proto.GroundStationService/controlServer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GroundStationService service

type GroundStationServiceServer interface {
	StreamPackets(*StreamRequest, GroundStationService_StreamPacketsServer) error
	SendCommand(context.Context, *Command) (*Ack, error)
	Ping(context.Context, *Ping) (*Pong, error)
	ControlServer(context.Context, *ServerControl) (*Ack, error)
}

func RegisterGroundStationServiceServer(s *grpc.Server, srv GroundStationServiceServer) {
	s.RegisterService(&_GroundStationService_serviceDesc, srv)
}

func _GroundStationService_StreamPackets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GroundStationServiceServer).StreamPackets(m, &groundStationServiceStreamPacketsServer{stream})
}

type GroundStationService_StreamPacketsServer interface {
	Send(*DataBundle) error
	grpc.ServerStream
}

type groundStationServiceStreamPacketsServer struct {
	grpc.ServerStream
}

func (x *groundStationServiceStreamPacketsServer) Send(m *DataBundle) error {
	return x.ServerStream.SendMsg(m)
}

func _GroundStationService_SendCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroundStationServiceServer).SendCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GroundStationService/SendCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroundStationServiceServer).SendCommand(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroundStationService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroundStationServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GroundStationService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroundStationServiceServer).Ping(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroundStationService_ControlServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerControl)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroundStationServiceServer).ControlServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GroundStationService/ControlServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroundStationServiceServer).ControlServer(ctx, req.(*ServerControl))
	}
	return interceptor(ctx, in, info, handler)
}

var _GroundStationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GroundStationService",
	HandlerType: (*GroundStationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sendCommand",
			Handler:    _GroundStationService_SendCommand_Handler,
		},
		{
			MethodName: "ping",
			Handler:    _GroundStationService_Ping_Handler,
		},
		{
			MethodName: "controlServer",
			Handler:    _GroundStationService_ControlServer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "streamPackets",
			Handler:       _GroundStationService_StreamPackets_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "groundstation.proto",
}

func init() { proto1.RegisterFile("groundstation.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 680 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0x6e, 0xda, 0xa6, 0xdb, 0xae, 0xeb, 0xda, 0xdd, 0x0a, 0x94, 0x09, 0xa1, 0xca, 0x12, 0x52,
	0xa5, 0xb1, 0x0a, 0x75, 0x08, 0x4d, 0xc0, 0x0b, 0xdb, 0xd0, 0x34, 0x69, 0x8c, 0xe2, 0x6c, 0xbc,
	0x7b, 0x8d, 0x15, 0x45, 0x6b, 0xec, 0x90, 0x38, 0x68, 0x3c, 0xf2, 0x1b, 0x90, 0x10, 0xbf, 0x8d,
	0x67, 0xfe, 0x07, 0xc8, 0x76, 0xd2, 0xa4, 0x1b, 0xd3, 0x9e, 0xe2, 0xbb, 0xef, 0xbb, 0xf3, 0xe7,
	0xbb, 0xcb, 0xc1, 0x56, 0x90, 0xc8, 0x4c, 0xf8, 0xa9, 0x62, 0x2a, 0x94, 0x62, 0x1c, 0x27, 0x52,
	0x49, 0x74, 0xcd, 0x87, 0x74, 0xa1, 0xe3, 0xa9, 0x84, 0xb3, 0x88, 0xf2, 0x2f, 0x19, 0x4f, 0x15,
	0xf9, 0xe3, 0x40, 0xc7, 0xe3, 0xc9, 0x57, 0x9e, 0x1c, 0x4a, 0xa1, 0x12, 0x39, 0xc7, 0xd7, 0xb0,
	0x72, 0x28, 0xa3, 0x88, 0x09, 0x7f, 0xe0, 0x0c, 0x9d, 0xd1, 0xc6, 0x64, 0x68, 0x53, 0x8c, 0x97,
	0x68, 0xe3, 0x9c, 0xf3, 0x5e, 0x64, 0x11, 0x2d, 0x02, 0xc8, 0x4f, 0x07, 0xda, 0x15, 0x00, 0xb7,
	0xa0, 0x7b, 0x2a, 0x03, 0x1d, 0x18, 0xce, 0xb8, 0xa7, 0x58, 0xa2, 0x7a, 0x35, 0x44, 0xd8, 0xa8,
	0x3a, 0x65, 0xdc, 0x73, 0xf0, 0x31, 0x3c, 0x38, 0x62, 0x8a, 0x79, 0x4a, 0x26, 0xfc, 0x03, 0x13,
	0x2c, 0xe0, 0x89, 0xa5, 0xd7, 0x71, 0x00, 0xfd, 0xdb, 0x90, 0x8c, 0x7b, 0x0d, 0xec, 0x43, 0xef,
	0x20, 0x91, 0xcc, 0x9f, 0xb1, 0x54, 0x15, 0xfc, 0xa6, 0xbe, 0x73, 0xc9, 0x2b, 0xe3, 0x9e, 0x4b,
	0x9e, 0x41, 0xd7, 0xe0, 0xa7, 0x32, 0xc8, 0xf5, 0x21, 0x42, 0x53, 0xb0, 0x88, 0x9b, 0x47, 0xae,
	0x51, 0x73, 0x26, 0x9f, 0x16, 0x6f, 0xd7, 0xf0, 0x99, 0xf4, 0x17, 0xb0, 0x3e, 0xe3, 0x53, 0x80,
	0x29, 0x9b, 0x5d, 0x71, 0x75, 0xfe, 0x2d, 0xe6, 0x83, 0xfa, 0xd0, 0x19, 0xb9, 0xb4, 0xe2, 0xd1,
	0x31, 0x5a, 0xea, 0xa0, 0x31, 0x6c, 0x8c, 0x5c, 0x6a, 0xce, 0xc4, 0x85, 0xc6, 0xbb, 0xd9, 0x15,
	0xf9, 0xeb, 0xc0, 0xba, 0x2d, 0xa0, 0xa7, 0x98, 0xca, 0x52, 0xdc, 0x87, 0x47, 0x37, 0x9f, 0x45,
	0x33, 0x21, 0x42, 0x11, 0x98, 0x2b, 0x57, 0xe9, 0x5d, 0x30, 0x3e, 0x87, 0xcd, 0x63, 0x3a, 0x3d,
	0xb4, 0xd9, 0x8a, 0x98, 0xba, 0x89, 0xb9, 0x0d, 0xe0, 0x18, 0xb0, 0x52, 0x8e, 0x82, 0xde, 0x30,
	0xf4, 0xff, 0x20, 0x38, 0x82, 0xee, 0xb1, 0x77, 0x2a, 0x83, 0x8a, 0x9e, 0xa6, 0x21, 0xdf, 0x74,
	0xe3, 0x2e, 0xac, 0x7d, 0x8c, 0xb9, 0x98, 0xca, 0x44, 0xa5, 0x03, 0x77, 0xd8, 0x18, 0xb5, 0x27,
	0xdd, 0x7c, 0x54, 0x0a, 0x3f, 0x2d, 0x19, 0x64, 0x1f, 0x56, 0x0b, 0x43, 0x17, 0x4a, 0x7f, 0xcd,
	0x4b, 0x5d, 0x6a, 0xce, 0x38, 0x80, 0x15, 0x33, 0x13, 0x8b, 0xc7, 0x14, 0x26, 0xd9, 0xd3, 0x6c,
	0x11, 0xe0, 0x0e, 0xb4, 0x6c, 0xf1, 0x4c, 0x5c, 0x7b, 0xb2, 0xb5, 0x34, 0x98, 0x16, 0xa2, 0x39,
	0x85, 0xb4, 0xa0, 0x39, 0xd5, 0xc1, 0x6f, 0x00, 0x74, 0x21, 0x0f, 0x32, 0xe1, 0xcf, 0x39, 0xee,
	0xea, 0x0e, 0x26, 0x2c, 0xe2, 0x8a, 0x27, 0x3a, 0x8d, 0x16, 0xdd, 0xc9, 0xd3, 0x18, 0x20, 0xa5,
	0x15, 0x02, 0xf9, 0xe1, 0xd8, 0x68, 0xdb, 0xe3, 0xb2, 0xff, 0x67, 0xe5, 0xe0, 0x54, 0x3c, 0xf7,
	0xce, 0xc7, 0x43, 0x68, 0xd1, 0xeb, 0xf3, 0x30, 0xe2, 0xa6, 0xfe, 0x0d, 0x9a, 0x5b, 0x37, 0x54,
	0x35, 0xef, 0x53, 0xf5, 0xcb, 0x81, 0x96, 0x75, 0xdf, 0xab, 0xe8, 0x09, 0xac, 0x19, 0xa6, 0x81,
	0xeb, 0x06, 0x2e, 0x1d, 0x77, 0xea, 0x21, 0xe0, 0x7e, 0x66, 0xf3, 0x8c, 0x9b, 0xce, 0xb7, 0x27,
	0xeb, 0xb9, 0x14, 0xe3, 0xa3, 0x16, 0xc2, 0x3e, 0xb8, 0x17, 0x22, 0x34, 0x9d, 0xd7, 0x59, 0xad,
	0x41, 0xbe, 0x3b, 0x50, 0xe2, 0x27, 0xc2, 0xe7, 0xd7, 0x79, 0x8f, 0xad, 0xa1, 0xf5, 0x9e, 0x08,
	0xf5, 0xea, 0xa5, 0x4d, 0x5f, 0x37, 0xb7, 0x56, 0x3c, 0x38, 0x84, 0xf6, 0x45, 0x58, 0x12, 0xb4,
	0xac, 0x26, 0xad, 0xba, 0x34, 0xe3, 0x48, 0x66, 0x97, 0x73, 0x5e, 0x2a, 0x74, 0x68, 0xd5, 0x35,
	0xf9, 0xed, 0x40, 0xff, 0xd8, 0xac, 0x40, 0xcf, 0xae, 0xc0, 0x7c, 0xd5, 0xe0, 0x5b, 0xe8, 0xa4,
	0x66, 0xf9, 0xd9, 0x02, 0xa5, 0xd8, 0x2f, 0x06, 0xa8, 0xba, 0x12, 0xb7, 0x37, 0x73, 0x6f, 0x39,
	0x36, 0xa4, 0xf6, 0xc2, 0xc1, 0x1d, 0x68, 0xa7, 0x5c, 0xf8, 0xc5, 0x7e, 0xd8, 0xc8, 0x59, 0xb9,
	0xbd, 0x0d, 0xb9, 0xad, 0x7f, 0xf6, 0x1a, 0x12, 0x68, 0xc6, 0xfa, 0x1f, 0x69, 0x17, 0x5d, 0x0c,
	0x45, 0xb0, 0xbd, 0x30, 0xa4, 0x08, 0x48, 0x0d, 0xf7, 0xa0, 0x33, 0xb3, 0xcb, 0xd4, 0x0e, 0x70,
	0x29, 0xa7, 0xba, 0x68, 0x97, 0x13, 0x5f, 0xb6, 0x8c, 0xb1, 0xf7, 0x2f, 0x00, 0x00, 0xff, 0xff,
	0x9c, 0x20, 0xc6, 0x83, 0xe5, 0x05, 0x00, 0x00,
}
