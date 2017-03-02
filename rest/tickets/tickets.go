package tickets

import (
	"encoding/json"

	"fmt"
	"net/http"

	pb "github.com/bobbydeveaux/efr-service-go/proto/tickets"
	pbu "github.com/bobbydeveaux/efr-service-go/proto/users"
	"github.com/bobbydeveaux/efr-service-go/rest/auth"
	s "github.com/bobbydeveaux/efr-service-go/services/tickets"
	"github.com/dvsekhvalnov/jose2go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"math/rand"
	"os"
	"strconv"
	"time"
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

	fmt.Println("WAAT")
	// decode the jwt to grab the email.
	passphrase := auth.GetPassphrase()

	strPayload, _, err := jose.Decode(jwt, passphrase)
	payload := []byte(strPayload)

	var User pbu.User
	err = json.Unmarshal(payload, &User)

	if err != nil {
		fmt.Println(err.Error())
	}

	if err != nil || User.GetEmail() == "" {
		var b = []byte("[]")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(b)
		return
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTicketsClient(conn)

	// Contact the server and print out its response.

	rpc, err := c.NewTicket(context.Background(), &pb.TicketRequest{Email: User.GetEmail(), Socialid: User.GetSocialID(), Referrer: referrer})
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

func ClaimWin(w http.ResponseWriter, r *http.Request) {

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

	rpc, err := c.ClaimWin(context.Background(), &pb.ClaimRequest{SocialID: User.GetSocialID()})

	b, err := json.Marshal(rpc)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(b)
	return

}

const PRIZE = 10

// THis really shouldnt be here. Horrible.
func PickWinner(w http.ResponseWriter, r *http.Request) {

	db := dynamo.New(session.New(), &aws.Config{
		Endpoint:                      aws.String(os.Getenv("DYNAMO_ADDR")),
		CredentialsChainVerboseErrors: aws.Bool(true),
		Region: aws.String("eu-west-1")})
	tblTickets := db.Table("Tickets")
	tblWinners := db.Table("Winners")

	var tickets []pb.TicketReply_Ticket
	err := tblTickets.Scan().All(&tickets)
	if err != nil {
		fmt.Printf(err.Error())
	}

	var moneyPot int64 = PRIZE
	tk := new(s.Tickets)
	winners := tk.GetWinners()
	if len(winners) > 0 {
		pastWinner := winners[0]
		fmt.Println(pastWinner)
		if !pastWinner.Claimed {
			moneyPot = moneyPot + pastWinner.MoneyPot
		}
	}

	fmt.Printf("We have %d raffle tickets in the pot\n", len(tickets))

	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	var luckyTicket = &tickets[rd.Intn(len(tickets))]
	fmt.Println("Magic 8-Ball says:", luckyTicket.GetEmail())

	var TimeNow = time.Now().Unix()
	strTimeNow := strconv.FormatInt(TimeNow, 10)
	var winner = &pb.WinnerReply_Winner{
		WinnerID:      TimeNow,
		DateTime:      strTimeNow,
		Entrants:      strconv.Itoa(len(tickets)),
		WinningTicket: luckyTicket,
		Claimed:       false,
		MoneyPot:      moneyPot,
	}
	tblWinners.Put(winner).Run()
}
