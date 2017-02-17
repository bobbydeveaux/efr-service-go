package gettickets

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	pb "github.com/bobbydeveaux/efr-service-go/proto/tickets"
	"github.com/guregu/dynamo"
	"strconv"
)

func main() {
	db := dynamo.New(session.New(), &aws.Config{
		Endpoint:                      aws.String("http://localhost:8000"),
		CredentialsChainVerboseErrors: aws.Bool(true),
		Region: aws.String("eu-west-1")})
	table := db.Table("Tickets")

	var tickets []pb.TicketReply_Ticket
	err := table.Scan().All(&tickets)
	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Println(len(tickets))

	for _, ticket := range tickets {
		fmt.Println(ticket.TicketID + " - " + ticket.Email + " - Bonus: " + strconv.FormatBool(ticket.Bonus))
	}
}
