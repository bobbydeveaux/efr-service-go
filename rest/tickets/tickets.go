package tickets

import (
	"encoding/json"

	"fmt"
	"net/http"

	pb "github.com/bobbydeveaux/efr-service-go/proto/tickets"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"strconv"
)

const (
	address = "localhost:50051"
)

func NewTicket(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GET params were:", r.URL.Query())

	email := r.URL.Query().Get("email")
	facebookid := r.URL.Query().Get("facebookid")
	referrer := r.URL.Query().Get("referrer")

	int64facebook, _ := strconv.ParseInt(facebookid, 10, 64)
	int64referrer, _ := strconv.ParseInt(referrer, 10, 64)
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTicketsClient(conn)

	// Contact the server and print out its response.

	rpc, err := c.NewTicket(context.Background(), &pb.TicketRequest{Email: email, Facebookid: int64facebook, Referrer: int64referrer})
	if err != nil {
		fmt.Println("could not greet: %v", err)
	}

	b, err := json.Marshal(rpc.Tickets)
	if err != nil {
		fmt.Println("error:", err)
	}

	if b == nil {
		b = []byte("[]")
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(b)
}

func GetTickets(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GET params were:", r.URL.Query())

	facebookid := r.URL.Query().Get("facebookid")

	int64facebook, _ := strconv.ParseInt(facebookid, 10, 64)

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTicketsClient(conn)

	// Contact the server and print out its response.

	rpc, err := c.GetTickets(context.Background(), &pb.TicketRequest{Facebookid: int64facebook})
	if err != nil {
		fmt.Println("could not greet: %v", err)
	}

	b, err := json.Marshal(rpc.Tickets)
	if err != nil {
		fmt.Println("error:", err)
	}

	if string(b) == "null" {
		b = []byte("[]")
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(b)
}
