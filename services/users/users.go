package users

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/bobbydeveaux/dynamo"
	user "github.com/bobbydeveaux/efr-service-go/proto/users"
	"log"
	"os"
)

type Users struct {
}

var db = dynamo.New(session.New(), &aws.Config{
	Endpoint:                      aws.String(os.Getenv("DYNAMO_ADDR")),
	CredentialsChainVerboseErrors: aws.Bool(true),
	Region: aws.String("eu-west-1")})
var table = db.Table("Users")

func (a *Users) UpdateUser(u2 *user.User) *user.User {

	log.Println("[DEBUG] User before update")
	u1 := a.GetUser(u2.GetSocialID())
	fmt.Println("u1=", u1)
	fmt.Println("u2=", u2)

	// if u1 is esmpty, user doest exist so create a new one
	if u1.GetSocialID() == "" {
		u1 = u2
		u1.Balance = 0
	} else {
		// iser exists, only update balance & last login
		u1.LastLogin = u2.GetLastLogin()
		if u2.Balance >= 0 {
			u1.Balance = u2.GetBalance()
		}
	}

	table.Put(u1).Run()

	log.Println("[DEBUG] User after update")
	fmt.Println("u1=", u1)
	return u1
}

func (a *Users) GetUser(SocialID string) *user.User {

	var exist []*user.User

	err := table.Scan().Filter("SocialID = ?", SocialID).Consistent(true).All(&exist)
	if err != nil {
		log.Println(err.Error())
	}

	if len(exist) == 0 {
		u := &user.User{}
		return u
	}

	return exist[0]
}
