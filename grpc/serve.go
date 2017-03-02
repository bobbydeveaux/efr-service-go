package grpc

import (
	"github.com/bobbydeveaux/efr-service-go/services/tickets"
	"github.com/bobbydeveaux/efr-service-go/services/users"

	"log"
	"net"

	pb "github.com/bobbydeveaux/efr-service-go/proto/tickets"
	user "github.com/bobbydeveaux/efr-service-go/proto/users"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) NewTicket(ctx context.Context, in *pb.TicketRequest) (*pb.TicketReply, error) {
	tk := new(tickets.Tickets)
	tickets := tk.NewTicket(in.GetEmail(), in.GetSocialid(), in.GetReferrer())
	return &pb.TicketReply{Tickets: tickets}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *server) GetTickets(ctx context.Context, in *pb.TicketRequest) (*pb.TicketReply, error) {
	tk := new(tickets.Tickets)
	tickets := tk.GetTickets(in.GetSocialid())
	return &pb.TicketReply{Tickets: tickets}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *server) GetWinners(ctx context.Context, in *pb.WinnerRequest) (*pb.WinnerReply, error) {
	tk := new(tickets.Tickets)
	winners := tk.GetWinners()
	return &pb.WinnerReply{Winners: winners}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *server) ClaimWin(ctx context.Context, in *pb.ClaimRequest) (*pb.ClaimReply, error) {
	tk := new(tickets.Tickets)
	claim := tk.ClaimWin(in.GetSocialID())
	return &pb.ClaimReply{Success: claim}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *server) UpdateUser(ctx context.Context, in *user.UserRequest) (*user.UserReply, error) {
	u := new(users.Users)
	ur := u.UpdateUser(in.GetUser())
	return &user.UserReply{User: ur}, nil
}

func Serve() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTicketsServer(s, &server{})
	user.RegisterUsersServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
