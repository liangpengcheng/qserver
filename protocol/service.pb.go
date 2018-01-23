// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package protocol

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

type S2CServiceRegister_MsgID int32

const (
	S2CServiceRegister_Zero S2CServiceRegister_MsgID = 0
	S2CServiceRegister_ID   S2CServiceRegister_MsgID = 201
)

var S2CServiceRegister_MsgID_name = map[int32]string{
	0:   "Zero",
	201: "ID",
}
var S2CServiceRegister_MsgID_value = map[string]int32{
	"Zero": 0,
	"ID":   201,
}

func (x S2CServiceRegister_MsgID) String() string {
	return proto.EnumName(S2CServiceRegister_MsgID_name, int32(x))
}
func (S2CServiceRegister_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 0} }

type C2SServiceRegisterResult_MsgID int32

const (
	C2SServiceRegisterResult_Zero C2SServiceRegisterResult_MsgID = 0
	C2SServiceRegisterResult_ID   C2SServiceRegisterResult_MsgID = 202
)

var C2SServiceRegisterResult_MsgID_name = map[int32]string{
	0:   "Zero",
	202: "ID",
}
var C2SServiceRegisterResult_MsgID_value = map[string]int32{
	"Zero": 0,
	"ID":   202,
}

func (x C2SServiceRegisterResult_MsgID) String() string {
	return proto.EnumName(C2SServiceRegisterResult_MsgID_name, int32(x))
}
func (C2SServiceRegisterResult_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor2, []int{2, 0}
}

type C2GServiceRegister_MsgID int32

const (
	C2GServiceRegister_Zero C2GServiceRegister_MsgID = 0
	C2GServiceRegister_ID   C2GServiceRegister_MsgID = 203
)

var C2GServiceRegister_MsgID_name = map[int32]string{
	0:   "Zero",
	203: "ID",
}
var C2GServiceRegister_MsgID_value = map[string]int32{
	"Zero": 0,
	"ID":   203,
}

func (x C2GServiceRegister_MsgID) String() string {
	return proto.EnumName(C2GServiceRegister_MsgID_name, int32(x))
}
func (C2GServiceRegister_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{3, 0} }

type C2GServiceRemove_MsgID int32

const (
	C2GServiceRemove_Zero C2GServiceRemove_MsgID = 0
	C2GServiceRemove_ID   C2GServiceRemove_MsgID = 204
)

var C2GServiceRemove_MsgID_name = map[int32]string{
	0:   "Zero",
	204: "ID",
}
var C2GServiceRemove_MsgID_value = map[string]int32{
	"Zero": 0,
	"ID":   204,
}

func (x C2GServiceRemove_MsgID) String() string {
	return proto.EnumName(C2GServiceRemove_MsgID_name, int32(x))
}
func (C2GServiceRemove_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{4, 0} }

type X2XSendMessage2User_MsgID int32

const (
	X2XSendMessage2User_Zero X2XSendMessage2User_MsgID = 0
	X2XSendMessage2User_ID   X2XSendMessage2User_MsgID = 205
)

var X2XSendMessage2User_MsgID_name = map[int32]string{
	0:   "Zero",
	205: "ID",
}
var X2XSendMessage2User_MsgID_value = map[string]int32{
	"Zero": 0,
	"ID":   205,
}

func (x X2XSendMessage2User_MsgID) String() string {
	return proto.EnumName(X2XSendMessage2User_MsgID_name, int32(x))
}
func (X2XSendMessage2User_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{5, 0} }

