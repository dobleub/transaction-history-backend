package internal

import (
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

func Routes(env config.Config) *mux.Router {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)

	// s := r.PathPrefix("/default").Subrouter()	// could use when run with API Gateway
	r.HandleFunc("/", handlers.HandleVersion).Methods(http.MethodGet)
	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleUsers(&env, w, r)
	}).Methods(http.MethodGet)
	r.HandleFunc("/transactions/user/{userid}", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleTransactions(&env, w, r)
	}).Methods(http.MethodGet)
	r.HandleFunc("/summary/user/{userid}", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleSummary(&env, w, r)
	}).Methods(http.MethodGet)
	r.HandleFunc("/summary/email/{userid}/to/{emailto}", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlerSendEmail(&env, w, r)
	}).Methods(http.MethodGet)

	r.Use(loggingMiddleware)

	return r
}
