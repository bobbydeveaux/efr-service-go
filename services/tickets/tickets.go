package tickets

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/bobbydeveaux/dynamo"
	pb "github.com/bobbydeveaux/efr-service-go/proto/tickets"
	"github.com/bobbydeveaux/efr-service-go/services/users"
	"github.com/bobbydeveaux/randomizr"
	"log"
	"os"
	"sort"
)

type Tickets struct {
}

var db = dynamo.New(session.New(), &aws.Config{
	Endpoint:                      aws.String(os.Getenv("DYNAMO_ADDR")),
	CredentialsChainVerboseErrors: aws.Bool(true),
	Region: aws.String("eu-west-1")})
var table = db.Table("Tickets")

func (a *Tickets) NewTicket(email string, socialID string, referrer string) []*pb.TicketReply_Ticket {

	log.Println("New Ticket")
	var exist []*pb.TicketReply_Ticket
	//err := table.Get("Email", email).All(&exist)
	err := table.Scan().Filter("Email = ?", email).Consistent(true).All(&exist)
	if err != nil {
		log.Println(err.Error())
	}

	// YUK - can't get the filter above to filter by !Bonus
	if len(exist) > 0 {
		for _, t := range exist {
			if t.GetBonus() {
				continue
			}

			var tickets = make([]*pb.TicketReply_Ticket, 1)
			tickets[0] = t
			return tickets
		}
	}

	// is FB ID even valid? Check against FB SDK API
	//@TODO

	// Is this a valid referrer?
	if valid, referrerEmail := checkReferrer(referrer); valid == false {
		referrer = ""
	} else {
		// Give the referra their bonus ticket
		go referralTicket(referrer, referrerEmail)
	}

	// retrieve new ID from ID service
	// @TODO
	ticketID := randomizr.Generate(22)

	// Make the ticket
	ticket := &pb.TicketReply_Ticket{
		TicketID: ticketID,
		Email:    email,
		SocialID: socialID,
		Referrer: referrer,
		Bonus:    false,
	}

	log.Println("Creating ticket...")
	log.Println(ticket)

	// Insert into dynamo
	go table.Put(ticket).Run()

	var tickets = make([]*pb.TicketReply_Ticket, 1)
	tickets[0] = ticket
	return tickets
}

func (a *Tickets) BonusTicket(email string, socialID string, referrer string) []*pb.TicketReply_Ticket {

	log.Println("Bonus Ticket")

	var exist []*pb.TicketReply_Ticket

	err := table.Scan().Filter("DaysRemaining = ? AND SocialID = ?", 1, socialID).Consistent(true).All(&exist)
	if err != nil {
		log.Println(err.Error())
	}
	if len(exist) >= 10 {
		var tickets = make([]*pb.TicketReply_Ticket, 1)
		tickets[0] = &pb.TicketReply_Ticket{}
		return tickets
	}

	// retrieve new ID from ID service
	// @TODO
	ticketID := randomizr.Generate(22)

	// Make the ticket
	ticket := &pb.TicketReply_Ticket{
		TicketID:      ticketID,
		Email:         email,
		SocialID:      socialID,
		Referrer:      referrer,
		Bonus:         true,
		DaysRemaining: 1,
	}

	log.Println("Creating ticket...")
	log.Println(ticket)

	// Insert into dynamo
	go table.Put(ticket).Run()

	var tickets = make([]*pb.TicketReply_Ticket, 1)
	tickets[0] = ticket
	return tickets
}

func (a *Tickets) GetTickets(socialID string, fullcount bool) []*pb.TicketReply_Ticket {
	var tickets []*pb.TicketReply_Ticket
	var err error
	if fullcount == true {
		err = table.Scan().Consistent(true).All(&tickets)
	} else {
		err = table.Scan().Filter("SocialID = ?", socialID).Consistent(true).All(&tickets)
	}

	if err != nil {
		log.Println(err.Error())
	}

	return tickets
}

// By is the type of a "less" function that defines the ordering of its Planet arguments.
type By func(w1, w2 *pb.WinnerReply_Winner) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(winners []*pb.WinnerReply_Winner) {
	ps := &winnerSorter{
		winners: winners,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

type winnerSorter struct {
	winners []*pb.WinnerReply_Winner
	by      func(w1, w2 *pb.WinnerReply_Winner) bool
}

// Len is part of sort.Interface.
func (s *winnerSorter) Len() int {
	return len(s.winners)
}

// Swap is part of sort.Interface.
func (s *winnerSorter) Swap(i, j int) {
	s.winners[i], s.winners[j] = s.winners[j], s.winners[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *winnerSorter) Less(i, j int) bool {
	return s.by(s.winners[i], s.winners[j])
}

func (a *Tickets) GetWinners() []*pb.WinnerReply_Winner {
	var winners []*pb.WinnerReply_Winner
	table := db.Table("Winners")
	err := table.Scan().Consistent(true).All(&winners)
	if err != nil {
		log.Println(err.Error())
	}

	dateTime := func(w1, w2 *pb.WinnerReply_Winner) bool {
		return w1.WinnerID > w2.WinnerID
	}

	By(dateTime).Sort(winners)

	return winners
}

func (a *Tickets) ClaimWin(socialID string) bool {

	winners := a.GetWinners()
	if winners[0].WinningTicket.GetSocialID() != socialID {
		return false
	}

	if winners[0].Claimed == true {
		return false
	}

	if winners[0].Claimed == false {
		u := new(users.Users)
		user := u.GetUser(socialID)
		log.Println("Updating balance with money pot of ", winners[0].MoneyPot)
		user.Balance = user.Balance + winners[0].MoneyPot
		u.UpdateUser(user)
		log.Println(user)
		log.Println("Updated")
	}

	log.Println("Claiming Win")
	winners[0].Claimed = true

	table := db.Table("Winners")
	go table.Put(winners[0]).Run()

	return true
}

func checkReferrer(referrer string) (bool, string) {
	var exist []pb.TicketReply_Ticket
	log.Println("Checking referrer")

	log.Println(exist)

	//err := table.Get("Email", email).All(&exist)
	err := table.Scan().Filter("SocialID = ?", referrer).Consistent(true).All(&exist)

	if err != nil {
		log.Println(err.Error())
	}

	// referrer no exist
	if len(exist) > 0 {
		return true, exist[0].GetEmail()
	}

	log.Println("Invalid referrer: " + referrer)
	return false, ""

}

func referralTicket(referrer string, referrerEmail string) {

	ticketID := randomizr.Generate(22)

	// Make the ticket
	ticket := &pb.TicketReply_Ticket{
		TicketID: ticketID,
		Email:    referrerEmail,
		SocialID: referrer,
		Referrer: "",
		Bonus:    true,
	}

	// Insert into dynamo
	table.Put(ticket).Run()

	// Connect to SES and inforn referrer of their bonus ticket

}
