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
	WinnerRequest
	WinnerReply
	ClaimRequest
	ClaimReply
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
	Socialid string `protobuf:"bytes,2,opt,name=socialid" json:"socialid,omitempty"`
	Referrer string `protobuf:"bytes,3,opt,name=referrer" json:"referrer,omitempty"`
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

func (m *TicketRequest) GetSocialid() string {
	if m != nil {
		return m.Socialid
	}
	return ""
}

func (m *TicketRequest) GetReferrer() string {
	if m != nil {
		return m.Referrer
	}
	return ""
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
	Email    string `protobuf:"bytes,2,opt,name=Email,json=email" json:"-"`
	SocialID string `protobuf:"bytes,3,opt,name=SocialID,json=socialID" json:"SocialID,omitempty"`
	Referrer string `protobuf:"bytes,4,opt,name=Referrer,json=referrer" json:"Referrer,omitempty"`
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

func (m *TicketReply_Ticket) GetSocialID() string {
	if m != nil {
		return m.SocialID
	}
	return ""
}

func (m *TicketReply_Ticket) GetReferrer() string {
	if m != nil {
		return m.Referrer
	}
	return ""
}

func (m *TicketReply_Ticket) GetBonus() bool {
	if m != nil {
		return m.Bonus
	}
	return false
}

// The request message containing the user's name.
type WinnerRequest struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *WinnerRequest) Reset()                    { *m = WinnerRequest{} }
func (m *WinnerRequest) String() string            { return proto.CompactTextString(m) }
func (*WinnerRequest) ProtoMessage()               {}
func (*WinnerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *WinnerRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type WinnerReply struct {
	Winners []*WinnerReply_Winner `protobuf:"bytes,1,rep,name=Winners,json=winners" json:"Winners,omitempty"`
}

func (m *WinnerReply) Reset()                    { *m = WinnerReply{} }
func (m *WinnerReply) String() string            { return proto.CompactTextString(m) }
func (*WinnerReply) ProtoMessage()               {}
func (*WinnerReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *WinnerReply) GetWinners() []*WinnerReply_Winner {
	if m != nil {
		return m.Winners
	}
	return nil
}

type WinnerReply_Winner struct {
	WinnerID      int64               `protobuf:"varint,1,opt,name=WinnerID,json=winnerID" json:"WinnerID,omitempty"`
	DateTime      string              `protobuf:"bytes,2,opt,name=DateTime,json=dateTime" json:"DateTime,omitempty"`
	Entrants      string              `protobuf:"bytes,3,opt,name=Entrants,json=entrants" json:"Entrants,omitempty"`
	WinningTicket *TicketReply_Ticket `protobuf:"bytes,4,opt,name=WinningTicket,json=winningTicket" json:"WinningTicket,omitempty"`
	Claimed       bool                `protobuf:"varint,5,opt,name=Claimed,json=claimed" json:"Claimed,omitempty"`
	MoneyPot      int64               `protobuf:"varint,6,opt,name=MoneyPot,json=moneyPot" json:"MoneyPot,omitempty"`
	Paid          bool                `protobuf:"varint,7,opt,name=Paid,json=paid" json:"Paid,omitempty"`
}

func (m *WinnerReply_Winner) Reset()                    { *m = WinnerReply_Winner{} }
func (m *WinnerReply_Winner) String() string            { return proto.CompactTextString(m) }
func (*WinnerReply_Winner) ProtoMessage()               {}
func (*WinnerReply_Winner) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *WinnerReply_Winner) GetWinnerID() int64 {
	if m != nil {
		return m.WinnerID
	}
	return 0
}

func (m *WinnerReply_Winner) GetDateTime() string {
	if m != nil {
		return m.DateTime
	}
	return ""
}

func (m *WinnerReply_Winner) GetEntrants() string {
	if m != nil {
		return m.Entrants
	}
	return ""
}

func (m *WinnerReply_Winner) GetWinningTicket() *TicketReply_Ticket {
	if m != nil {
		return m.WinningTicket
	}
	return nil
}

func (m *WinnerReply_Winner) GetClaimed() bool {
	if m != nil {
		return m.Claimed
	}
	return false
}

func (m *WinnerReply_Winner) GetMoneyPot() int64 {
	if m != nil {
		return m.MoneyPot
	}
	return 0
}

func (m *WinnerReply_Winner) GetPaid() bool {
	if m != nil {
		return m.Paid
	}
	return false
}

type ClaimRequest struct {
	SocialID string `protobuf:"bytes,1,opt,name=SocialID,json=socialID" json:"SocialID,omitempty"`
}

