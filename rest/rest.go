package rest

import (
	"flag"

	"net/http"
	"time"

	"github.com/bobbydeveaux/efr-service-go/rest/auth"
	"github.com/bobbydeveaux/efr-service-go/rest/tickets"
	"github.com/gorilla/mux"
)

func Serve() {
	flag.Parse()
	router := mux.NewRouter()
	http.Handle("/", httpInterceptor(router))

	router.HandleFunc("/newticket", tickets.NewTicket).Methods("GET")
	router.HandleFunc("/gettickets", tickets.GetTickets).Methods("GET")
	router.HandleFunc("/bonusticket", tickets.BonusTicket).Methods("GET")
	router.HandleFunc("/token", auth.GetToken).Methods("GET")
	router.HandleFunc("/winners", tickets.GetWinners).Methods("GET")
	router.HandleFunc("/claim", tickets.ClaimWin).Methods("GET")
	router.HandleFunc("/markaspaid", tickets.MarkAsPaid).Methods("GET")
	router.HandleFunc("/ILrjBUFER1QuA1jLLYmc", tickets.PickWinner).Methods("GET")

	http.ListenAndServe(":8181", nil)
}

func httpInterceptor(router http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		if origin := req.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers",
				"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		}
		// Stop here if its Preflighted OPTIONS request
		if req.Method == "OPTIONS" {
			return
		}
		startTime := time.Now()

		router.ServeHTTP(w, req)

		finishTime := time.Now()
		elapsedTime := finishTime.Sub(startTime)

		switch req.Method {
		case "GET":
			// We may not always want to StatusOK, but for the sake of
			// this example we will
			LogAccess(w, req, elapsedTime)
		case "POST":
			// here we might use http.StatusCreated
		}

	})
}
