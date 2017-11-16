// Code generated by protoc-gen-go.
// source: msg.proto
// DO NOT EDIT!

/*
Package msg is a generated protocol buffer package.

It is generated from these files:
	msg.proto

It has these top-level messages:
	SCSysBusy
	CSLogin
	SCLogin
	CSEnterWorld
	SCEnterWorld
*/
package msg

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

// msgid
type EMsg int32

const (
	EMsg_ID_SCSysBusy    EMsg = 2
	EMsg_ID_CSLogin      EMsg = 1001
	EMsg_ID_SCLogin      EMsg = 1002
	EMsg_ID_CSEnterWorld EMsg = 2001
	EMsg_ID_SCEnterWorld EMsg = 2002
)

var EMsg_name = map[int32]string{
	2:    "ID_SCSysBusy",
	1001: "ID_CSLogin",
	1002: "ID_SCLogin",
	2001: "ID_CSEnterWorld",
	2002: "ID_SCEnterWorld",
}
var EMsg_value = map[string]int32{
	"ID_SCSysBusy":    2,
	"ID_CSLogin":      1001,
	"ID_SCLogin":      1002,
	"ID_CSEnterWorld": 2001,
	"ID_SCEnterWorld": 2002,
}

func (x EMsg) Enum() *EMsg {
	p := new(EMsg)
	*p = x
	return p
}
func (x EMsg) String() string {
	return proto.EnumName(EMsg_name, int32(x))
}
func (x *EMsg) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EMsg_value, data, "EMsg")
	if err != nil {
		return err
	}
	*x = EMsg(value)
	return nil
}
func (EMsg) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

