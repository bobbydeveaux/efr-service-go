// Code generated by protoc-gen-go.
// source: user.proto
// DO NOT EDIT!

/*
Package User is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	User
*/
package User

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

// The request message containing the user's name.
type User struct {
	SocialID  string `protobuf:"bytes,1,opt,name=SocialID,json=socialID" json:"SocialID,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=Email,json=email" json:"Email,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=FirstName,json=firstName" json:"FirstName,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=Name,json=name" json:"Name,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetSocialID() string {
	if m != nil {
		return m.SocialID
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "User.User")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x09, 0x2d, 0x4e, 0x2d, 0x52, 0xca, 0xe2, 0x02,
	0xd3, 0x42, 0x52, 0x5c, 0x1c, 0xc1, 0xf9, 0xc9, 0x99, 0x89, 0x39, 0x9e, 0x2e, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0x1c, 0xc5, 0x50, 0xbe, 0x90, 0x08, 0x17, 0xab, 0x6b, 0x6e, 0x62, 0x66,
	0x8e, 0x04, 0x13, 0x58, 0x82, 0x35, 0x15, 0xc4, 0x11, 0x92, 0xe1, 0xe2, 0x74, 0xcb, 0x2c, 0x2a,
	0x2e, 0xf1, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x06, 0xcb, 0x70, 0xa6, 0xc1, 0x04, 0x84, 0x84, 0xb8,
	0x58, 0xc0, 0x12, 0x2c, 0x60, 0x09, 0x96, 0xbc, 0xc4, 0xdc, 0xd4, 0x24, 0x36, 0xb0, 0xc5, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0xfb, 0x56, 0xd0, 0x86, 0x00, 0x00, 0x00,
}
