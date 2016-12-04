// Code generated by protoc-gen-go.
// source: bblwheel.proto
// DO NOT EDIT!

/*
Package bblwheel is a generated protocol buffer package.

It is generated from these files:
	bblwheel.proto
	rpc.proto

It has these top-level messages:
	Void
	RegisterResult
	Service
	ConfigEntry
	Config
	UpdateConfigReq
	Stats
	Event
	LookupConfigReq
	LookupConfigResp
	LookupServiceReq
	LookupServiceResp
	Request
	Response
	Message
*/
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Service_Status int32

const (
	Service_INIT        Service_Status = 0
	Service_ONLINE      Service_Status = 1
	Service_MAINTENANCE Service_Status = 2
	Service_OFFLINE     Service_Status = 3
	Service_FAULT       Service_Status = 4
	Service_UNAUTHORIZE Service_Status = 5
)

var Service_Status_name = map[int32]string{
	0: "INIT",
	1: "ONLINE",
	2: "MAINTENANCE",
	3: "OFFLINE",
	4: "FAULT",
	5: "UNAUTHORIZE",
}
var Service_Status_value = map[string]int32{
	"INIT":        0,
	"ONLINE":      1,
	"MAINTENANCE": 2,
	"OFFLINE":     3,
	"FAULT":       4,
	"UNAUTHORIZE": 5,
}

