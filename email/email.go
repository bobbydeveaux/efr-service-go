package email

import (
	"fmt"
	"github.com/sourcegraph/go-ses"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func sendEmail(to string, subject string, plainText string, htmlBody string) {
	// Change the From address to a sender address that is verified in your Amazon SES account.
	from := "EasyFreeRaffle.com<bobby@easyfreeraffle.com>"

	// EnvConfig uses the AWS credentials in the environment variables $AWS_ACCESS_KEY_ID and
	// $AWS_SECRET_KEY.
	res, err := ses.EnvConfig.SendEmailHTML(from, to, subject, plainText, htmlBody)
	if err == nil {
		fmt.Printf("Sent email: %s...\n", res[:32])
	} else {
		fmt.Printf("Error sending email: %s\n", err)
	}

	// output:
	// Sent email: <SendEmailResponse xmlns="http:/...
}

func SendDailyReminderEmail(to string, prize int) {
	today := time.Now().Weekday().String()
	subject := today + " Draw Reminder"
	fmt.Println(subject)

	pwd, _ := os.Getwd()
	b, _ := ioutil.ReadFile(pwd + "/email/templates/reminder.html")

	body := strings.Replace(string(b), "{{prize}}", strconv.Itoa(prize), -1)
	body = strings.Replace(body, "{{day}}", today, -1)
	sendEmail(to, subject, "", body)
}