func (m *ClaimRequest) Reset()                    { *m = ClaimRequest{} }
func (m *ClaimRequest) String() string            { return proto.CompactTextString(m) }
func (*ClaimRequest) ProtoMessage()               {}
func (*ClaimRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ClaimRequest) GetSocialID() string {
	if m != nil {
		return m.SocialID
	}
	return ""
}

type ClaimReply struct {
	Success bool `protobuf:"varint,1,opt,name=Success,json=success" json:"Success,omitempty"`
}

func (m *ClaimReply) Reset()                    { *m = ClaimReply{} }
func (m *ClaimReply) String() string            { return proto.CompactTextString(m) }
func (*ClaimReply) ProtoMessage()               {}
func (*ClaimReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ClaimReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*TicketRequest)(nil), "Tickets.TicketRequest")
	proto.RegisterType((*TicketReply)(nil), "Tickets.TicketReply")
	proto.RegisterType((*TicketReply_Ticket)(nil), "Tickets.TicketReply.Ticket")
	proto.RegisterType((*WinnerRequest)(nil), "Tickets.WinnerRequest")
	proto.RegisterType((*WinnerReply)(nil), "Tickets.WinnerReply")
	proto.RegisterType((*WinnerReply_Winner)(nil), "Tickets.WinnerReply.Winner")
	proto.RegisterType((*ClaimRequest)(nil), "Tickets.ClaimRequest")
	proto.RegisterType((*ClaimReply)(nil), "Tickets.ClaimReply")
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
	GetWinners(ctx context.Context, in *WinnerRequest, opts ...grpc.CallOption) (*WinnerReply, error)
	ClaimWin(ctx context.Context, in *ClaimRequest, opts ...grpc.CallOption) (*ClaimReply, error)
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

func (c *ticketsClient) GetWinners(ctx context.Context, in *WinnerRequest, opts ...grpc.CallOption) (*WinnerReply, error) {
	out := new(WinnerReply)
	err := grpc.Invoke(ctx, "/Tickets.Tickets/GetWinners", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketsClient) ClaimWin(ctx context.Context, in *ClaimRequest, opts ...grpc.CallOption) (*ClaimReply, error) {
	out := new(ClaimReply)
	err := grpc.Invoke(ctx, "/Tickets.Tickets/ClaimWin", in, out, c.cc, opts...)
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
	GetWinners(context.Context, *WinnerRequest) (*WinnerReply, error)
	ClaimWin(context.Context, *ClaimRequest) (*ClaimReply, error)
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

func _Tickets_GetWinners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WinnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketsServer).GetWinners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tickets.Tickets/GetWinners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketsServer).GetWinners(ctx, req.(*WinnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tickets_ClaimWin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClaimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketsServer).ClaimWin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tickets.Tickets/ClaimWin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketsServer).ClaimWin(ctx, req.(*ClaimRequest))
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
		{
			MethodName: "GetWinners",
			Handler:    _Tickets_GetWinners_Handler,
		},
		{
			MethodName: "ClaimWin",
			Handler:    _Tickets_ClaimWin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tickets.proto",
}