type Service struct {
	Handler []int32 `protobuf:"varint,1,rep,packed,name=handler" json:"handler,omitempty"`
	Version int32   `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	Name    string  `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Address string  `protobuf:"bytes,4,opt,name=address" json:"address,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Service) GetHandler() []int32 {
	if m != nil {
		return m.Handler
	}
	return nil
}

func (m *Service) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// service 2 center server 注册
type S2CServiceRegister struct {
	Serv *Service `protobuf:"bytes,1,opt,name=serv" json:"serv,omitempty"`
}

func (m *S2CServiceRegister) Reset()                    { *m = S2CServiceRegister{} }
func (m *S2CServiceRegister) String() string            { return proto.CompactTextString(m) }
func (*S2CServiceRegister) ProtoMessage()               {}
func (*S2CServiceRegister) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *S2CServiceRegister) GetServ() *Service {
	if m != nil {
		return m.Serv
	}
	return nil
}

// center 2 service 注册结果
type C2SServiceRegisterResult struct {
	Result string `protobuf:"bytes,1,opt,name=Result" json:"Result,omitempty"`
}

func (m *C2SServiceRegisterResult) Reset()                    { *m = C2SServiceRegisterResult{} }
func (m *C2SServiceRegisterResult) String() string            { return proto.CompactTextString(m) }
func (*C2SServiceRegisterResult) ProtoMessage()               {}
func (*C2SServiceRegisterResult) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *C2SServiceRegisterResult) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

// center 2 gateway 注册 service,广播
type C2GServiceRegister struct {
	Serv *Service `protobuf:"bytes,1,opt,name=serv" json:"serv,omitempty"`
}

func (m *C2GServiceRegister) Reset()                    { *m = C2GServiceRegister{} }
func (m *C2GServiceRegister) String() string            { return proto.CompactTextString(m) }
func (*C2GServiceRegister) ProtoMessage()               {}
func (*C2GServiceRegister) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *C2GServiceRegister) GetServ() *Service {
	if m != nil {
		return m.Serv
	}
	return nil
}

// center2 gateway 删除service ,广播
type C2GServiceRemove struct {
	Serv *Service `protobuf:"bytes,1,opt,name=serv" json:"serv,omitempty"`
}

func (m *C2GServiceRemove) Reset()                    { *m = C2GServiceRemove{} }
func (m *C2GServiceRemove) String() string            { return proto.CompactTextString(m) }
func (*C2GServiceRemove) ProtoMessage()               {}
func (*C2GServiceRemove) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *C2GServiceRemove) GetServ() *Service {
	if m != nil {
		return m.Serv
	}
	return nil
}

type X2XSendMessage2User struct {
	Sendto  uint64 `protobuf:"varint,1,opt,name=sendto" json:"sendto,omitempty"`
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *X2XSendMessage2User) Reset()                    { *m = X2XSendMessage2User{} }
func (m *X2XSendMessage2User) String() string            { return proto.CompactTextString(m) }
func (*X2XSendMessage2User) ProtoMessage()               {}
func (*X2XSendMessage2User) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *X2XSendMessage2User) GetSendto() uint64 {
	if m != nil {
		return m.Sendto
	}
	return 0
}

func (m *X2XSendMessage2User) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type RPCRequest struct {
	Userid    uint64 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	Messageid int32  `protobuf:"varint,2,opt,name=messageid" json:"messageid,omitempty"`
	Body      []byte `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *RPCRequest) Reset()                    { *m = RPCRequest{} }
func (m *RPCRequest) String() string            { return proto.CompactTextString(m) }
func (*RPCRequest) ProtoMessage()               {}
func (*RPCRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *RPCRequest) GetUserid() uint64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *RPCRequest) GetMessageid() int32 {
	if m != nil {
		return m.Messageid
	}
	return 0
}

