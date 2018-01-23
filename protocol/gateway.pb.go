// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway.proto

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	gateway.proto
	login.proto
	service.proto

It has these top-level messages:
	Gateway
	G2CRegisterGateway
	C2GRegisterGatewayResult
	G2CRegisterGatewayUser
	C2GRegisterGatewayUserResult
	G2CRemoveGatewayUser
	C2LRegister
	L2CRegisterResult
	C2LLogin
	C2LLoginThirdSDK
	L2CLoginResult
	C2GTokenLogin
	G2CTokenLoginResult
	Service
	S2CServiceRegister
	C2SServiceRegisterResult
	C2GServiceRegister
	C2GServiceRemove
	X2XSendMessage2User
	RPCRequest
	RPCResponse
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type G2CRegisterGateway_MsgID int32

const (
	G2CRegisterGateway_Zero G2CRegisterGateway_MsgID = 0
	G2CRegisterGateway_ID   G2CRegisterGateway_MsgID = 1
)

var G2CRegisterGateway_MsgID_name = map[int32]string{
	0: "Zero",
	1: "ID",
}
var G2CRegisterGateway_MsgID_value = map[string]int32{
	"Zero": 0,
	"ID":   1,
}

func (x G2CRegisterGateway_MsgID) String() string {
	return proto.EnumName(G2CRegisterGateway_MsgID_name, int32(x))
}
func (G2CRegisterGateway_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type C2GRegisterGatewayResult_MsgID int32

const (
	C2GRegisterGatewayResult_Zero C2GRegisterGatewayResult_MsgID = 0
	C2GRegisterGatewayResult_ID   C2GRegisterGatewayResult_MsgID = 2
)

var C2GRegisterGatewayResult_MsgID_name = map[int32]string{
	0: "Zero",
	2: "ID",
}
var C2GRegisterGatewayResult_MsgID_value = map[string]int32{
	"Zero": 0,
	"ID":   2,
}

func (x C2GRegisterGatewayResult_MsgID) String() string {
	return proto.EnumName(C2GRegisterGatewayResult_MsgID_name, int32(x))
}
func (C2GRegisterGatewayResult_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 0}
}

type G2CRegisterGatewayUser_MsgID int32

const (
	G2CRegisterGatewayUser_Zero G2CRegisterGatewayUser_MsgID = 0
	G2CRegisterGatewayUser_ID   G2CRegisterGatewayUser_MsgID = 3
)

var G2CRegisterGatewayUser_MsgID_name = map[int32]string{
	0: "Zero",
	3: "ID",
}
var G2CRegisterGatewayUser_MsgID_value = map[string]int32{
	"Zero": 0,
	"ID":   3,
}

func (x G2CRegisterGatewayUser_MsgID) String() string {
	return proto.EnumName(G2CRegisterGatewayUser_MsgID_name, int32(x))
}
func (G2CRegisterGatewayUser_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

type C2GRegisterGatewayUserResult_MsgID int32

const (
	C2GRegisterGatewayUserResult_Zero C2GRegisterGatewayUserResult_MsgID = 0
	C2GRegisterGatewayUserResult_ID   C2GRegisterGatewayUserResult_MsgID = 4
)

var C2GRegisterGatewayUserResult_MsgID_name = map[int32]string{
	0: "Zero",
	4: "ID",
}
var C2GRegisterGatewayUserResult_MsgID_value = map[string]int32{
	"Zero": 0,
	"ID":   4,
}

func (x C2GRegisterGatewayUserResult_MsgID) String() string {
	return proto.EnumName(C2GRegisterGatewayUserResult_MsgID_name, int32(x))
}
func (C2GRegisterGatewayUserResult_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0}
}

type G2CRemoveGatewayUser_MsgID int32

const (
	G2CRemoveGatewayUser_Zero G2CRemoveGatewayUser_MsgID = 0
	G2CRemoveGatewayUser_ID   G2CRemoveGatewayUser_MsgID = 5
)

var G2CRemoveGatewayUser_MsgID_name = map[int32]string{
	0: "Zero",
	5: "ID",
}
var G2CRemoveGatewayUser_MsgID_value = map[string]int32{
	"Zero": 0,
	"ID":   5,
}

func (x G2CRemoveGatewayUser_MsgID) String() string {
	return proto.EnumName(G2CRemoveGatewayUser_MsgID_name, int32(x))
}
func (G2CRemoveGatewayUser_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 0}
}

type Gateway struct {
	Id      int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
}

func (m *Gateway) Reset()                    { *m = Gateway{} }
func (m *Gateway) String() string            { return proto.CompactTextString(m) }
func (*Gateway) ProtoMessage()               {}
func (*Gateway) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Gateway) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Gateway) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type G2CRegisterGateway struct {
	Gate *Gateway `protobuf:"bytes,1,opt,name=gate" json:"gate,omitempty"`
}

func (m *G2CRegisterGateway) Reset()                    { *m = G2CRegisterGateway{} }
func (m *G2CRegisterGateway) String() string            { return proto.CompactTextString(m) }
func (*G2CRegisterGateway) ProtoMessage()               {}
func (*G2CRegisterGateway) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *G2CRegisterGateway) GetGate() *Gateway {
	if m != nil {
		return m.Gate
	}
	return nil
}