func init() { proto.RegisterFile("tickets.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x94, 0xcd, 0x8a, 0x13, 0x41,
	0x10, 0xc7, 0xb7, 0xf3, 0xd5, 0xb3, 0x15, 0x73, 0x69, 0x57, 0x69, 0xc6, 0x4b, 0x18, 0x50, 0x82,
	0x87, 0x1c, 0x56, 0x16, 0x04, 0xbd, 0xa8, 0x59, 0x24, 0x07, 0x65, 0x99, 0x5d, 0xd8, 0x93, 0x87,
	0xd9, 0x99, 0x52, 0x1a, 0x93, 0x99, 0x38, 0xdd, 0x21, 0xe4, 0x0d, 0x7c, 0x1c, 0x5f, 0x4b, 0x7c,
	0x05, 0x0f, 0xd2, 0x5d, 0xdd, 0xa1, 0x27, 0x8a, 0x82, 0xb7, 0xfa, 0x4d, 0x7d, 0x74, 0xfd, 0xab,
	0x8a, 0x81, 0x89, 0x51, 0xe5, 0x67, 0x34, 0x7a, 0xbe, 0x69, 0x1b, 0xd3, 0x08, 0x7e, 0x43, 0x98,
	0x7d, 0x80, 0x09, 0x99, 0x39, 0x7e, 0xd9, 0xa2, 0x36, 0xe2, 0x0c, 0x86, 0xb8, 0x2e, 0xd4, 0x4a,
	0xb2, 0x29, 0x9b, 0x9d, 0xe6, 0x04, 0x22, 0x85, 0x44, 0x37, 0xa5, 0x2a, 0x56, 0xaa, 0x92, 0x3d,
	0xe7, 0x38, 0xb0, 0xf5, 0xb5, 0xf8, 0x11, 0xdb, 0x16, 0x5b, 0xd9, 0x27, 0x5f, 0xe0, 0xec, 0x07,
	0x83, 0x71, 0xa8, 0xbf, 0x59, 0xed, 0x6d, 0x2c, 0x35, 0xa2, 0x2a, 0xff, 0xc0, 0x81, 0xc5, 0x05,
	0x84, 0xae, 0x64, 0x6f, 0xda, 0x9f, 0x8d, 0xcf, 0x1f, 0xcd, 0x3d, 0xcf, 0xa3, 0x12, 0xc1, 0xe6,
	0x5e, 0x50, 0xfa, 0x95, 0xc1, 0x88, 0xbe, 0xd9, 0xea, 0x64, 0x2d, 0x17, 0xdd, 0xea, 0xcb, 0x85,
	0xd5, 0x75, 0xe9, 0x74, 0xf5, 0x8e, 0x74, 0x5d, 0x3b, 0x1d, 0xcb, 0x45, 0xe8, 0x5d, 0x7b, 0xb6,
	0xbe, 0x3c, 0xe8, 0x1a, 0x74, 0x75, 0xd9, 0x6a, 0xaf, 0x9b, 0x7a, 0xab, 0xe5, 0x70, 0xca, 0x66,
	0x49, 0x3e, 0xbc, 0xb3, 0x90, 0x3d, 0x86, 0xc9, 0xad, 0xaa, 0x6b, 0x6c, 0xff, 0x3a, 0xcc, 0xec,
	0x5b, 0x0f, 0xc6, 0x21, 0xce, 0x0e, 0xe5, 0x02, 0x38, 0xa1, 0x96, 0xec, 0x48, 0x78, 0x14, 0x16,
	0x6c, 0xbe, 0xa3, 0xd8, 0xf4, 0x3b, 0x83, 0x11, 0x7d, 0xb3, 0xad, 0x92, 0xe5, 0x85, 0xf7, 0xf3,
	0x64, 0xe7, 0xd9, 0xfa, 0x16, 0x85, 0xc1, 0x1b, 0xb5, 0xc6, 0xb0, 0xba, 0xca, 0xb3, 0xf5, 0x5d,
	0xd6, 0xa6, 0x2d, 0x6a, 0xa3, 0x83, 0x7c, 0xf4, 0x2c, 0x5e, 0x91, 0x18, 0x55, 0x7f, 0xa2, 0x66,
	0xdc, 0x0c, 0xfe, 0xb1, 0x94, 0xc9, 0x2e, 0xce, 0x10, 0x12, 0xf8, 0x9b, 0x55, 0xa1, 0xd6, 0x58,
	0xf9, 0x39, 0xf1, 0x92, 0xd0, 0x3e, 0xfc, 0xae, 0xa9, 0x71, 0x7f, 0xd5, 0x18, 0x39, 0xa2, 0x86,
	0xd7, 0x9e, 0x85, 0x80, 0xc1, 0x55, 0xa1, 0x2a, 0xc9, 0x5d, 0xca, 0x60, 0x53, 0xa8, 0x2a, 0x7b,
	0x0a, 0xf7, 0x5c, 0xa5, 0x30, 0xd8, 0x78, 0x6f, 0xac, 0xbb, 0xb7, 0xec, 0x09, 0x80, 0x8f, 0xb5,
	0xc3, 0x95, 0xc0, 0xaf, 0xb7, 0x65, 0x89, 0x5a, 0xbb, 0xc0, 0x24, 0xe7, 0x9a, 0xf0, 0xfc, 0x27,
	0x3b, 0x1c, 0x9c, 0x78, 0x01, 0xa7, 0xef, 0x71, 0xe7, 0xdb, 0x7e, 0xf8, 0x9b, 0x44, 0xf7, 0x68,
	0x7a, 0xf6, 0x27, 0xe9, 0xd9, 0x89, 0x78, 0x09, 0xf0, 0x16, 0x4d, 0x28, 0xf5, 0x7f, 0xd9, 0xfe,
	0x00, 0xa2, 0xec, 0xce, 0x25, 0x45, 0xd9, 0xd1, 0x49, 0x64, 0x27, 0xe2, 0x39, 0x24, 0x4e, 0xec,
	0xad, 0xaa, 0xc5, 0x83, 0x43, 0x4c, 0x3c, 0xab, 0xf4, 0xfe, 0xf1, 0x67, 0x97, 0x79, 0x37, 0x72,
	0x7f, 0x82, 0x67, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x37, 0x77, 0x61, 0x65, 0x1a, 0x04, 0x00,
	0x00,
}
