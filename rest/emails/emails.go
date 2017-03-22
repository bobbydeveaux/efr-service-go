package emails

import (
	"net/http"

	"github.com/bobbydeveaux/efr-service-go/email"
)

const (
	address = "localhost:50051"
)

func SendEmail(w http.ResponseWriter, r *http.Request) {

	email.SendDailyReminderEmail("me@bobbyjason.co.uk", 10)
}
