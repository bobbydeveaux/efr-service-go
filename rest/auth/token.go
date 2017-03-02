package auth

import (
	"encoding/json"

	"fmt"
	pb "github.com/bobbydeveaux/efr-service-go/proto/user"
	"github.com/dvsekhvalnov/jose2go"
	fb "github.com/huandu/facebook"
	"net/http"
	"time"
)

const passphrase string = "arse string"

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

	res, _ := fb.Get("/me", fb.Params{
		"fields":       "id,name,email,age_range,birthday,currency,first_name,devices,about,friends",
		"access_token": accessToken,
	})

	if res["id"] == nil || res["email"] == nil {
		var b = []byte("{\"error\": \"invalid access token\"}")

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
	}

	updateUser(User)

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

func updateUser(user *pb.User) {

}
