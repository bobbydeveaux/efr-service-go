package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	pb "github.com/bobbydeveaux/efr-service-go/proto/tickets"
	s "github.com/bobbydeveaux/efr-service-go/services/tickets"
	"github.com/guregu/dynamo"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const PRIZE = 10

func test() {

	db := dynamo.New(session.New(), &aws.Config{
		Endpoint:                      aws.String(os.Getenv("DYNAMO_ADDR")),
		CredentialsChainVerboseErrors: aws.Bool(true),
		Region: aws.String("eu-west-1")})
	tblWinners := db.Table("Winners")
	var winners []*pb.WinnerReply_Winner
	//err := table.Get("Email", email).All(&exist)
	err := tblWinners.Scan().All(&winners)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("Winner count:", len(winners))

}

func pickWinner() {
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

	tk := new(s.Tickets)
	pastWinner := tk.GetWinners()[0]
	fmt.Println(pastWinner)

	var moneyPot int64 = PRIZE

	if !pastWinner.Claimed {
		moneyPot = moneyPot + pastWinner.MoneyPot
	}

	fmt.Printf("We have %d raffle tickets in the pot\n", len(tickets))

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var luckyTicket = &tickets[r.Intn(len(tickets))]
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
