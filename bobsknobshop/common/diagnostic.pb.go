// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bobsknobshop/common/diagnostic.proto

package bobsknobshop_common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	math "math"
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

type ServiceCallLog struct {
	// Name of the service
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// The request.
	Request *any.Any `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	// The response.
	Response             *any.Any `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceCallLog) Reset()         { *m = ServiceCallLog{} }
func (m *ServiceCallLog) String() string { return proto.CompactTextString(m) }
func (*ServiceCallLog) ProtoMessage()    {}
func (*ServiceCallLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_75d1c2931c22cc7b, []int{0}
}

func (m *ServiceCallLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceCallLog.Unmarshal(m, b)
}
func (m *ServiceCallLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceCallLog.Marshal(b, m, deterministic)
}
func (m *ServiceCallLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceCallLog.Merge(m, src)
}
func (m *ServiceCallLog) XXX_Size() int {
	return xxx_messageInfo_ServiceCallLog.Size(m)
}
func (m *ServiceCallLog) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceCallLog.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceCallLog proto.InternalMessageInfo

func (m *ServiceCallLog) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *ServiceCallLog) GetRequest() *any.Any {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *ServiceCallLog) GetResponse() *any.Any {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceCallLog)(nil), "bobsknobshop.common.ServiceCallLog")
}

func init() {
	proto.RegisterFile("bobsknobshop/common/diagnostic.proto", fileDescriptor_75d1c2931c22cc7b)
}

var fileDescriptor_75d1c2931c22cc7b = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0xca, 0x4f, 0x2a,
	0xce, 0xce, 0xcb, 0x4f, 0x2a, 0xce, 0xc8, 0x2f, 0xd0, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3,
	0x4f, 0xc9, 0x4c, 0x4c, 0xcf, 0xcb, 0x2f, 0x2e, 0xc9, 0x4c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0x46, 0x56, 0xa5, 0x07, 0x51, 0x25, 0x25, 0x99, 0x9e, 0x9f, 0x9f, 0x9e, 0x93, 0xaa,
	0x0f, 0x56, 0x92, 0x54, 0x9a, 0xa6, 0x9f, 0x98, 0x57, 0x09, 0x51, 0xaf, 0xd4, 0xc3, 0xc8, 0xc5,
	0x17, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0xea, 0x9c, 0x98, 0x93, 0xe3, 0x93, 0x9f, 0x2e, 0x24,
	0xc1, 0xc5, 0x5e, 0x0c, 0x11, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x85, 0xf4,
	0xb8, 0xd8, 0x8b, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x24, 0x98, 0x14, 0x18, 0x35, 0xb8, 0x8d,
	0x44, 0xf4, 0x20, 0x26, 0xeb, 0xc1, 0x4c, 0xd6, 0x73, 0xcc, 0xab, 0x0c, 0x82, 0x29, 0x12, 0x32,
	0xe0, 0xe2, 0x28, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x95, 0x60, 0xc6, 0xa3, 0x01, 0xae,
	0xca, 0x49, 0x8a, 0x4b, 0x2c, 0x3d, 0x47, 0x0f, 0x8b, 0x1f, 0x02, 0x18, 0x93, 0xd8, 0xc0, 0x7a,
	0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x96, 0x2a, 0xc9, 0xc0, 0x09, 0x01, 0x00, 0x00,
}