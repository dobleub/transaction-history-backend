package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
	"github.com/sethvargo/go-envconfig"

	"github.com/dobleub/transaction-history-backend/internal/config"
	"github.com/dobleub/transaction-history-backend/internal/transactions"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	var env config.Config
	envconfig.ProcessWith(context.Background(), &env, envconfig.OsLookuper())

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Not found", r.RequestURI)
		http.Error(w, fmt.Sprintf("Not found: %s", r.RequestURI), http.StatusNotFound)
	})

	// s := r.PathPrefix("/default").Subrouter()	// could use when run with API Gateway
	r.HandleFunc("/", transactions.HandleVersion).Methods(http.MethodGet)
	r.Use(loggingMiddleware)

	runtime_api := env.AWSConfig.Lambda
	if runtime_api != "" {
		log.Println("Starting up in Lambda Runtime")
		adapter := gorillamux.NewV2(r)
		lambda.Start(adapter.ProxyWithContext)
	} else {
		log.Println("Starting up on localhost")
		srv := &http.Server{
			Addr:    ":" + env.Port,
			Handler: r,
		}
		_ = srv.ListenAndServe()
	}
}
