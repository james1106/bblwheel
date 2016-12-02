// Code generated by protoc-gen-go.
// source: rpc.proto
// DO NOT EDIT!

package bblwheel

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

type Request struct {
	ID         int64             `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	ClintID    string            `protobuf:"bytes,2,opt,name=ClintID" json:"ClintID,omitempty"`
	ClientAddr string            `protobuf:"bytes,3,opt,name=ClientAddr" json:"ClientAddr,omitempty"`
	Timestamp  int64             `protobuf:"varint,4,opt,name=Timestamp" json:"Timestamp,omitempty"`
	ForwardFor []string          `protobuf:"bytes,5,rep,name=ForwardFor" json:"ForwardFor,omitempty"`
	Path       string            `protobuf:"bytes,6,opt,name=Path" json:"Path,omitempty"`
	Token      string            `protobuf:"bytes,7,opt,name=Token" json:"Token,omitempty"`
	Content    []byte            `protobuf:"bytes,8,opt,name=Content,proto3" json:"Content,omitempty"`
	Params     map[string]string `protobuf:"bytes,9,rep,name=Params" json:"Params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Request) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

type Response struct {
	ID         int64             `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	ClientID   string            `protobuf:"bytes,2,opt,name=ClientID" json:"ClientID,omitempty"`
	Timestamp  int64             `protobuf:"varint,3,opt,name=Timestamp" json:"Timestamp,omitempty"`
	Status     int32             `protobuf:"varint,5,opt,name=Status" json:"Status,omitempty"`
	StatusText string            `protobuf:"bytes,6,opt,name=StatusText" json:"StatusText,omitempty"`
	NewToken   string            `protobuf:"bytes,7,opt,name=NewToken" json:"NewToken,omitempty"`
	Content    []byte            `protobuf:"bytes,8,opt,name=Content,proto3" json:"Content,omitempty"`
	Params     map[string]string `protobuf:"bytes,9,rep,name=Params" json:"Params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Response) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

type Message struct {
	ID         int64             `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Path       string            `protobuf:"bytes,2,opt,name=Path" json:"Path,omitempty"`
	Timestamp  int64             `protobuf:"varint,3,opt,name=Timestamp" json:"Timestamp,omitempty"`
	ForwardFor []string          `protobuf:"bytes,4,rep,name=ForwardFor" json:"ForwardFor,omitempty"`
	Token      string            `protobuf:"bytes,5,opt,name=Token" json:"Token,omitempty"`
	Content    []byte            `protobuf:"bytes,6,opt,name=Content,proto3" json:"Content,omitempty"`
	Params     map[string]string `protobuf:"bytes,7,rep,name=Params" json:"Params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Message) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "bblwheel.Request")
	proto.RegisterType((*Response)(nil), "bblwheel.Response")
	proto.RegisterType((*Message)(nil), "bblwheel.Message")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Rpc service

type RpcClient interface {
	Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Channel(ctx context.Context, opts ...grpc.CallOption) (Rpc_ChannelClient, error)
}

type rpcClient struct {
	cc *grpc.ClientConn
}

func NewRpcClient(cc *grpc.ClientConn) RpcClient {
	return &rpcClient{cc}
}

func (c *rpcClient) Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/bblwheel.Rpc/Call", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) Channel(ctx context.Context, opts ...grpc.CallOption) (Rpc_ChannelClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Rpc_serviceDesc.Streams[0], c.cc, "/bblwheel.Rpc/Channel", opts...)
	if err != nil {
		return nil, err
	}
	x := &rpcChannelClient{stream}
	return x, nil
}

type Rpc_ChannelClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type rpcChannelClient struct {
	grpc.ClientStream
}

func (x *rpcChannelClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rpcChannelClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Rpc service

type RpcServer interface {
	Call(context.Context, *Request) (*Response, error)
	Channel(Rpc_ChannelServer) error
}

func RegisterRpcServer(s *grpc.Server, srv RpcServer) {
	s.RegisterService(&_Rpc_serviceDesc, srv)
}

func _Rpc_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bblwheel.Rpc/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Call(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_Channel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RpcServer).Channel(&rpcChannelServer{stream})
}

type Rpc_ChannelServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type rpcChannelServer struct {
	grpc.ServerStream
}

func (x *rpcChannelServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rpcChannelServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Rpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bblwheel.Rpc",
	HandlerType: (*RpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Rpc_Call_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Channel",
			Handler:       _Rpc_Channel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x94, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0xe5, 0xa4, 0x49, 0x9a, 0x33, 0x84, 0xc0, 0x9a, 0x90, 0x55, 0x60, 0x8a, 0x76, 0x95,
	0x1b, 0x22, 0xb4, 0x09, 0x04, 0xdc, 0x41, 0xcb, 0x44, 0x2f, 0x40, 0x93, 0xe9, 0x0b, 0xb8, 0xed,
	0x11, 0xad, 0x96, 0x38, 0xc1, 0x76, 0x29, 0x7d, 0x08, 0xde, 0x81, 0x97, 0xe0, 0x19, 0x78, 0xad,
	0x29, 0xce, 0x9f, 0xa5, 0x8d, 0xb6, 0xab, 0xde, 0x9d, 0xef, 0xb3, 0x8f, 0x7b, 0xce, 0xaf, 0x9f,
	0x02, 0xa1, 0x2a, 0x16, 0x49, 0xa1, 0x72, 0x93, 0xd3, 0xe1, 0x7c, 0x9e, 0x6e, 0x57, 0x88, 0xe9,
	0xf9, 0x7f, 0x07, 0x02, 0x8e, 0x3f, 0x37, 0xa8, 0x0d, 0x7d, 0x0c, 0xce, 0x74, 0xc2, 0x48, 0x44,
	0x62, 0x97, 0x3b, 0xd3, 0x09, 0x65, 0x10, 0x8c, 0xd3, 0xb5, 0x34, 0xd3, 0x09, 0x73, 0x22, 0x12,
	0x87, 0xbc, 0x91, 0xf4, 0x0c, 0x60, 0x9c, 0xae, 0x51, 0x9a, 0x8f, 0xcb, 0xa5, 0x62, 0xae, 0x3d,
	0xec, 0x38, 0xf4, 0x05, 0x84, 0xb3, 0x75, 0x86, 0xda, 0x88, 0xac, 0x60, 0x03, 0xfb, 0xe0, 0x9d,
	0x51, 0x76, 0x5f, 0xe5, 0x6a, 0x2b, 0xd4, 0xf2, 0x2a, 0x57, 0xcc, 0x8b, 0xdc, 0xb2, 0xfb, 0xce,
	0xa1, 0x14, 0x06, 0xd7, 0xc2, 0xac, 0x98, 0x6f, 0xdf, 0xb5, 0x35, 0x3d, 0x05, 0x6f, 0x96, 0xdf,
	0xa0, 0x64, 0x81, 0x35, 0x2b, 0x61, 0x27, 0xcc, 0xa5, 0x41, 0x69, 0xd8, 0x30, 0x22, 0xf1, 0x23,
	0xde, 0x48, 0xfa, 0x06, 0xfc, 0x6b, 0xa1, 0x44, 0xa6, 0x59, 0x18, 0xb9, 0xf1, 0xc9, 0xc5, 0xcb,
	0xa4, 0x59, 0x39, 0xa9, 0xd7, 0x4d, 0xaa, 0xf3, 0xcf, 0xd2, 0xa8, 0x1d, 0xaf, 0x2f, 0x8f, 0xde,
	0xc3, 0x49, 0xc7, 0xa6, 0x4f, 0xc0, 0xbd, 0xc1, 0x9d, 0x45, 0x12, 0xf2, 0xb2, 0x2c, 0xe7, 0xf8,
	0x25, 0xd2, 0x0d, 0xd6, 0x44, 0x2a, 0xf1, 0xc1, 0x79, 0x47, 0xce, 0xff, 0x39, 0x30, 0xe4, 0xa8,
	0x8b, 0x5c, 0x6a, 0xec, 0xa1, 0x1c, 0xc1, 0xb0, 0xc2, 0xd3, 0xb2, 0x6c, 0xf5, 0x3e, 0x2c, 0xf7,
	0x10, 0xd6, 0x33, 0xf0, 0xbf, 0x1b, 0x61, 0x36, 0x9a, 0x79, 0x11, 0x89, 0x3d, 0x5e, 0xab, 0x12,
	0x62, 0x55, 0xcd, 0xf0, 0xb7, 0xa9, 0x51, 0x75, 0x9c, 0xf2, 0x17, 0xbf, 0xe1, 0xb6, 0xcb, 0xac,
	0xd5, 0x0f, 0x60, 0x7b, 0x7b, 0x80, 0xed, 0xac, 0x8b, 0xad, 0xda, 0xed, 0xd8, 0xdc, 0xfe, 0x38,
	0x10, 0x7c, 0x45, 0xad, 0xc5, 0x8f, 0x3e, 0xb6, 0x26, 0x09, 0x4e, 0x27, 0x09, 0x0f, 0xe3, 0xda,
	0xcf, 0xd6, 0xa0, 0x97, 0xad, 0x36, 0x47, 0xde, 0x3d, 0x39, 0xf2, 0xef, 0xcb, 0x51, 0x70, 0x98,
	0xa3, 0x7a, 0xe8, 0x23, 0xf3, 0xb8, 0x58, 0x83, 0xcb, 0x8b, 0x05, 0x7d, 0x05, 0x83, 0xb1, 0x48,
	0x53, 0xfa, 0xb4, 0x17, 0xdc, 0x11, 0xed, 0xff, 0x29, 0xf4, 0x12, 0x82, 0xf1, 0x4a, 0x48, 0x89,
	0x7b, 0x1d, 0xf5, 0x88, 0xa3, 0xbe, 0x15, 0x93, 0xd7, 0xe4, 0xd3, 0x73, 0x38, 0x5d, 0xe4, 0x59,
	0xb2, 0xc2, 0x4c, 0xec, 0x36, 0xb2, 0xbd, 0xf3, 0x85, 0xfc, 0x25, 0x64, 0xee, 0xdb, 0x4f, 0xc5,
	0xe5, 0x6d, 0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0x3b, 0x00, 0xe5, 0x37, 0x04, 0x00, 0x00,
}