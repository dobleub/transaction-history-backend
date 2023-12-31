package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/handlers"
	"github.com/sethvargo/go-envconfig"

	"github.com/dobleub/transaction-history-backend/internal"
	"github.com/dobleub/transaction-history-backend/internal/config"
)

func main() {
	var env config.Config
	envconfig.ProcessWith(context.Background(), &env, envconfig.OsLookuper())

	r := internal.Routes(env)

	runtime_api := env.AWSConfig.Lambda
	if runtime_api != "" {
		log.Println("Starting up in Lambda Runtime")
		adapter := gorillamux.NewV2(r)
		lambda.Start(adapter.ProxyWithContext)
	} else {
		log.Println("Starting up on localhost at port ", env.Port)
		headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
		originsOk := handlers.AllowedOrigins([]string{"*"})
		methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

		srv := &http.Server{
			Addr:    ":" + env.Port,
			Handler: handlers.CORS(originsOk, headersOk, methodsOk)(r),
		}
		_ = srv.ListenAndServe()
	}
}