func (m *RPCRequest) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type RPCResponse struct {
	Response []byte `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (m *RPCResponse) Reset()                    { *m = RPCResponse{} }
func (m *RPCResponse) String() string            { return proto.CompactTextString(m) }
func (*RPCResponse) ProtoMessage()               {}
func (*RPCResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *RPCResponse) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*Service)(nil), "protocol.Service")
	proto.RegisterType((*S2CServiceRegister)(nil), "protocol.S2CServiceRegister")
	proto.RegisterType((*C2SServiceRegisterResult)(nil), "protocol.C2SServiceRegisterResult")
	proto.RegisterType((*C2GServiceRegister)(nil), "protocol.C2GServiceRegister")
	proto.RegisterType((*C2GServiceRemove)(nil), "protocol.C2GServiceRemove")
	proto.RegisterType((*X2XSendMessage2User)(nil), "protocol.X2XSendMessage2User")
	proto.RegisterType((*RPCRequest)(nil), "protocol.RPCRequest")
	proto.RegisterType((*RPCResponse)(nil), "protocol.RPCResponse")
	proto.RegisterEnum("protocol.S2CServiceRegister_MsgID", S2CServiceRegister_MsgID_name, S2CServiceRegister_MsgID_value)
	proto.RegisterEnum("protocol.C2SServiceRegisterResult_MsgID", C2SServiceRegisterResult_MsgID_name, C2SServiceRegisterResult_MsgID_value)
	proto.RegisterEnum("protocol.C2GServiceRegister_MsgID", C2GServiceRegister_MsgID_name, C2GServiceRegister_MsgID_value)
	proto.RegisterEnum("protocol.C2GServiceRemove_MsgID", C2GServiceRemove_MsgID_name, C2GServiceRemove_MsgID_value)
	proto.RegisterEnum("protocol.X2XSendMessage2User_MsgID", X2XSendMessage2User_MsgID_name, X2XSendMessage2User_MsgID_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GatewayCallService service

type GatewayCallServiceClient interface {
	Request(ctx context.Context, in *RPCRequest, opts ...grpc.CallOption) (*RPCResponse, error)
}

type gatewayCallServiceClient struct {
	cc *grpc.ClientConn
}

func NewGatewayCallServiceClient(cc *grpc.ClientConn) GatewayCallServiceClient {
	return &gatewayCallServiceClient{cc}
}

func (c *gatewayCallServiceClient) Request(ctx context.Context, in *RPCRequest, opts ...grpc.CallOption) (*RPCResponse, error) {
	out := new(RPCResponse)
	err := grpc.Invoke(ctx, "/protocol.GatewayCallService/Request", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GatewayCallService service

type GatewayCallServiceServer interface {
	Request(context.Context, *RPCRequest) (*RPCResponse, error)
}

func RegisterGatewayCallServiceServer(s *grpc.Server, srv GatewayCallServiceServer) {
	s.RegisterService(&_GatewayCallService_serviceDesc, srv)
}

func _GatewayCallService_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayCallServiceServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.GatewayCallService/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayCallServiceServer).Request(ctx, req.(*RPCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GatewayCallService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.GatewayCallService",
	HandlerType: (*GatewayCallServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _GatewayCallService_Request_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcd, 0xae, 0xda, 0x30,
	0x10, 0x85, 0xaf, 0xef, 0x0d, 0x04, 0x06, 0x2a, 0x51, 0xf7, 0x47, 0x11, 0xea, 0x22, 0xb2, 0x54,
	0x89, 0x6e, 0x58, 0xa4, 0x9b, 0xee, 0x83, 0x84, 0x58, 0x80, 0x2a, 0x47, 0xb4, 0xa8, 0xbb, 0x10,
	0x8f, 0x68, 0xd4, 0x60, 0x53, 0x3b, 0x50, 0xf1, 0x88, 0xfd, 0x7b, 0xa7, 0xca, 0x4e, 0x5c, 0x2a,
	0x24, 0x56, 0xbd, 0xab, 0xcc, 0x99, 0xb1, 0xbe, 0x39, 0xb1, 0x0f, 0x3c, 0x31, 0xa8, 0x4f, 0x65,
	0x81, 0xd3, 0x83, 0x56, 0xb5, 0xa2, 0x3d, 0xf7, 0x29, 0x54, 0xc5, 0xbe, 0x40, 0x98, 0x35, 0x23,
	0x1a, 0x41, 0xf8, 0x39, 0x97, 0xa2, 0x42, 0x1d, 0x91, 0xf8, 0x61, 0xd2, 0xe1, 0x5e, 0xda, 0xc9,
	0x09, 0xb5, 0x29, 0x95, 0x8c, 0xee, 0x63, 0x62, 0x27, 0xad, 0xa4, 0x14, 0x02, 0x99, 0xef, 0x31,
	0x7a, 0x88, 0xc9, 0xa4, 0xcf, 0x5d, 0x6d, 0x4f, 0xe7, 0x42, 0x68, 0x34, 0x26, 0x0a, 0x5c, 0xdb,
	0x4b, 0xf6, 0x11, 0x68, 0x96, 0xa4, 0xed, 0x3e, 0x8e, 0xbb, 0xd2, 0xd4, 0xa8, 0xe9, 0x6b, 0x08,
	0xac, 0xbb, 0x88, 0xc4, 0x64, 0x32, 0x48, 0x9e, 0x4e, 0xbd, 0xb7, 0xa9, 0x3f, 0xe8, 0xc6, 0x6c,
	0x0c, 0x9d, 0xa5, 0xd9, 0x2d, 0x66, 0xb4, 0x07, 0xc1, 0x27, 0xd4, 0x6a, 0x74, 0x47, 0x43, 0xb8,
	0x5f, 0xcc, 0x46, 0xdf, 0x09, 0x5b, 0x41, 0x94, 0x26, 0xd9, 0x15, 0x98, 0xa3, 0x39, 0x56, 0x35,
	0x7d, 0x09, 0xdd, 0xa6, 0x72, 0x0b, 0xfa, 0xbc, 0x55, 0xb7, 0x79, 0x3f, 0x88, 0x35, 0x9a, 0x26,
	0xf3, 0xc7, 0x36, 0xfa, 0x93, 0xb0, 0x35, 0x8c, 0xfe, 0x05, 0xef, 0xd5, 0x09, 0xff, 0x1b, 0xfb,
	0x8b, 0xb0, 0x02, 0x9e, 0x6d, 0x92, 0x4d, 0x86, 0x52, 0x2c, 0xd1, 0x98, 0x7c, 0x87, 0xc9, 0xda,
	0xa0, 0xb6, 0xbf, 0x6e, 0x50, 0x8a, 0x5a, 0x39, 0x76, 0xc0, 0x5b, 0x65, 0x5f, 0xa8, 0x50, 0xb2,
	0x46, 0x59, 0xbb, 0xf7, 0x1c, 0x72, 0x2f, 0x6f, 0x2f, 0xf9, 0x4d, 0xd8, 0x07, 0x00, 0xfe, 0x3e,
	0xe5, 0xf8, 0xf5, 0x88, 0xc6, 0x5d, 0xeb, 0xd1, 0xa0, 0x2e, 0x85, 0x67, 0x37, 0x8a, 0xbe, 0x82,
	0xfe, 0xbe, 0xf1, 0x50, 0x8a, 0x36, 0x2d, 0x97, 0x86, 0xcd, 0xcb, 0x56, 0x89, 0xb3, 0xcb, 0xcb,
	0x90, 0xbb, 0x9a, 0xbd, 0x81, 0x81, 0xe3, 0x9a, 0x83, 0x92, 0x06, 0xe9, 0x18, 0x7a, 0xba, 0xad,
	0x1d, 0x7a, 0xc8, 0xff, 0xea, 0x64, 0x05, 0x74, 0x9e, 0xd7, 0xf8, 0x2d, 0x3f, 0xa7, 0x79, 0x55,
	0xf9, 0xe0, 0xbe, 0x83, 0xd0, 0xbb, 0x7a, 0x7e, 0xb9, 0xbd, 0x8b, 0xd7, 0xf1, 0x8b, 0xab, 0x6e,
	0x43, 0x63, 0x77, 0xdb, 0xae, 0xeb, 0xbf, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0x04, 0x0c, 0xcb,
	0x6e, 0x1f, 0x03, 0x00, 0x00,
}
