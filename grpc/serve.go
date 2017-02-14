package grpc

import (
	"github.com/bobbydeveaux/efr-service-go/services/tickets"

	"log"
	"net"

	pb "github.com/bobbydeveaux/efr-service-go/proto/tickets"
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
	tickets := tk.NewTicket(in.GetEmail(), in.GetFacebookid(), in.GetReferrer())
	return &pb.TicketReply{Tickets: tickets}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *server) GetTickets(ctx context.Context, in *pb.TicketRequest) (*pb.TicketReply, error) {
	tk := new(tickets.Tickets)
	tickets := tk.GetTickets(in.GetFacebookid())
	return &pb.TicketReply{Tickets: tickets}, nil
}

func Serve() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTicketsServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
