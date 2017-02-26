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

func (m *TicketRequest) GetSocialid() string {
	if m != nil {
		return m.Socialid
	}
	return ""
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
	SocialID string `protobuf:"bytes,3,opt,name=SocialID,json=socialID" json:"SocialID,omitempty"`
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

func (m *TicketReply_Ticket) GetSocialID() string {
	if m != nil {
		return m.SocialID
	}
	return ""
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
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x93, 0xcd, 0x8a, 0x13, 0x41,
	0x10, 0xc7, 0xb7, 0x13, 0x93, 0x9e, 0xad, 0x98, 0x4b, 0xbb, 0x4a, 0x33, 0x5e, 0x42, 0x83, 0x12,
	0x3c, 0xe4, 0xb0, 0xb2, 0x20, 0xe8, 0x45, 0x9d, 0x45, 0x72, 0x50, 0x64, 0x76, 0x61, 0x4f, 0x1e,
	0x66, 0x27, 0xa5, 0x34, 0x26, 0x3d, 0xeb, 0x74, 0x87, 0x61, 0xdf, 0xc0, 0x87, 0xf0, 0xb5, 0x7c,
	0x0a, 0x5f, 0xc1, 0x83, 0xf4, 0xd7, 0xd0, 0x13, 0x45, 0xc1, 0x5b, 0xfd, 0xbb, 0x3e, 0x7f, 0x55,
	0x33, 0x30, 0x37, 0xb2, 0xfe, 0x8c, 0x46, 0xaf, 0x6e, 0xda, 0xc6, 0x34, 0x8c, 0x5e, 0x7a, 0x29,
	0x3e, 0xc0, 0xdc, 0x9b, 0x25, 0x7e, 0xd9, 0xa3, 0x36, 0xec, 0x04, 0x26, 0xb8, 0xab, 0xe4, 0x96,
	0x93, 0x05, 0x59, 0x1e, 0x97, 0x5e, 0xb0, 0x1c, 0x32, 0xdd, 0xd4, 0xb2, 0xda, 0xca, 0x0d, 0x1f,
	0x39, 0x47, 0xaf, 0xad, 0xaf, 0xc5, 0x8f, 0xd8, 0xb6, 0xd8, 0xf2, 0xf1, 0x82, 0x2c, 0xc7, 0x65,
	0xaf, 0xc5, 0x0f, 0x02, 0xb3, 0x58, 0xff, 0x66, 0x7b, 0x6b, 0x63, 0xfd, 0x20, 0x72, 0x13, 0x1a,
	0xf4, 0x9a, 0x9d, 0x41, 0x9c, 0x8a, 0x8f, 0x16, 0xe3, 0xe5, 0xec, 0xf4, 0xe1, 0x2a, 0xe8, 0x55,
	0x52, 0x22, 0xda, 0x34, 0x00, 0xe5, 0x5f, 0x09, 0x4c, 0xfd, 0x9b, 0xad, 0xee, 0xad, 0x75, 0x31,
	0xac, 0xbe, 0x2e, 0x2c, 0xd7, 0xb9, 0xe3, 0x1a, 0x1d, 0x70, 0x5d, 0x38, 0x8e, 0x75, 0xe1, 0x66,
	0xef, 0xb9, 0xd6, 0x85, 0xf5, 0x95, 0x91, 0xeb, 0xce, 0x90, 0xcb, 0x56, 0x7b, 0xd5, 0xa8, 0xbd,
	0xe6, 0x93, 0x05, 0x59, 0x66, 0xe5, 0xe4, 0xda, 0x0a, 0xf1, 0x08, 0xe6, 0x57, 0x52, 0x29, 0x6c,
	0xff, 0xba, 0x4c, 0xf1, 0x6d, 0x04, 0xb3, 0x18, 0x67, 0x97, 0x72, 0x06, 0xd4, 0x4b, 0xcd, 0xc9,
	0x01, 0x78, 0x12, 0x16, 0x6d, 0xda, 0xf9, 0xd8, 0xfc, 0x3b, 0x81, 0xa9, 0x7f, 0xb3, 0xa3, 0x7a,
	0x2b, 0x80, 0x8f, 0xcb, 0xac, 0x0b, 0xda, 0xfa, 0x8a, 0xca, 0xe0, 0xa5, 0xdc, 0x61, 0x3c, 0xdd,
	0x26, 0x68, 0xeb, 0x3b, 0x57, 0xa6, 0xad, 0x94, 0xd1, 0x11, 0x1f, 0x83, 0x66, 0x2f, 0x3d, 0x8c,
	0x54, 0x9f, 0xfc, 0x30, 0x6e, 0x07, 0xff, 0x38, 0xca, 0xbc, 0x4b, 0x33, 0x18, 0x07, 0xfa, 0x7a,
	0x5b, 0xc9, 0x1d, 0x6e, 0xc2, 0x9e, 0x68, 0xed, 0xa5, 0x6d, 0xfc, 0xb6, 0x51, 0x78, 0xfb, 0xbe,
	0x31, 0x7c, 0xea, 0x07, 0xde, 0x05, 0x2d, 0x9e, 0xc0, 0x5d, 0x97, 0x15, 0x97, 0x98, 0xde, 0x88,
	0x0c, 0x6f, 0x24, 0x1e, 0x03, 0x84, 0x58, 0xbb, 0x48, 0x0e, 0xf4, 0x62, 0x5f, 0xd7, 0xa8, 0xb5,
	0x0b, 0xcc, 0x4a, 0xaa, 0xbd, 0x3c, 0xfd, 0x49, 0xfa, 0x8f, 0x8b, 0x3d, 0x87, 0xe3, 0x77, 0xd8,
	0x85, 0x11, 0x1f, 0xfc, 0x86, 0xe3, 0x9a, 0xe6, 0x27, 0x7f, 0xc2, 0x14, 0x47, 0xec, 0x05, 0xc0,
	0x1b, 0x34, 0xb1, 0xd4, 0xff, 0x65, 0x87, 0x63, 0x27, 0xd9, 0x83, 0xaf, 0x26, 0xc9, 0x4e, 0xce,
	0x2f, 0x8e, 0xd8, 0x33, 0xc8, 0x1c, 0xec, 0x95, 0x54, 0xec, 0x7e, 0x1f, 0x93, 0xee, 0x2a, 0xbf,
	0x77, 0xf8, 0xec, 0x32, 0xaf, 0xa7, 0xee, 0xaf, 0x7f, 0xfa, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x96,
	0xf3, 0x89, 0x10, 0x06, 0x04, 0x00, 0x00,
}
