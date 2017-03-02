package users

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/bobbydeveaux/dynamo"
	user "github.com/bobbydeveaux/efr-service-go/proto/users"
	"os"
)

type Users struct {
}

var db = dynamo.New(session.New(), &aws.Config{
	Endpoint:                      aws.String(os.Getenv("DYNAMO_ADDR")),
	CredentialsChainVerboseErrors: aws.Bool(true),
	Region: aws.String("eu-west-1")})
var table = db.Table("Users")

func (a *Users) UpdateUser(user *user.User) *user.User {

	go table.Put(user).Run()

	return user
}
