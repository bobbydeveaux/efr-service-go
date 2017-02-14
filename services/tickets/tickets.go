package tickets

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/bobbydeveaux/dynamo"
	pb "github.com/bobbydeveaux/efr-service-go/proto/tickets"
	"github.com/bobbydeveaux/randomizr"
	"log"
	"strconv"
)

type Tickets struct {
}

var db = dynamo.New(session.New(), &aws.Config{
	Endpoint:                      aws.String("http://localhost:8000"),
	CredentialsChainVerboseErrors: aws.Bool(true),
	Region: aws.String("eu-west-1")})
var table = db.Table("Tickets")

func (a *Tickets) NewTicket(email string, facebookID int64, referrer int64) []*pb.TicketReply_Ticket {

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
		referrer = 0
	} else {
		// Give the referra their bonus ticket
		go referralTicket(referrer, referrerEmail)
	}

	// retrieve new ID from ID service
	// @TODO
	ticketID := randomizr.Generate(22)

	// Make the ticket
	ticket := &pb.TicketReply_Ticket{
		TicketID:   ticketID,
		Email:      email,
		FacebookID: facebookID,
		Referrer:   referrer,
		Bonus:      false,
	}

	// Insert into dynamo
	go table.Put(ticket).Run()

	var tickets = make([]*pb.TicketReply_Ticket, 1)
	tickets[0] = ticket
	return tickets
}

func (a *Tickets) GetTickets(facebookID int64) []*pb.TicketReply_Ticket {
	var tickets []*pb.TicketReply_Ticket
	//err := table.Get("Email", email).All(&exist)
	err := table.Scan().Filter("FacebookID = ?", facebookID).Consistent(true).All(&tickets)
	if err != nil {
		log.Println(err.Error())
	}

	return tickets
}

func checkReferrer(referrer int64) (bool, string) {
	var exist []pb.TicketReply_Ticket
	//err := table.Get("Email", email).All(&exist)
	err := table.Scan().Filter("FacebookID = ?", referrer).Consistent(true).All(&exist)

	if err != nil {
		log.Println(err.Error())
	}

	// referrer no exist
	if len(exist) > 0 {
		return true, exist[0].GetEmail()
	}

	log.Println("Invalid referrer: " + strconv.FormatInt(referrer, 10))
	return false, ""

}

func referralTicket(referrer int64, referrerEmail string) {

	ticketID := randomizr.Generate(22)

	// Make the ticket
	ticket := &pb.TicketReply_Ticket{
		TicketID:   ticketID,
		Email:      referrerEmail,
		FacebookID: referrer,
		Referrer:   0,
		Bonus:      true,
	}

	// Insert into dynamo
	table.Put(ticket).Run()

	// Connect to SES and inforn referrer of their bonus ticket

}
