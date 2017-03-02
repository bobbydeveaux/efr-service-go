// Code generated by protoc-gen-go.
// source: users.proto
// DO NOT EDIT!

/*
Package Users is a generated protocol buffer package.

It is generated from these files:
	users.proto

It has these top-level messages:
	User
	UserRequest
	UserReply
*/
package Users

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

// The request message containing the user's name.
type User struct {
	SocialID  string `protobuf:"bytes,1,opt,name=SocialID,json=socialID" json:"SocialID,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=Email,json=email" json:"Email,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=FirstName,json=firstName" json:"FirstName,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=Name,json=name" json:"Name,omitempty"`
	LastLogin int64  `protobuf:"varint,5,opt,name=LastLogin,json=lastLogin" json:"LastLogin,omitempty"`
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

func (m *User) GetLastLogin() int64 {
	if m != nil {
		return m.LastLogin
	}
	return 0
}

// The request message containing the user's name.
type UserRequest struct {
	User *User `protobuf:"bytes,1,opt,name=User,json=user" json:"User,omitempty"`
}

func (m *UserRequest) Reset()                    { *m = UserRequest{} }
func (m *UserRequest) String() string            { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()               {}
func (*UserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// The response message containing the greetings
type UserReply struct {
	User *User `protobuf:"bytes,1,opt,name=User,json=user" json:"User,omitempty"`
}

func (m *UserReply) Reset()                    { *m = UserReply{} }
func (m *UserReply) String() string            { return proto.CompactTextString(m) }
func (*UserReply) ProtoMessage()               {}
func (*UserReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UserReply) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "Users.User")
	proto.RegisterType((*UserRequest)(nil), "Users.UserRequest")
	proto.RegisterType((*UserReply)(nil), "Users.UserReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Users service

type UsersClient interface {
	// Sends a greeting
	UpdateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserReply, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) UpdateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := grpc.Invoke(ctx, "/Users.Users/UpdateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Users service

type UsersServer interface {
	// Sends a greeting
	UpdateUser(context.Context, *UserRequest) (*UserReply, error)
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Users.Users/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).UpdateUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Users.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateUser",
			Handler:    _Users_UpdateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}

func init() { proto.RegisterFile("users.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x90, 0xbf, 0x4a, 0x43, 0x31,
	0x14, 0xc6, 0x8d, 0x4d, 0xa4, 0x39, 0x59, 0xe4, 0xe0, 0x10, 0x4a, 0xc1, 0x92, 0xa9, 0x83, 0xdc,
	0xa1, 0xba, 0xba, 0xa9, 0x20, 0x14, 0x87, 0x48, 0x1f, 0x20, 0x6a, 0x94, 0x40, 0xda, 0xc4, 0x9b,
	0x74, 0xe8, 0x1b, 0xf8, 0xd8, 0x92, 0x13, 0x0b, 0x3a, 0xb9, 0x84, 0x7c, 0x7f, 0x20, 0xbf, 0x2f,
	0xa0, 0xf6, 0xc5, 0x8f, 0x65, 0xc8, 0x63, 0xaa, 0x09, 0xc5, 0xa6, 0x09, 0xf3, 0xc5, 0x80, 0xb7,
	0x1b, 0xce, 0x60, 0xfa, 0x9c, 0x5e, 0x83, 0x8b, 0x8f, 0x77, 0x9a, 0x2d, 0xd8, 0x52, 0xda, 0x69,
	0xf9, 0xd1, 0x78, 0x01, 0xe2, 0x7e, 0xeb, 0x42, 0xd4, 0xa7, 0x14, 0x08, 0xdf, 0x04, 0xce, 0x41,
	0x3e, 0x84, 0xb1, 0xd4, 0x27, 0xb7, 0xf5, 0x7a, 0x42, 0x89, 0x7c, 0x3f, 0x1a, 0x88, 0xc0, 0x29,
	0xe0, 0x14, 0xf0, 0x5d, 0xf3, 0xe6, 0x20, 0xd7, 0xae, 0xd4, 0x75, 0xfa, 0x08, 0x3b, 0x2d, 0x16,
	0x6c, 0x39, 0xb1, 0x32, 0x1e, 0x0d, 0x33, 0x80, 0x6a, 0x24, 0xd6, 0x7f, 0xee, 0x7d, 0xa9, 0x78,
	0xd9, 0xc1, 0x08, 0x46, 0xad, 0xd4, 0x40, 0xbc, 0x74, 0x5a, 0xde, 0x86, 0x98, 0x2b, 0x90, 0xbd,
	0x9f, 0xe3, 0xe1, 0xdf, 0xf6, 0xea, 0x16, 0xfa, 0x62, 0xbc, 0x01, 0xd8, 0xe4, 0x37, 0x57, 0x3d,
	0xcd, 0xc6, 0xdf, 0xcd, 0xfe, 0xf2, 0xec, 0xfc, 0x8f, 0x97, 0xe3, 0xc1, 0x9c, 0xbc, 0x9c, 0xd1,
	0xaf, 0x5d, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x54, 0xf4, 0xa2, 0xbb, 0x44, 0x01, 0x00, 0x00,
}