//
type SCSysBusy struct {
	Gateway          *string `protobuf:"bytes,99,opt,name=gateway,def=to=client" json:"gateway,omitempty"`
	SrvType          *int32  `protobuf:"varint,1,req,name=SrvType" json:"SrvType,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SCSysBusy) Reset()                    { *m = SCSysBusy{} }
func (m *SCSysBusy) String() string            { return proto.CompactTextString(m) }
func (*SCSysBusy) ProtoMessage()               {}
func (*SCSysBusy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_SCSysBusy_Gateway string = "to=client"

func (m *SCSysBusy) GetGateway() string {
	if m != nil && m.Gateway != nil {
		return *m.Gateway
	}
	return Default_SCSysBusy_Gateway
}

func (m *SCSysBusy) GetSrvType() int32 {
	if m != nil && m.SrvType != nil {
		return *m.SrvType
	}
	return 0
}

//
type CSLogin struct {
	Gateway          *string `protobuf:"bytes,99,opt,name=gateway,def=to=logon" json:"gateway,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CSLogin) Reset()                    { *m = CSLogin{} }
func (m *CSLogin) String() string            { return proto.CompactTextString(m) }
func (*CSLogin) ProtoMessage()               {}
func (*CSLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

const Default_CSLogin_Gateway string = "to=logon"

func (m *CSLogin) GetGateway() string {
	if m != nil && m.Gateway != nil {
		return *m.Gateway
	}
	return Default_CSLogin_Gateway
}

type SCLogin struct {
	Gateway          *string `protobuf:"bytes,99,opt,name=gateway,def=to=client|accId=" json:"gateway,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SCLogin) Reset()                    { *m = SCLogin{} }
func (m *SCLogin) String() string            { return proto.CompactTextString(m) }
func (*SCLogin) ProtoMessage()               {}
func (*SCLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

const Default_SCLogin_Gateway string = "to=client|accId="

func (m *SCLogin) GetGateway() string {
	if m != nil && m.Gateway != nil {
		return *m.Gateway
	}
	return Default_SCLogin_Gateway
}

//
type CSEnterWorld struct {
	Gateway          *string `protobuf:"bytes,99,opt,name=gateway,def=to=user|url=auto" json:"gateway,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CSEnterWorld) Reset()                    { *m = CSEnterWorld{} }
func (m *CSEnterWorld) String() string            { return proto.CompactTextString(m) }
func (*CSEnterWorld) ProtoMessage()               {}
func (*CSEnterWorld) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

const Default_CSEnterWorld_Gateway string = "to=user|url=auto"

func (m *CSEnterWorld) GetGateway() string {
	if m != nil && m.Gateway != nil {
		return *m.Gateway
	}
	return Default_CSEnterWorld_Gateway
}

type SCEnterWorld struct {
	Gateway          *string `protobuf:"bytes,99,opt,name=gateway,def=to=client|url=set" json:"gateway,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SCEnterWorld) Reset()                    { *m = SCEnterWorld{} }
func (m *SCEnterWorld) String() string            { return proto.CompactTextString(m) }
func (*SCEnterWorld) ProtoMessage()               {}
func (*SCEnterWorld) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

const Default_SCEnterWorld_Gateway string = "to=client|url=set"

func (m *SCEnterWorld) GetGateway() string {
	if m != nil && m.Gateway != nil {
		return *m.Gateway
	}
	return Default_SCEnterWorld_Gateway
}

func init() {
	proto.RegisterType((*SCSysBusy)(nil), "msg.SCSysBusy")
	proto.RegisterType((*CSLogin)(nil), "msg.CSLogin")
	proto.RegisterType((*SCLogin)(nil), "msg.SCLogin")
	proto.RegisterType((*CSEnterWorld)(nil), "msg.CSEnterWorld")
	proto.RegisterType((*SCEnterWorld)(nil), "msg.SCEnterWorld")
	proto.RegisterEnum("msg.EMsg", EMsg_name, EMsg_value)
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0xd9, 0x54, 0x62, 0xff, 0x14, 0x1a, 0x83, 0x87, 0x1e, 0x47, 0xbd, 0x8c, 0x89, 0xde,
	0xbc, 0x4c, 0x7a, 0x31, 0xdb, 0xa1, 0xa2, 0x97, 0x45, 0xf0, 0x38, 0x4a, 0x17, 0xc3, 0xa0, 0xeb,
	0x7f, 0x24, 0xa9, 0x12, 0xd8, 0x97, 0xd4, 0x6f, 0xa1, 0x9f, 0x42, 0xd4, 0xac, 0x19, 0xc5, 0x63,
	0x5e, 0xde, 0xef, 0xbd, 0xbc, 0x40, 0xb4, 0x31, 0xea, 0x7a, 0xab, 0xd1, 0x22, 0x3b, 0xda, 0x18,
	0x95, 0xdd, 0x43, 0x24, 0xb8, 0x70, 0xe6, 0xae, 0x35, 0x8e, 0x5d, 0x00, 0x51, 0xa5, 0x95, 0x6f,
	0xa5, 0x4b, 0xab, 0xd1, 0x60, 0x1c, 0x4d, 0x23, 0x8b, 0x79, 0x55, 0xaf, 0x65, 0x63, 0x17, 0xfb,
	0x1b, 0x96, 0x02, 0x11, 0xfa, 0xf5, 0xc9, 0x6d, 0x65, 0x3a, 0x18, 0x0d, 0xc7, 0x27, 0x8b, 0xfd,
	0x31, 0xbb, 0x02, 0xc2, 0xc5, 0x03, 0xaa, 0x75, 0xc3, 0xb2, 0x7e, 0xd2, 0xa9, 0xc5, 0xbc, 0x46,
	0x85, 0x4d, 0x17, 0x94, 0xdd, 0x00, 0x11, 0xfc, 0xcf, 0x3e, 0xe9, 0xdb, 0x69, 0x57, 0xbc, 0x2b,
	0xab, 0xaa, 0x58, 0xe5, 0x01, 0x9b, 0x42, 0xcc, 0xc5, 0xbc, 0xb1, 0x52, 0x3f, 0xa3, 0xae, 0x57,
	0xff, 0xb3, 0xad, 0x91, 0x7a, 0xd7, 0xea, 0x3a, 0x2f, 0x5b, 0x8b, 0x81, 0xbd, 0x85, 0x58, 0xf0,
	0x03, 0xf6, 0xb2, 0xcf, 0x9e, 0x85, 0xde, 0x1f, 0xda, 0xc8, 0x30, 0x7c, 0xf2, 0x02, 0xc7, 0xf3,
	0x47, 0xa3, 0x18, 0x85, 0xb8, 0x98, 0x2d, 0xbb, 0x5f, 0xa3, 0x43, 0x96, 0x00, 0x14, 0xb3, 0xa5,
	0xdf, 0x4e, 0x3f, 0x89, 0x17, 0xfc, 0x3a, 0xfa, 0x45, 0xd8, 0x39, 0x24, 0xbf, 0x8e, 0xd0, 0x4d,
	0xdf, 0x13, 0xaf, 0x1e, 0xbe, 0x88, 0x7e, 0x24, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xab, 0x9a,
	0x60, 0x56, 0xa3, 0x01, 0x00, 0x00,
}