type C2GRegisterGatewayResult struct {
	Result string `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *C2GRegisterGatewayResult) Reset()                    { *m = C2GRegisterGatewayResult{} }
func (m *C2GRegisterGatewayResult) String() string            { return proto.CompactTextString(m) }
func (*C2GRegisterGatewayResult) ProtoMessage()               {}
func (*C2GRegisterGatewayResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *C2GRegisterGatewayResult) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type G2CRegisterGatewayUser struct {
	Uid uint64 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *G2CRegisterGatewayUser) Reset()                    { *m = G2CRegisterGatewayUser{} }
func (m *G2CRegisterGatewayUser) String() string            { return proto.CompactTextString(m) }
func (*G2CRegisterGatewayUser) ProtoMessage()               {}
func (*G2CRegisterGatewayUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *G2CRegisterGatewayUser) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type C2GRegisterGatewayUserResult struct {
	Result string `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *C2GRegisterGatewayUserResult) Reset()                    { *m = C2GRegisterGatewayUserResult{} }
func (m *C2GRegisterGatewayUserResult) String() string            { return proto.CompactTextString(m) }
func (*C2GRegisterGatewayUserResult) ProtoMessage()               {}
func (*C2GRegisterGatewayUserResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *C2GRegisterGatewayUserResult) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type G2CRemoveGatewayUser struct {
	Uid uint64 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *G2CRemoveGatewayUser) Reset()                    { *m = G2CRemoveGatewayUser{} }
func (m *G2CRemoveGatewayUser) String() string            { return proto.CompactTextString(m) }
func (*G2CRemoveGatewayUser) ProtoMessage()               {}
func (*G2CRemoveGatewayUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *G2CRemoveGatewayUser) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func init() {
	proto.RegisterType((*Gateway)(nil), "protocol.Gateway")
	proto.RegisterType((*G2CRegisterGateway)(nil), "protocol.G2CRegisterGateway")
	proto.RegisterType((*C2GRegisterGatewayResult)(nil), "protocol.C2GRegisterGatewayResult")
	proto.RegisterType((*G2CRegisterGatewayUser)(nil), "protocol.G2CRegisterGatewayUser")
	proto.RegisterType((*C2GRegisterGatewayUserResult)(nil), "protocol.C2GRegisterGatewayUserResult")
	proto.RegisterType((*G2CRemoveGatewayUser)(nil), "protocol.G2CRemoveGatewayUser")
	proto.RegisterEnum("protocol.G2CRegisterGateway_MsgID", G2CRegisterGateway_MsgID_name, G2CRegisterGateway_MsgID_value)
	proto.RegisterEnum("protocol.C2GRegisterGatewayResult_MsgID", C2GRegisterGatewayResult_MsgID_name, C2GRegisterGatewayResult_MsgID_value)
	proto.RegisterEnum("protocol.G2CRegisterGatewayUser_MsgID", G2CRegisterGatewayUser_MsgID_name, G2CRegisterGatewayUser_MsgID_value)
	proto.RegisterEnum("protocol.C2GRegisterGatewayUserResult_MsgID", C2GRegisterGatewayUserResult_MsgID_name, C2GRegisterGatewayUserResult_MsgID_value)
	proto.RegisterEnum("protocol.G2CRemoveGatewayUser_MsgID", G2CRemoveGatewayUser_MsgID_name, G2CRemoveGatewayUser_MsgID_value)
}

func init() { proto.RegisterFile("gateway.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4f, 0x2c, 0x49,
	0x2d, 0x4f, 0xac, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39,
	0x4a, 0xc6, 0x5c, 0xec, 0xee, 0x10, 0x29, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xd6, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x09, 0x2e, 0xf6, 0xc4, 0x94, 0x94, 0xa2, 0xd4,
	0xe2, 0x62, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x29, 0x8c, 0x4b, 0xc8, 0xdd,
	0xc8, 0x39, 0x28, 0x35, 0x3d, 0xb3, 0xb8, 0x24, 0xb5, 0x08, 0xa6, 0x5f, 0x95, 0x8b, 0x05, 0x64,
	0x0b, 0xd8, 0x04, 0x6e, 0x23, 0x41, 0x3d, 0x98, 0x1d, 0x7a, 0x50, 0x05, 0x41, 0x60, 0x69, 0x25,
	0x49, 0x2e, 0x56, 0xdf, 0xe2, 0x74, 0x4f, 0x17, 0x21, 0x0e, 0x2e, 0x96, 0xa8, 0xd4, 0xa2, 0x7c,
	0x01, 0x06, 0x21, 0x36, 0x2e, 0x26, 0x4f, 0x17, 0x01, 0x46, 0x25, 0x5f, 0x2e, 0x09, 0x67, 0x23,
	0x77, 0x34, 0x73, 0x83, 0x52, 0x8b, 0x4b, 0x73, 0x4a, 0x84, 0xc4, 0xb8, 0xd8, 0x8a, 0xc0, 0x2c,
	0xb0, 0xf9, 0x9c, 0x41, 0x50, 0x1e, 0x6e, 0xe3, 0x98, 0x94, 0x5c, 0xb9, 0xc4, 0x30, 0x9d, 0x19,
	0x5a, 0x9c, 0x5a, 0x24, 0x24, 0xc0, 0xc5, 0x5c, 0x0a, 0xf5, 0x2b, 0x4b, 0x10, 0x88, 0x89, 0xdb,
	0x18, 0x66, 0xa5, 0x40, 0x2e, 0x19, 0x4c, 0x57, 0x81, 0x8c, 0x21, 0xd7, 0x65, 0x2c, 0x4a, 0xce,
	0x5c, 0x22, 0x60, 0x97, 0xe5, 0xe6, 0x97, 0xa5, 0x92, 0xe9, 0x2e, 0xd6, 0x24, 0x36, 0x70, 0x00,
	0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x55, 0x6f, 0x4f, 0x44, 0xdc, 0x01, 0x00, 0x00,
}
