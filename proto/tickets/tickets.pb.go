// Code generated by protoc-gen-go.
// source: tickets.proto
// DO NOT EDIT!

/*
Package Tickets is a generated protocol buffer package.

It is generated from these files:
	tickets.proto

It has these top-level messages:
	TicketRequest
	TicketReply
*/
package Tickets

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
type TicketRequest struct {
	Email    string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Socialid int64  `protobuf:"varint,2,opt,name=socialid" json:"socialid,omitempty"`
	Referrer int64  `protobuf:"varint,3,opt,name=referrer" json:"referrer,omitempty"`
}

func (m *TicketRequest) Reset()                    { *m = TicketRequest{} }
func (m *TicketRequest) String() string            { return proto.CompactTextString(m) }
func (*TicketRequest) ProtoMessage()               {}
func (*TicketRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TicketRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *TicketRequest) GetSocialid() int64 {
	if m != nil {
		return m.Socialid
	}
	return 0
}

func (m *TicketRequest) GetReferrer() int64 {
	if m != nil {
		return m.Referrer
	}
	return 0
}

// The response message containing the greetings
type TicketReply struct {
	Ticketid string                `protobuf:"bytes,1,opt,name=ticketid" json:"ticketid,omitempty"`
	Tickets  []*TicketReply_Ticket `protobuf:"bytes,2,rep,name=Tickets,json=tickets" json:"Tickets,omitempty"`
}

func (m *TicketReply) Reset()                    { *m = TicketReply{} }
func (m *TicketReply) String() string            { return proto.CompactTextString(m) }
func (*TicketReply) ProtoMessage()               {}
func (*TicketReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TicketReply) GetTicketid() string {
	if m != nil {
		return m.Ticketid
	}
	return ""
}

func (m *TicketReply) GetTickets() []*TicketReply_Ticket {
	if m != nil {
		return m.Tickets
	}
	return nil
}

type TicketReply_Ticket struct {
	TicketID string `protobuf:"bytes,1,opt,name=TicketID,json=ticketID" json:"TicketID,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=Email,json=email" json:"Email,omitempty"`
	SocialID int64  `protobuf:"varint,3,opt,name=SocialID,json=socialID" json:"SocialID,omitempty"`
	Referrer int64  `protobuf:"varint,4,opt,name=Referrer,json=referrer" json:"Referrer,omitempty"`
	Bonus    bool   `protobuf:"varint,5,opt,name=Bonus,json=bonus" json:"Bonus,omitempty"`
}

func (m *TicketReply_Ticket) Reset()                    { *m = TicketReply_Ticket{} }
func (m *TicketReply_Ticket) String() string            { return proto.CompactTextString(m) }
func (*TicketReply_Ticket) ProtoMessage()               {}
func (*TicketReply_Ticket) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *TicketReply_Ticket) GetTicketID() string {
	if m != nil {
		return m.TicketID
	}
	return ""
}

func (m *TicketReply_Ticket) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *TicketReply_Ticket) GetSocialID() int64 {
	if m != nil {
		return m.SocialID
	}
	return 0
}

func (m *TicketReply_Ticket) GetReferrer() int64 {
	if m != nil {
		return m.Referrer
	}
	return 0
}

func (m *TicketReply_Ticket) GetBonus() bool {
	if m != nil {
		return m.Bonus
	}
	return false
}