func (x Service_Status) String() string {
	return proto.EnumName(Service_Status_name, int32(x))
}
func (Service_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type Event_EventType int32

const (
	Event_DISCOVERY       Event_EventType = 0
	Event_CONFIGUPDATE    Event_EventType = 1
	Event_REGISTER_RESULT Event_EventType = 2
	Event_KEEPALIVE       Event_EventType = 3
	Event_CONTROL         Event_EventType = 4
	Event_EXEC            Event_EventType = 5
)

var Event_EventType_name = map[int32]string{
	0: "DISCOVERY",
	1: "CONFIGUPDATE",
	2: "REGISTER_RESULT",
	3: "KEEPALIVE",
	4: "CONTROL",
	5: "EXEC",
}
var Event_EventType_value = map[string]int32{
	"DISCOVERY":       0,
	"CONFIGUPDATE":    1,
	"REGISTER_RESULT": 2,
	"KEEPALIVE":       3,
	"CONTROL":         4,
	"EXEC":            5,
}

func (x Event_EventType) String() string {
	return proto.EnumName(Event_EventType_name, int32(x))
}
func (Event_EventType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

type Void struct {
}

func (m *Void) Reset()                    { *m = Void{} }
func (m *Void) String() string            { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()               {}
func (*Void) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RegisterResult struct {
	Desc    string             `protobuf:"bytes,1,opt,name=Desc" json:"Desc,omitempty"`
	Service []*Service         `protobuf:"bytes,2,rep,name=Service" json:"Service,omitempty"`
	Configs map[string]*Config `protobuf:"bytes,3,rep,name=Configs" json:"Configs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *RegisterResult) Reset()                    { *m = RegisterResult{} }
func (m *RegisterResult) String() string            { return proto.CompactTextString(m) }
func (*RegisterResult) ProtoMessage()               {}
func (*RegisterResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterResult) GetService() []*Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *RegisterResult) GetConfigs() map[string]*Config {
	if m != nil {
		return m.Configs
	}
	return nil
}

type Service struct {
	ID                string         `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Name              string         `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Tags              []string       `protobuf:"bytes,3,rep,name=Tags" json:"Tags,omitempty"`
	Address           string         `protobuf:"bytes,4,opt,name=Address" json:"Address,omitempty"`
	DataCenter        string         `protobuf:"bytes,5,opt,name=DataCenter" json:"DataCenter,omitempty"`
	Node              string         `protobuf:"bytes,6,opt,name=Node" json:"Node,omitempty"`
	PID               string         `protobuf:"bytes,7,opt,name=PID" json:"PID,omitempty"`
	Weight            int32          `protobuf:"varint,8,opt,name=Weight" json:"Weight,omitempty"`
	Single            bool           `protobuf:"varint,9,opt,name=Single" json:"Single,omitempty"`
	DependentServices []string       `protobuf:"bytes,10,rep,name=DependentServices" json:"DependentServices,omitempty"`
	DependentConfigs  []string       `protobuf:"bytes,11,rep,name=DependentConfigs" json:"DependentConfigs,omitempty"`
	Status            Service_Status `protobuf:"varint,12,opt,name=status,enum=bblwheel.Service_Status" json:"status,omitempty"`
	Stats             *Stats         `protobuf:"bytes,13,opt,name=Stats" json:"Stats,omitempty"`
	Error             string         `protobuf:"bytes,14,opt,name=Error" json:"Error,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Service) GetStats() *Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type ConfigEntry struct {
	Key   string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=Value" json:"Value,omitempty"`
}

func (m *ConfigEntry) Reset()                    { *m = ConfigEntry{} }
func (m *ConfigEntry) String() string            { return proto.CompactTextString(m) }
func (*ConfigEntry) ProtoMessage()               {}
func (*ConfigEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Config struct {
	Items []*ConfigEntry `protobuf:"bytes,1,rep,name=Items" json:"Items,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Config) GetItems() []*ConfigEntry {
	if m != nil {
		return m.Items
	}
	return nil
}

type UpdateConfigReq struct {
	ServiceID   string  `protobuf:"bytes,1,opt,name=ServiceID" json:"ServiceID,omitempty"`
	ServiceName string  `protobuf:"bytes,2,opt,name=ServiceName" json:"ServiceName,omitempty"`
	Config      *Config `protobuf:"bytes,3,opt,name=config" json:"config,omitempty"`
}

func (m *UpdateConfigReq) Reset()                    { *m = UpdateConfigReq{} }
func (m *UpdateConfigReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateConfigReq) ProtoMessage()               {}
func (*UpdateConfigReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateConfigReq) GetConfig() *Config {
	if m != nil {
		return m.Config
	}
	return nil
}

type Stats struct {
	ServiceID      string            `protobuf:"bytes,1,opt,name=ServiceID" json:"ServiceID,omitempty"`
	ServiceName    string            `protobuf:"bytes,2,opt,name=ServiceName" json:"ServiceName,omitempty"`
	Count          int64             `protobuf:"varint,3,opt,name=Count" json:"Count,omitempty"`
	UpTime         int64             `protobuf:"varint,4,opt,name=UpTime" json:"UpTime,omitempty"`
	UsedMem        int64             `protobuf:"varint,5,opt,name=UsedMem" json:"UsedMem,omitempty"`
	FreeMem        int64             `protobuf:"varint,6,opt,name=FreeMem" json:"FreeMem,omitempty"`
	Threads        int64             `protobuf:"varint,7,opt,name=Threads" json:"Threads,omitempty"`
	AvgRespTime    int64             `protobuf:"varint,8,opt,name=AvgRespTime" json:"AvgRespTime,omitempty"`
	MinRespTime    int64             `protobuf:"varint,9,opt,name=MinRespTime" json:"MinRespTime,omitempty"`
	MaxRespTime    int64             `protobuf:"varint,10,opt,name=MaxRespTime" json:"MaxRespTime,omitempty"`
	LastActiveTime int64             `protobuf:"varint,11,opt,name=LastActiveTime" json:"LastActiveTime,omitempty"`
	Other          map[string]string `protobuf:"bytes,12,rep,name=Other" json:"Other,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Errors         string            `protobuf:"bytes,13,opt,name=Errors" json:"Errors,omitempty"`
}

func (m *Stats) Reset()                    { *m = Stats{} }
func (m *Stats) String() string            { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()               {}
func (*Stats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Stats) GetOther() map[string]string {
	if m != nil {
		return m.Other
	}
	return nil
}

type Event struct {
	Type    Event_EventType `protobuf:"varint,1,opt,name=Type,enum=bblwheel.Event_EventType" json:"Type,omitempty"`
	Service *Service        `protobuf:"bytes,2,opt,name=Service" json:"Service,omitempty"`
	Item    *ConfigEntry    `protobuf:"bytes,3,opt,name=Item" json:"Item,omitempty"`
	Command string          `protobuf:"bytes,4,opt,name=Command" json:"Command,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Event) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *Event) GetItem() *ConfigEntry {
	if m != nil {
		return m.Item
	}
	return nil
}

type LookupConfigReq struct {
	DependentConfigs []string `protobuf:"bytes,1,rep,name=DependentConfigs" json:"DependentConfigs,omitempty"`
}

func (m *LookupConfigReq) Reset()                    { *m = LookupConfigReq{} }
func (m *LookupConfigReq) String() string            { return proto.CompactTextString(m) }
func (*LookupConfigReq) ProtoMessage()               {}
func (*LookupConfigReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type LookupConfigResp struct {
	Configs map[string]*Config `protobuf:"bytes,1,rep,name=Configs" json:"Configs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *LookupConfigResp) Reset()                    { *m = LookupConfigResp{} }
func (m *LookupConfigResp) String() string            { return proto.CompactTextString(m) }
func (*LookupConfigResp) ProtoMessage()               {}
func (*LookupConfigResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *LookupConfigResp) GetConfigs() map[string]*Config {
	if m != nil {
		return m.Configs
	}
	return nil
}

type LookupServiceReq struct {
	ServiceID         string   `protobuf:"bytes,1,opt,name=ServiceID" json:"ServiceID,omitempty"`
	ServiceName       string   `protobuf:"bytes,2,opt,name=ServiceName" json:"ServiceName,omitempty"`
	DependentServices []string `protobuf:"bytes,3,rep,name=DependentServices" json:"DependentServices,omitempty"`
}

func (m *LookupServiceReq) Reset()                    { *m = LookupServiceReq{} }
func (m *LookupServiceReq) String() string            { return proto.CompactTextString(m) }
func (*LookupServiceReq) ProtoMessage()               {}
func (*LookupServiceReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type LookupServiceResp struct {
	Services []*Service `protobuf:"bytes,1,rep,name=Services" json:"Services,omitempty"`
}

func (m *LookupServiceResp) Reset()                    { *m = LookupServiceResp{} }
func (m *LookupServiceResp) String() string            { return proto.CompactTextString(m) }
func (*LookupServiceResp) ProtoMessage()               {}
func (*LookupServiceResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *LookupServiceResp) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

func init() {
	proto.RegisterType((*Void)(nil), "bblwheel.Void")
	proto.RegisterType((*RegisterResult)(nil), "bblwheel.RegisterResult")
	proto.RegisterType((*Service)(nil), "bblwheel.Service")
	proto.RegisterType((*ConfigEntry)(nil), "bblwheel.ConfigEntry")
	proto.RegisterType((*Config)(nil), "bblwheel.Config")
	proto.RegisterType((*UpdateConfigReq)(nil), "bblwheel.UpdateConfigReq")
	proto.RegisterType((*Stats)(nil), "bblwheel.Stats")
	proto.RegisterType((*Event)(nil), "bblwheel.Event")
	proto.RegisterType((*LookupConfigReq)(nil), "bblwheel.LookupConfigReq")
	proto.RegisterType((*LookupConfigResp)(nil), "bblwheel.LookupConfigResp")
	proto.RegisterType((*LookupServiceReq)(nil), "bblwheel.LookupServiceReq")
	proto.RegisterType((*LookupServiceResp)(nil), "bblwheel.LookupServiceResp")
	proto.RegisterEnum("bblwheel.Service_Status", Service_Status_name, Service_Status_value)
	proto.RegisterEnum("bblwheel.Event_EventType", Event_EventType_name, Event_EventType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BblWheel service

type BblWheelClient interface {
	LookupConfig(ctx context.Context, in *LookupConfigReq, opts ...grpc.CallOption) (*LookupConfigResp, error)
	LookupService(ctx context.Context, in *LookupServiceReq, opts ...grpc.CallOption) (*LookupServiceResp, error)
	UpdateConfig(ctx context.Context, in *UpdateConfigReq, opts ...grpc.CallOption) (*Void, error)
	Register(ctx context.Context, in *Service, opts ...grpc.CallOption) (*RegisterResult, error)
	Unregister(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Void, error)
	Events(ctx context.Context, opts ...grpc.CallOption) (BblWheel_EventsClient, error)
}

type bblWheelClient struct {
	cc *grpc.ClientConn
}

func NewBblWheelClient(cc *grpc.ClientConn) BblWheelClient {
	return &bblWheelClient{cc}
}

func (c *bblWheelClient) LookupConfig(ctx context.Context, in *LookupConfigReq, opts ...grpc.CallOption) (*LookupConfigResp, error) {
	out := new(LookupConfigResp)
	err := grpc.Invoke(ctx, "/bblwheel.BblWheel/LookupConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bblWheelClient) LookupService(ctx context.Context, in *LookupServiceReq, opts ...grpc.CallOption) (*LookupServiceResp, error) {
	out := new(LookupServiceResp)
	err := grpc.Invoke(ctx, "/bblwheel.BblWheel/LookupService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bblWheelClient) UpdateConfig(ctx context.Context, in *UpdateConfigReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := grpc.Invoke(ctx, "/bblwheel.BblWheel/UpdateConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bblWheelClient) Register(ctx context.Context, in *Service, opts ...grpc.CallOption) (*RegisterResult, error) {
	out := new(RegisterResult)
	err := grpc.Invoke(ctx, "/bblwheel.BblWheel/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bblWheelClient) Unregister(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := grpc.Invoke(ctx, "/bblwheel.BblWheel/Unregister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bblWheelClient) Events(ctx context.Context, opts ...grpc.CallOption) (BblWheel_EventsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BblWheel_serviceDesc.Streams[0], c.cc, "/bblwheel.BblWheel/Events", opts...)
	if err != nil {
		return nil, err
	}
	x := &bblWheelEventsClient{stream}
	return x, nil
}

type BblWheel_EventsClient interface {
	Send(*Event) error
	Recv() (*Event, error)
	grpc.ClientStream
}

type bblWheelEventsClient struct {
	grpc.ClientStream
}

func (x *bblWheelEventsClient) Send(m *Event) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bblWheelEventsClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for BblWheel service

type BblWheelServer interface {
	LookupConfig(context.Context, *LookupConfigReq) (*LookupConfigResp, error)
	LookupService(context.Context, *LookupServiceReq) (*LookupServiceResp, error)
	UpdateConfig(context.Context, *UpdateConfigReq) (*Void, error)
	Register(context.Context, *Service) (*RegisterResult, error)
	Unregister(context.Context, *Service) (*Void, error)
	Events(BblWheel_EventsServer) error
}

func RegisterBblWheelServer(s *grpc.Server, srv BblWheelServer) {
	s.RegisterService(&_BblWheel_serviceDesc, srv)
}

func _BblWheel_LookupConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BblWheelServer).LookupConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bblwheel.BblWheel/LookupConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BblWheelServer).LookupConfig(ctx, req.(*LookupConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BblWheel_LookupService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupServiceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BblWheelServer).LookupService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bblwheel.BblWheel/LookupService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BblWheelServer).LookupService(ctx, req.(*LookupServiceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BblWheel_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BblWheelServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bblwheel.BblWheel/UpdateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BblWheelServer).UpdateConfig(ctx, req.(*UpdateConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BblWheel_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BblWheelServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bblwheel.BblWheel/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BblWheelServer).Register(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

func _BblWheel_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BblWheelServer).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bblwheel.BblWheel/Unregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BblWheelServer).Unregister(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

func _BblWheel_Events_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BblWheelServer).Events(&bblWheelEventsServer{stream})
}

type BblWheel_EventsServer interface {
	Send(*Event) error
	Recv() (*Event, error)
	grpc.ServerStream
}

type bblWheelEventsServer struct {
	grpc.ServerStream
}

func (x *bblWheelEventsServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bblWheelEventsServer) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _BblWheel_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bblwheel.BblWheel",
	HandlerType: (*BblWheelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LookupConfig",
			Handler:    _BblWheel_LookupConfig_Handler,
		},
		{
			MethodName: "LookupService",
			Handler:    _BblWheel_LookupService_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _BblWheel_UpdateConfig_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _BblWheel_Register_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _BblWheel_Unregister_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Events",
			Handler:       _BblWheel_Events_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "bblwheel.proto",
}

func init() { proto.RegisterFile("bblwheel.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1056 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x56, 0xdf, 0x6e, 0xe3, 0xc4,
	0x17, 0x5e, 0xc7, 0xb1, 0x13, 0x9f, 0xa4, 0x89, 0x7b, 0x7e, 0xfd, 0x21, 0x6f, 0x16, 0xa1, 0xc8,
	0xd2, 0x2e, 0x81, 0x65, 0x43, 0x55, 0xb4, 0x62, 0x41, 0x42, 0x28, 0x4d, 0xdc, 0xad, 0xd5, 0x34,
	0xa9, 0x26, 0x4e, 0x77, 0xe1, 0x06, 0xb9, 0xf5, 0x90, 0x46, 0x9b, 0xd8, 0x59, 0xdb, 0x29, 0x54,
	0xdc, 0xf0, 0x0a, 0xbc, 0x01, 0x12, 0xcf, 0xc4, 0x3d, 0x8f, 0xc0, 0x05, 0x0f, 0x80, 0x66, 0x3c,
	0xfe, 0xd3, 0xb4, 0x81, 0x0b, 0x10, 0x37, 0xd6, 0x9c, 0xf3, 0x7d, 0x67, 0xe6, 0xcc, 0xcc, 0x77,
	0x8e, 0x07, 0x1a, 0x17, 0x17, 0x8b, 0xef, 0xae, 0x28, 0x5d, 0x74, 0x57, 0x61, 0x10, 0x07, 0x58,
	0x4d, 0x6d, 0x53, 0x85, 0xf2, 0x79, 0x30, 0xf7, 0xcc, 0xdf, 0x24, 0x68, 0x10, 0x3a, 0x9b, 0x47,
	0x31, 0x0d, 0x09, 0x8d, 0xd6, 0x8b, 0x18, 0x11, 0xca, 0x03, 0x1a, 0x5d, 0x1a, 0x52, 0x5b, 0xea,
	0x68, 0x84, 0x8f, 0xf1, 0x29, 0x54, 0x26, 0x34, 0xbc, 0x9e, 0x5f, 0x52, 0xa3, 0xd4, 0x96, 0x3b,
	0xb5, 0x83, 0xdd, 0x6e, 0x36, 0xb5, 0x00, 0x48, 0xca, 0xc0, 0x2f, 0xa1, 0xd2, 0x0f, 0xfc, 0x6f,
	0xe7, 0xb3, 0xc8, 0x90, 0x39, 0xf9, 0x71, 0x4e, 0xbe, 0xbd, 0x56, 0x57, 0xf0, 0x2c, 0x3f, 0x0e,
	0x6f, 0x48, 0x1a, 0xd5, 0x1a, 0x42, 0xbd, 0x08, 0xa0, 0x0e, 0xf2, 0x1b, 0x7a, 0x23, 0x12, 0x62,
	0x43, 0x7c, 0x02, 0xca, 0xb5, 0xbb, 0x58, 0xb3, 0x6c, 0xa4, 0x4e, 0xed, 0x40, 0xcf, 0x17, 0x48,
	0x02, 0x49, 0x02, 0x7f, 0x5e, 0x7a, 0x21, 0x99, 0x7f, 0xc8, 0x59, 0xf2, 0xd8, 0x80, 0x92, 0x3d,
	0x10, 0x13, 0x95, 0xec, 0x01, 0xdb, 0xeb, 0xc8, 0x5d, 0x26, 0xd3, 0x68, 0x84, 0x8f, 0x99, 0xcf,
	0x71, 0x45, 0xee, 0x1a, 0xe1, 0x63, 0x34, 0xa0, 0xd2, 0xf3, 0xbc, 0x90, 0x46, 0x91, 0x51, 0xe6,
	0xd4, 0xd4, 0xc4, 0xf7, 0x00, 0x06, 0x6e, 0xec, 0xf6, 0xa9, 0x1f, 0xd3, 0xd0, 0x50, 0x38, 0x58,
	0xf0, 0xf0, 0x15, 0x02, 0x8f, 0x1a, 0xaa, 0x58, 0x21, 0xf0, 0x28, 0xdb, 0xcf, 0x99, 0x3d, 0x30,
	0x2a, 0xc9, 0x7e, 0xce, 0xec, 0x01, 0xbe, 0x03, 0xea, 0x2b, 0x3a, 0x9f, 0x5d, 0xc5, 0x46, 0xb5,
	0x2d, 0x75, 0x14, 0x22, 0x2c, 0xe6, 0x9f, 0xcc, 0xfd, 0xd9, 0x82, 0x1a, 0x5a, 0x5b, 0xea, 0x54,
	0x89, 0xb0, 0xf0, 0x23, 0xd8, 0x1d, 0xd0, 0x15, 0xf5, 0x3d, 0xea, 0xc7, 0x62, 0x6f, 0x91, 0x01,
	0x3c, 0xe1, 0xbb, 0x00, 0x7e, 0x08, 0x7a, 0xe6, 0x4c, 0x6f, 0xa6, 0xc6, 0xc9, 0x77, 0xfc, 0xb8,
	0x0f, 0x6a, 0x14, 0xbb, 0xf1, 0x3a, 0x32, 0xea, 0x6d, 0xa9, 0xd3, 0x38, 0x30, 0xee, 0x5c, 0x74,
	0x77, 0xc2, 0x71, 0x22, 0x78, 0xf8, 0x18, 0x14, 0xe6, 0x89, 0x8c, 0x1d, 0x7e, 0x17, 0xcd, 0x42,
	0x00, 0x73, 0x93, 0x04, 0xc5, 0x3d, 0x50, 0xac, 0x30, 0x0c, 0x42, 0xa3, 0xc1, 0xb7, 0x9d, 0x18,
	0xe6, 0x6b, 0x50, 0x93, 0xe9, 0xb0, 0x0a, 0x65, 0x7b, 0x64, 0x3b, 0xfa, 0x03, 0x04, 0x50, 0xc7,
	0xa3, 0xa1, 0x3d, 0xb2, 0x74, 0x09, 0x9b, 0x50, 0x3b, 0xed, 0xd9, 0x23, 0xc7, 0x1a, 0xf5, 0x46,
	0x7d, 0x4b, 0x2f, 0x61, 0x0d, 0x2a, 0xe3, 0xa3, 0x23, 0x8e, 0xca, 0xa8, 0x81, 0x72, 0xd4, 0x9b,
	0x0e, 0x1d, 0xbd, 0xcc, 0x88, 0xd3, 0x51, 0x6f, 0xea, 0x1c, 0x8f, 0x89, 0xfd, 0xb5, 0xa5, 0x2b,
	0xe6, 0x73, 0xa8, 0x25, 0x7b, 0xca, 0x34, 0x74, 0x92, 0x6b, 0xe8, 0x84, 0xde, 0xb0, 0x84, 0xce,
	0x33, 0x0d, 0x69, 0x24, 0x31, 0xcc, 0xe7, 0xa0, 0x26, 0x61, 0xf8, 0x14, 0x14, 0x3b, 0xa6, 0xcb,
	0xc8, 0x90, 0xb8, 0x88, 0xff, 0xbf, 0xa9, 0xb1, 0x44, 0xb4, 0x09, 0xc7, 0xfc, 0x01, 0x9a, 0xd3,
	0x95, 0xe7, 0xc6, 0x54, 0xe8, 0x8f, 0xbe, 0xc5, 0x77, 0x41, 0x13, 0x27, 0x96, 0x49, 0x2e, 0x77,
	0x60, 0x1b, 0x6a, 0xc2, 0x28, 0x08, 0xb0, 0xe8, 0xc2, 0x0e, 0xa8, 0x97, 0x7c, 0x32, 0x43, 0xde,
	0x22, 0x72, 0x81, 0x9b, 0xbf, 0xca, 0xe2, 0x0a, 0xfe, 0xf1, 0x9a, 0x7b, 0xa0, 0xf4, 0x83, 0xb5,
	0x1f, 0xf3, 0x25, 0x65, 0x92, 0x18, 0x4c, 0x85, 0xd3, 0x95, 0x33, 0x5f, 0x52, 0x2e, 0x7e, 0x99,
	0x08, 0x8b, 0x55, 0xc5, 0x34, 0xa2, 0xde, 0x29, 0x5d, 0x72, 0xe1, 0xcb, 0x24, 0x35, 0x19, 0x72,
	0x14, 0x52, 0xca, 0x10, 0x35, 0x41, 0x84, 0xc9, 0x10, 0xe7, 0x2a, 0xa4, 0xae, 0x17, 0x71, 0xfd,
	0xcb, 0x24, 0x35, 0x59, 0x76, 0xbd, 0xeb, 0x19, 0xa1, 0x51, 0xb2, 0x54, 0x95, 0xa3, 0x45, 0x17,
	0x63, 0x9c, 0xce, 0xfd, 0x8c, 0xa1, 0x25, 0x8c, 0x82, 0x8b, 0x33, 0xdc, 0xef, 0x33, 0x06, 0x08,
	0x46, 0xee, 0xc2, 0x27, 0xd0, 0x18, 0xba, 0x51, 0xdc, 0xbb, 0x8c, 0xe7, 0xd7, 0x94, 0x93, 0x6a,
	0x9c, 0xb4, 0xe1, 0xc5, 0x7d, 0x50, 0xc6, 0xf1, 0x15, 0x0d, 0x8d, 0x3a, 0xbf, 0xfd, 0xd6, 0x86,
	0xaa, 0xbb, 0x1c, 0x14, 0x12, 0xe0, 0x63, 0x76, 0x4a, 0x5c, 0xd3, 0x49, 0x21, 0x68, 0x44, 0x58,
	0xad, 0x17, 0x00, 0x39, 0xf9, 0x9e, 0x5e, 0xb6, 0x57, 0xec, 0x65, 0x5a, 0xb1, 0x73, 0xfd, 0x54,
	0x02, 0xc5, 0xba, 0xa6, 0x7e, 0x8c, 0xcf, 0xa0, 0xec, 0xdc, 0xac, 0x28, 0x0f, 0x6b, 0x1c, 0x3c,
	0xcc, 0x93, 0xe1, 0x70, 0xf2, 0x65, 0x04, 0xc2, 0x69, 0xb7, 0xdb, 0xb5, 0xf4, 0x37, 0xed, 0xfa,
	0x03, 0x28, 0x33, 0x0d, 0x0b, 0x95, 0x6d, 0x91, 0x39, 0xa7, 0xb0, 0xcb, 0xeb, 0x07, 0xcb, 0xa5,
	0xeb, 0x7b, 0x69, 0x1b, 0x14, 0xa6, 0x39, 0x03, 0x2d, 0x4b, 0x02, 0x77, 0x40, 0x1b, 0xd8, 0x93,
	0xfe, 0xf8, 0xdc, 0x22, 0x5f, 0xe9, 0x0f, 0x50, 0x87, 0x7a, 0x7f, 0x3c, 0x3a, 0xb2, 0x5f, 0x4e,
	0xcf, 0x06, 0x3d, 0x87, 0x55, 0xf5, 0xff, 0xa0, 0x49, 0xac, 0x97, 0xf6, 0xc4, 0xb1, 0xc8, 0x37,
	0xc4, 0x9a, 0xb0, 0x0a, 0x2e, 0xb1, 0xa8, 0x13, 0xcb, 0x3a, 0xeb, 0x0d, 0xed, 0x73, 0x56, 0xdb,
	0x35, 0xa8, 0xf4, 0xc7, 0x23, 0x87, 0x8c, 0x87, 0x7a, 0x99, 0x35, 0x07, 0xeb, 0xb5, 0xd5, 0xd7,
	0x15, 0xf3, 0x0b, 0x68, 0x0e, 0x83, 0xe0, 0xcd, 0x7a, 0x95, 0x17, 0xda, 0x7d, 0xed, 0x4d, 0xba,
	0xbf, 0xbd, 0x99, 0xbf, 0x48, 0xa0, 0xdf, 0x8e, 0x8f, 0x56, 0xd8, 0xcb, 0x7f, 0x58, 0x49, 0xad,
	0xbf, 0x9f, 0x1f, 0xc2, 0x26, 0xf9, 0x3f, 0xf9, 0x65, 0xfd, 0x98, 0x65, 0x99, 0xde, 0xd6, 0xbf,
	0xd0, 0x4f, 0xee, 0xfd, 0x67, 0xc8, 0x5b, 0xfe, 0x19, 0xe6, 0x21, 0xec, 0x6e, 0x64, 0x10, 0xad,
	0xf0, 0x19, 0x54, 0xb3, 0x48, 0x69, 0xdb, 0x3b, 0x20, 0xa3, 0x1c, 0xfc, 0x5e, 0x82, 0xea, 0xe1,
	0xc5, 0xe2, 0x15, 0x83, 0xd1, 0x82, 0x7a, 0xf1, 0x2c, 0xf1, 0xe1, 0xb6, 0x33, 0x7e, 0xdb, 0x6a,
	0x6d, 0x3f, 0x7e, 0x3c, 0x86, 0x9d, 0x5b, 0x79, 0xe1, 0x1d, 0x72, 0x7e, 0x64, 0xad, 0x47, 0x5b,
	0xb1, 0x68, 0x85, 0x9f, 0x41, 0xbd, 0xd8, 0xb2, 0x8b, 0x09, 0x6d, 0xb4, 0xf2, 0x56, 0x23, 0x87,
	0xd8, 0xab, 0x09, 0x3f, 0x85, 0x6a, 0xfa, 0x90, 0xc1, 0xbb, 0x27, 0xd0, 0x32, 0xb6, 0xbd, 0x77,
	0xf0, 0x63, 0x80, 0xa9, 0x1f, 0xfe, 0x45, 0xe8, 0xe6, 0x4a, 0x5d, 0x50, 0x79, 0x5d, 0x45, 0xd8,
	0xdc, 0x28, 0xfa, 0xd6, 0xa6, 0xa3, 0x23, 0xed, 0x4b, 0x87, 0x8f, 0x60, 0xef, 0x32, 0x58, 0x76,
	0xaf, 0xe8, 0xd2, 0xbd, 0x59, 0xfb, 0x19, 0xe3, 0x58, 0xfa, 0x59, 0x92, 0x2e, 0x54, 0xfe, 0x0a,
	0xfc, 0xe4, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xd7, 0x7a, 0xe7, 0x17, 0x0a, 0x00, 0x00,
}
