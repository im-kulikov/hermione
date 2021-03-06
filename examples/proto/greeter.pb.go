// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/im-kulikov/hermione/examples/proto/greeter.proto

/*
Package greeter is a generated protocol buffer package.

It is generated from these files:
	github.com/im-kulikov/hermione/examples/proto/greeter.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
*/
package greeter

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

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Greeting string `protobuf:"bytes,2,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "greeter.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "greeter.HelloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := grpc.Invoke(ctx, "/greeter.Greeter/Hello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greeter.Greeter/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greeter.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Greeter_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/im-kulikov/hermione/examples/proto/greeter.proto",
}

func init() {
	proto.RegisterFile("github.com/im-kulikov/hermione/examples/proto/greeter.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x2e, 0xaa, 0x2c, 0x28, 0xc9, 0x2f, 0x48, 0xac,
	0xd4, 0x4d, 0x49, 0x2d, 0xd3, 0x4f, 0x4f, 0xcd, 0x2d, 0x2e, 0xc9, 0xcf, 0x4b, 0xd5, 0x4f, 0xad,
	0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0x2f, 0x4a,
	0x4d, 0x2d, 0x49, 0x2d, 0xd2, 0x03, 0xf3, 0x84, 0xd8, 0xa1, 0x5c, 0x25, 0x25, 0x2e, 0x1e, 0x8f,
	0xd4, 0x9c, 0x9c, 0xfc, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xbc,
	0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49, 0x9b, 0x8b, 0x17,
	0xaa, 0xa6, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x8a, 0x8b, 0x03, 0xac, 0x3f, 0x33, 0x2f,
	0x5d, 0x82, 0x09, 0xac, 0x10, 0xce, 0x37, 0x72, 0xe6, 0x62, 0x77, 0x87, 0x98, 0x2d, 0x64, 0xc1,
	0xc5, 0x0a, 0xd6, 0x27, 0x24, 0xaa, 0x07, 0xb3, 0x1d, 0xd9, 0x2e, 0x29, 0x31, 0x74, 0x61, 0x88,
	0xf1, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x57, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x10, 0x19,
	0x90, 0x6c, 0xe7, 0x00, 0x00, 0x00,
}
