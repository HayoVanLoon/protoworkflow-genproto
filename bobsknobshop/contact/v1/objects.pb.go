// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bobsknobshop/contact/v1/objects.proto

package contact

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Details about a person that has sent a message.
type Sender struct {
	// The sender id, generated by the platform.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// E-mail address of the sender.
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// Name to address the sender with (optional).
	RealName string `protobuf:"bytes,3,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	// How the sender would like to be addressed (optional).
	Title                string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sender) Reset()         { *m = Sender{} }
func (m *Sender) String() string { return proto.CompactTextString(m) }
func (*Sender) ProtoMessage()    {}
func (*Sender) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8d31499ab9ffafd, []int{0}
}

func (m *Sender) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sender.Unmarshal(m, b)
}
func (m *Sender) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sender.Marshal(b, m, deterministic)
}
func (m *Sender) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sender.Merge(m, src)
}
func (m *Sender) XXX_Size() int {
	return xxx_messageInfo_Sender.Size(m)
}
func (m *Sender) XXX_DiscardUnknown() {
	xxx_messageInfo_Sender.DiscardUnknown(m)
}

var xxx_messageInfo_Sender proto.InternalMessageInfo

func (m *Sender) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Sender) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Sender) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *Sender) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func init() {
	proto.RegisterType((*Sender)(nil), "bobsknobshop.contact.v1.Sender")
}

func init() {
	proto.RegisterFile("bobsknobshop/contact/v1/objects.proto", fileDescriptor_d8d31499ab9ffafd)
}

var fileDescriptor_d8d31499ab9ffafd = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0xca, 0x4f, 0x2a,
	0xce, 0xce, 0xcb, 0x4f, 0x2a, 0xce, 0xc8, 0x2f, 0xd0, 0x4f, 0xce, 0xcf, 0x2b, 0x49, 0x4c, 0x2e,
	0xd1, 0x2f, 0x33, 0xd4, 0xcf, 0x4f, 0xca, 0x4a, 0x4d, 0x2e, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x47, 0x56, 0xa6, 0x07, 0x55, 0xa6, 0x57, 0x66, 0xa8, 0x94, 0xca, 0xc5, 0x16,
	0x9c, 0x9a, 0x97, 0x92, 0x5a, 0x24, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8,
	0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x0b, 0x89, 0x70, 0xb1, 0xa6, 0xe6, 0x26, 0x66, 0xe6, 0x48,
	0x30, 0x81, 0x05, 0x21, 0x1c, 0x21, 0x69, 0x2e, 0xce, 0xa2, 0xd4, 0xc4, 0x9c, 0x78, 0xb0, 0x72,
	0x66, 0xb0, 0x0c, 0x07, 0x48, 0xc0, 0x0f, 0xaa, 0xa5, 0x24, 0xb3, 0x24, 0x27, 0x55, 0x82, 0x05,
	0xa2, 0x05, 0xcc, 0x71, 0xca, 0xe5, 0x92, 0x4a, 0xcf, 0xd1, 0xc3, 0xe1, 0x88, 0x00, 0xc6, 0x28,
	0xdf, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x8f, 0xc4, 0xca, 0xfc,
	0xb0, 0xc4, 0x3c, 0x9f, 0xfc, 0xfc, 0x3c, 0x7d, 0xb0, 0xeb, 0xcb, 0xf3, 0x8b, 0xb2, 0xd3, 0x72,
	0xf2, 0xcb, 0x75, 0xd3, 0x53, 0xf3, 0xc0, 0x02, 0xfa, 0x38, 0x3c, 0x6d, 0x0d, 0x65, 0x26, 0xb1,
	0x81, 0x95, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xbb, 0x28, 0xd9, 0x1e, 0x01, 0x00,
	0x00,
}
