package tickets

import (
	"encoding/json"

	"fmt"
	"net/http"

	pb "github.com/bobbydeveaux/efr-service-go/proto/tickets"
	pbu "github.com/bobbydeveaux/efr-service-go/proto/user"
	"github.com/bobbydeveaux/efr-service-go/rest/auth"
	"github.com/dvsekhvalnov/jose2go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"strconv"
)

const (
	address = "localhost:50051"
)

func NewTicket(w http.ResponseWriter, r *http.Request) {

	var err error

	fmt.Println("GET params were:", r.URL.Query())

	jwt := r.URL.Query().Get("jwt")
	if jwt == "" {
		b := []byte("[]")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(b)
		return
	}

	referrer := r.URL.Query().Get("referrer")

	int64referrer, _ := strconv.ParseInt(referrer, 10, 64)

	// decode the jwt to grab the email.
	passphrase := auth.GetPassphrase()

	strPayload, _, err := jose.Decode(jwt, passphrase)
	payload := []byte(strPayload)

	var User pbu.User
	err = json.Unmarshal(payload, &User)

	if err != nil {
		fmt.Println(err.Error())
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTicketsClient(conn)

	// Contact the server and print out its response.

	rpc, err := c.NewTicket(context.Background(), &pb.TicketRequest{Email: User.GetEmail(), Socialid: User.GetSocialID(), Referrer: int64referrer})
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

	var err error

	jwt := r.URL.Query().Get("jwt")
	if jwt == "" {
		b := []byte("[]")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(b)
		return
	}

	// decode the jwt to grab the email.
	passphrase := auth.GetPassphrase()

	strPayload, _, err := jose.Decode(jwt, passphrase)
	payload := []byte(strPayload)

	var User pbu.User
	err = json.Unmarshal(payload, &User)

	if err != nil {
		fmt.Println(err.Error())
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTicketsClient(conn)

	// Contact the server and print out its response.

	rpc, err := c.GetTickets(context.Background(), &pb.TicketRequest{Socialid: User.GetSocialID()})
	if err != nil {
		fmt.Println("could not greet: %s", err)
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

func GetWinners(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GET params were:", r.URL.Query())

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTicketsClient(conn)

	// Contact the server and print out its response.

	rpc, err := c.GetWinners(context.Background(), &pb.WinnerRequest{Email: "roger@gmail.com"})
	if err != nil {
		fmt.Println("could not greet: %s", err)
	}

	b, err := json.Marshal(rpc.Winners)
	if err != nil {
		fmt.Println("error:", err)
	}

	if string(b) == "null" {
		b = []byte("[]")
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(b)
}
