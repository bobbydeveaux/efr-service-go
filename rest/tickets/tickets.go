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

	"github.com/bobbydeveaux/efr-service-go/email"

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

var db = dynamo.New(session.New(), &aws.Config{
	Endpoint:                      aws.String(os.Getenv("DYNAMO_ADDR")),
	CredentialsChainVerboseErrors: aws.Bool(true),
	Region: aws.String("eu-west-1")})

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

func BonusTicket(w http.ResponseWriter, r *http.Request) {

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

	fmt.Println("BONUS TICKET")
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

	rpc, err := c.BonusTicket(context.Background(), &pb.TicketRequest{Email: User.GetEmail(), Socialid: User.GetSocialID(), Referrer: referrer})
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

	var fullcount = false
	if r.URL.Query().Get("fullcount") == "true" {
		fullcount = true
		fmt.Println("FULL COUNT")
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

	rpc, err := c.GetTickets(context.Background(), &pb.TicketRequest{Socialid: User.GetSocialID(), Fullcount: fullcount})
	if err != nil {
		fmt.Println("could not greet: %s", err)
	}

	var b = []byte("[]")

	if fullcount == true {
		b = []byte("{\"count\":" + strconv.Itoa(len(rpc.Tickets)) + "}")
	} else {
		b, err = json.Marshal(rpc.Tickets)
	}

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

	rpc, err := c.GetWinners(context.Background(), &pb.WinnerRequest{Email: "roger@gmail.com"})
	if err != nil {
		fmt.Println("could not greet: %s", err)
	}

	if User.GetSocialID() == "FB10153502750990419" {
		for _, v := range rpc.Winners {
			v.WinningTicket.PrivateEmail = v.WinningTicket.Email
		}
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

	deleteExpiredTickets()

	emailUsers(moneyPot)

}

func emailUsers(moneyPot int64) {
	tblUsers := db.Table("Users")

	var users []pbu.User
	err := tblUsers.Scan().All(&users)
	if err != nil {
		fmt.Printf(err.Error())
	}

	for _, u := range users {
		fmt.Println(u.GetEmail())
		go email.SendDailyReminderEmail("me@bobbyjason.co.uk", int(moneyPot))
	}

}

func deleteExpiredTickets() {
	tblTickets := db.Table("Tickets")

	var tickets []pb.TicketReply_Ticket
	err := tblTickets.Scan().Filter("DaysRemaining = ?", 1).All(&tickets)
	if err != nil {
		fmt.Printf(err.Error())
	}

	for _, t := range tickets {
		fmt.Println(t.GetTicketID())
		go tblTickets.Delete("TicketID", t.GetTicketID()).Run()
	}

}

// I don't think there is a valid exxcuse for this.
// but its 22:14 on 8th March 2017 and im tired.
// this needs to work before i go to bed.
func MarkAsPaid(w http.ResponseWriter, r *http.Request) {

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

	var b = []byte("[]")
	if User.GetSocialID() != "FB10153502750990419" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(b)
		return
	}

	tblWinners := db.Table("Winners")

	winnerid := r.URL.Query().Get("winnerid")
	int64WinnerID, _ := strconv.ParseInt(winnerid, 10, 64)
	var winners []pb.WinnerReply_Winner
	err = tblWinners.Scan().Filter("WinnerID = ?", int64WinnerID).All(&winners)
	if err != nil {
		fmt.Printf(err.Error())
	}

	winner := winners[0]

	if winner.Paid == true {
		b = []byte("{\"failure\": \"already paid\"}")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(b)
		return
	}
	winner.Paid = true

	tblUsers := db.Table("Users")

	var users []pbu.User
	err = tblUsers.Scan().Filter("SocialID = ?", winner.WinningTicket.SocialID).All(&users)
	if err != nil {
		fmt.Printf(err.Error())
	}

	u := users[0]
	fmt.Println(u)
	u.Balance = u.Balance - winner.MoneyPot

	tblWinners.Put(winner).Run()
	tblUsers.Put(u).Run()

	b = []byte("{\"success\": \"paid\"}")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(b)
	return
}
