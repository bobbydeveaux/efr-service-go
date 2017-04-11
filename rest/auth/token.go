package auth

import (
	"encoding/json"

	"bytes"
	"fmt"
	pb "github.com/bobbydeveaux/efr-service-go/proto/users"
	"github.com/dvsekhvalnov/jose2go"
	fb "github.com/huandu/facebook"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"time"
)

const passphrase string = "arse string"

const (
	address = "localhost:50051"
)

type Token struct {
	Token     string   `json:"token"`
	FirstName string   `json:"first_name"`
	SocialID  string   `json:"social_id"`
	User      *pb.User `json:"user"`
}

func GetPassphrase() string {
	return passphrase
}

func GetToken(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GET params were:", r.URL.Query())

	//socialNetwork := r.URL.Query().Get("social_network")
	accessToken := r.URL.Query().Get("access_token")
	referrer := r.URL.Query().Get("referrer")

	res, _ := fb.Get("/me", fb.Params{
		"fields":       "id,name,email,age_range,birthday,currency,first_name,devices,about,friends",
		"access_token": accessToken,
	})

	if res["id"] == nil || res["email"] == nil {
		var buffer bytes.Buffer
		buffer.WriteString("{\"error\": \"invalid access token\", \"email\":")
		buffer.WriteString("\"")
		buffer.WriteString(res["email"].(string))
		buffer.WriteString("\"")
		buffer.WriteString("}")

		var b = []byte(buffer.String())

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(b)
		return
	}

	var SocialID = "FB" + res["id"].(string)

	var User = &pb.User{
		SocialID:  SocialID,
		Email:     res["email"].(string),
		FirstName: res["first_name"].(string),
		Name:      res["name"].(string),
		LastLogin: time.Now().Unix(),
		Referrer:  referrer,
		Balance:   -1,
	}

	User = updateUser(User)

	payload, err := json.Marshal(User)
	strPayload := string(payload[:])
	fmt.Printf("payload is %v, ", strPayload)
	if err != nil {
		fmt.Println("error:", err)
	}

	secureToken, err := jose.Encrypt(strPayload, jose.PBES2_HS256_A128KW, jose.A256GCM, passphrase)

	fmt.Println("secureToken: ", secureToken)
	var token = &Token{
		Token:     secureToken,
		FirstName: User.GetFirstName(),
		SocialID:  SocialID,
		User:      User,
	}

	response, _ := json.Marshal(token)

	fmt.Println("response: ", string(response[:]))
	/*
		if err == nil {
			//go use token
			fmt.Printf("\nPBES2_HS256_A128KW A256GCM = %v\n", token)
		}
	*/

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(response)
}

func updateUser(user *pb.User) *pb.User {
	log.Println("Updating user")
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUsersClient(conn)

	// Contact the server and print out its response.

	rpc, err := c.UpdateUser(context.Background(), &pb.UserRequest{User: user})
	if err != nil {
		fmt.Println("could not greet: %v", err)
	}

	return rpc.User
}