func init() {
	proto.RegisterType((*TicketRequest)(nil), "Tickets.TicketRequest")
	proto.RegisterType((*TicketReply)(nil), "Tickets.TicketReply")
	proto.RegisterType((*TicketReply_Ticket)(nil), "Tickets.TicketReply.Ticket")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Tickets service

type TicketsClient interface {
	// Sends a greeting
	NewTicket(ctx context.Context, in *TicketRequest, opts ...grpc.CallOption) (*TicketReply, error)
	GetTickets(ctx context.Context, in *TicketRequest, opts ...grpc.CallOption) (*TicketReply, error)
}

type ticketsClient struct {
	cc *grpc.ClientConn
}

func NewTicketsClient(cc *grpc.ClientConn) TicketsClient {
	return &ticketsClient{cc}
}

func (c *ticketsClient) NewTicket(ctx context.Context, in *TicketRequest, opts ...grpc.CallOption) (*TicketReply, error) {
	out := new(TicketReply)
	err := grpc.Invoke(ctx, "/Tickets.Tickets/NewTicket", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketsClient) GetTickets(ctx context.Context, in *TicketRequest, opts ...grpc.CallOption) (*TicketReply, error) {
	out := new(TicketReply)
	err := grpc.Invoke(ctx, "/Tickets.Tickets/GetTickets", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Tickets service

type TicketsServer interface {
	// Sends a greeting
	NewTicket(context.Context, *TicketRequest) (*TicketReply, error)
	GetTickets(context.Context, *TicketRequest) (*TicketReply, error)
}

func RegisterTicketsServer(s *grpc.Server, srv TicketsServer) {
	s.RegisterService(&_Tickets_serviceDesc, srv)
}

func _Tickets_NewTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketsServer).NewTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tickets.Tickets/NewTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketsServer).NewTicket(ctx, req.(*TicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tickets_GetTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketsServer).GetTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tickets.Tickets/GetTickets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketsServer).GetTickets(ctx, req.(*TicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tickets_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Tickets.Tickets",
	HandlerType: (*TicketsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewTicket",
			Handler:    _Tickets_NewTicket_Handler,
		},
		{
			MethodName: "GetTickets",
			Handler:    _Tickets_GetTickets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tickets.proto",
}

func init() { proto.RegisterFile("tickets.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0xc9, 0x4c, 0xce,
	0x4e, 0x2d, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x0f, 0x81, 0x70, 0x95, 0x62,
	0xb9, 0x78, 0x21, 0xcc, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x11, 0x2e, 0xd6, 0xd4,
	0xdc, 0xc4, 0xcc, 0x1c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0x48, 0x8a, 0x8b,
	0xa3, 0x38, 0x3f, 0x39, 0x33, 0x31, 0x27, 0x33, 0x45, 0x82, 0x49, 0x81, 0x51, 0x83, 0x39, 0x08,
	0xce, 0x07, 0xc9, 0x15, 0xa5, 0xa6, 0xa5, 0x16, 0x15, 0xa5, 0x16, 0x49, 0x30, 0x43, 0xe4, 0x60,
	0x7c, 0xa5, 0xd7, 0x8c, 0x5c, 0xdc, 0x30, 0xf3, 0x0b, 0x72, 0x2a, 0x41, 0x6a, 0x21, 0x0e, 0xc9,
	0x4c, 0x81, 0x5a, 0x00, 0xe7, 0x0b, 0x99, 0x72, 0xc1, 0x5c, 0x25, 0xc1, 0xa4, 0xc0, 0xac, 0xc1,
	0x6d, 0x24, 0xad, 0x07, 0xe5, 0xeb, 0x21, 0x19, 0x01, 0x63, 0xb3, 0x43, 0x3d, 0x24, 0xd5, 0xc1,
	0xc8, 0xc5, 0x06, 0x11, 0x03, 0x99, 0x0e, 0x61, 0x79, 0xba, 0xa0, 0x9a, 0xee, 0xe9, 0x02, 0xf2,
	0x97, 0x2b, 0xd8, 0x5f, 0x4c, 0x68, 0xfe, 0x0a, 0x06, 0xfb, 0xc3, 0xd3, 0x05, 0xe6, 0xf6, 0x62,
	0x28, 0x1f, 0x24, 0x17, 0x04, 0xf3, 0x17, 0x0b, 0xaa, 0xbf, 0x40, 0xa6, 0x39, 0xe5, 0xe7, 0x95,
	0x16, 0x4b, 0xb0, 0x2a, 0x30, 0x6a, 0x70, 0x04, 0xb1, 0x26, 0x81, 0x38, 0x46, 0x2d, 0x8c, 0x70,
	0x2f, 0x08, 0x59, 0x73, 0x71, 0xfa, 0xa5, 0x96, 0x43, 0x1d, 0x26, 0x86, 0xe1, 0x13, 0x70, 0x60,
	0x4b, 0x89, 0x60, 0xf3, 0xa1, 0x12, 0x83, 0x90, 0x0d, 0x17, 0x97, 0x7b, 0x6a, 0x09, 0xcc, 0x28,
	0x12, 0x75, 0x27, 0xb1, 0x81, 0xe3, 0xd8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x65, 0x28,
	0x39, 0xf4, 0x01, 0x00, 0x00,
}
