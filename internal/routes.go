package internal

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dobleub/transaction-history-backend/internal/config"
	"github.com/dobleub/transaction-history-backend/internal/handlers"
	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Not found", r.RequestURI)
	http.Error(w, fmt.Sprintf("Not found: %s", r.RequestURI), http.StatusNotFound)
}

func Routes(env config.Config) *mux.Router {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	// s := r.PathPrefix("/default").Subrouter()	// could use when run with API Gateway
	r.HandleFunc("/", handlers.HandleVersion).Methods(http.MethodGet)
	r.HandleFunc("/transactions/user/{userid}", handlers.HandleTransactions).Methods(http.MethodGet)
	r.HandleFunc("/summary/user/{userid}", handlers.HandleSummary).Methods(http.MethodGet)
	r.HandleFunc("/summary/email/{userid}/to/{emailto}", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlerSendEmail(&env.EmailConfig, w, r)
	}).Methods(http.MethodGet)

	r.Use(loggingMiddleware)

	return r
}
