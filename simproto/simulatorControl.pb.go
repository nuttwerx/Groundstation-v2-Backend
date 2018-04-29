// Code generated by protoc-gen-go. DO NOT EDIT.
// source: simulatorControl.proto

/*
Package simproto is a generated protocol buffer package.

It is generated from these files:
	simulatorControl.proto

It has these top-level messages:
	SimCommand
	Ack
*/
package simproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SimCommand_SimCommandEnum int32

const (
	SimCommand_StartSimulator  SimCommand_SimCommandEnum = 0
	SimCommand_PauseSimulator  SimCommand_SimCommandEnum = 1
	SimCommand_ResumeSimulator SimCommand_SimCommandEnum = 2
	SimCommand_StopSimulator   SimCommand_SimCommandEnum = 3
)

var SimCommand_SimCommandEnum_name = map[int32]string{
	0: "StartSimulator",
	1: "PauseSimulator",
	2: "ResumeSimulator",
	3: "StopSimulator",
}
var SimCommand_SimCommandEnum_value = map[string]int32{
	"StartSimulator":  0,
	"PauseSimulator":  1,
	"ResumeSimulator": 2,
	"StopSimulator":   3,
}

func (x SimCommand_SimCommandEnum) String() string {
	return proto.EnumName(SimCommand_SimCommandEnum_name, int32(x))
}
func (SimCommand_SimCommandEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type SimCommand struct {
	Command SimCommand_SimCommandEnum `protobuf:"varint,1,opt,name=Command,enum=simproto.SimCommand_SimCommandEnum" json:"Command,omitempty"`
}

func (m *SimCommand) Reset()                    { *m = SimCommand{} }
func (m *SimCommand) String() string            { return proto.CompactTextString(m) }
func (*SimCommand) ProtoMessage()               {}
func (*SimCommand) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SimCommand) GetCommand() SimCommand_SimCommandEnum {
	if m != nil {
		return m.Command
	}
	return SimCommand_StartSimulator
}

type Ack struct {
	CommandExecuted bool `protobuf:"varint,1,opt,name=CommandExecuted" json:"CommandExecuted,omitempty"`
}

func (m *Ack) Reset()                    { *m = Ack{} }
func (m *Ack) String() string            { return proto.CompactTextString(m) }
func (*Ack) ProtoMessage()               {}
func (*Ack) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ack) GetCommandExecuted() bool {
	if m != nil {
		return m.CommandExecuted
	}
	return false
}

func init() {
	proto.RegisterType((*SimCommand)(nil), "simproto.SimCommand")
	proto.RegisterType((*Ack)(nil), "simproto.Ack")
	proto.RegisterEnum("simproto.SimCommand_SimCommandEnum", SimCommand_SimCommandEnum_name, SimCommand_SimCommandEnum_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SimControlService service

type SimControlServiceClient interface {
	SendSimCommand(ctx context.Context, in *SimCommand, opts ...grpc.CallOption) (*Ack, error)
}

type simControlServiceClient struct {
	cc *grpc.ClientConn
}

func NewSimControlServiceClient(cc *grpc.ClientConn) SimControlServiceClient {
	return &simControlServiceClient{cc}
}

func (c *simControlServiceClient) SendSimCommand(ctx context.Context, in *SimCommand, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := grpc.Invoke(ctx, "/simproto.SimControlService/SendSimCommand", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SimControlService service

type SimControlServiceServer interface {
	SendSimCommand(context.Context, *SimCommand) (*Ack, error)
}

func RegisterSimControlServiceServer(s *grpc.Server, srv SimControlServiceServer) {
	s.RegisterService(&_SimControlService_serviceDesc, srv)
}

func _SimControlService_SendSimCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimControlServiceServer).SendSimCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simproto.SimControlService/SendSimCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimControlServiceServer).SendSimCommand(ctx, req.(*SimCommand))
	}
	return interceptor(ctx, in, info, handler)
}

var _SimControlService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "simproto.SimControlService",
	HandlerType: (*SimControlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSimCommand",
			Handler:    _SimControlService_SendSimCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "simulatorControl.proto",
}

func init() { proto.RegisterFile("simulatorControl.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0xce, 0xcc, 0x2d,
	0xcd, 0x49, 0x2c, 0xc9, 0x2f, 0x72, 0xce, 0xcf, 0x2b, 0x29, 0xca, 0xcf, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x28, 0xce, 0xcc, 0x05, 0xb3, 0x94, 0xd6, 0x32, 0x72, 0x71, 0x05, 0x67,
	0xe6, 0x3a, 0xe7, 0xe7, 0xe6, 0x26, 0xe6, 0xa5, 0x08, 0xd9, 0x72, 0xb1, 0x43, 0x99, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x7c, 0x46, 0xca, 0x7a, 0x30, 0xa5, 0x7a, 0x08, 0x65, 0x48, 0x4c, 0xd7, 0xbc,
	0xd2, 0xdc, 0x20, 0x98, 0x1e, 0xa5, 0x04, 0x2e, 0x3e, 0x54, 0x29, 0x21, 0x21, 0x2e, 0xbe, 0xe0,
	0x92, 0xc4, 0xa2, 0x92, 0x60, 0x98, 0x43, 0x04, 0x18, 0x40, 0x62, 0x01, 0x89, 0xa5, 0xc5, 0xa9,
	0x08, 0x31, 0x46, 0x21, 0x61, 0x2e, 0xfe, 0xa0, 0xd4, 0xe2, 0xd2, 0x5c, 0x24, 0x41, 0x26, 0x21,
	0x41, 0x2e, 0xde, 0xe0, 0x92, 0xfc, 0x02, 0x84, 0x10, 0xb3, 0x92, 0x3e, 0x17, 0xb3, 0x63, 0x72,
	0xb6, 0x90, 0x06, 0x17, 0x3f, 0xcc, 0x96, 0x8a, 0xd4, 0xe4, 0xd2, 0x92, 0x54, 0x88, 0x7b, 0x39,
	0x82, 0xd0, 0x85, 0x8d, 0x7c, 0xb8, 0x04, 0xc1, 0x4e, 0x02, 0x7b, 0x3f, 0x38, 0xb5, 0xa8, 0x2c,
	0x33, 0x39, 0x55, 0xc8, 0x9c, 0x8b, 0x2f, 0x38, 0x35, 0x2f, 0x05, 0xc9, 0xe3, 0x22, 0xd8, 0xfc,
	0x29, 0xc5, 0x8b, 0x10, 0x75, 0x4c, 0xce, 0x56, 0x62, 0x48, 0x62, 0x03, 0x73, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x3f, 0x84, 0x89, 0xe9, 0x59, 0x01, 0x00, 0x00,
}
