package tests

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/dobleub/transaction-history-backend/internal/config"
	"github.com/dobleub/transaction-history-backend/internal/handlers"
	"github.com/dobleub/transaction-history-backend/internal/helpers"
	"github.com/dobleub/transaction-history-backend/internal/models"
	"github.com/gorilla/mux"
	"github.com/sethvargo/go-envconfig"
	"github.com/stretchr/testify/assert"
)

var env *config.Config

func TestMain(m *testing.M) {
	var tmpEnv config.Config
	envconfig.ProcessWith(context.Background(), &tmpEnv, envconfig.OsLookuper())

	env = &tmpEnv
	m.Run()
}

// HandleVersion is a test function for the transactions package
func TestHandlerVersion(t *testing.T) {
	var tests []struct {
		Name string `json:"name"`
		Req  struct {
			URL    string `json:"url"`
			Method string `json:"method"`
		} `json:"req,omitempty"`
		Expected struct {
			Body   models.Version `json:"body"`
			Status int            `json:"status"`
		} `json:"expected,omitempty"`
	}
	err := helpers.ReadJsonFile("internal/handlers/tests/handlerVersion.json", &tests)
	if err != nil {
		t.Error(err)
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			req, err := http.NewRequest(tt.Req.Method, tt.Req.URL, nil)
			if err != nil {
				t.Error(err)
			}

			rr, err := helpers.ServeHandlerFunc(handlers.HandleVersion, req)
			if err != nil {
				t.Error(err)
			}

			// Check the response body is what we expect.
			if err != nil {
				t.Error(err)
			}

			response := strings.TrimSpace(rr.Body.String())
			if !assert.Equal(t, helpers.ObjectToJsonString(tt.Expected.Body), response) {
				t.Errorf("handler returned unexpected body: got %v want %v", response, tt.Expected.Body)
			}
		})
	}
}

func TestHandlerTransactions(t *testing.T) {
	var tests []struct {
		Name string `json:"name"`
		Req  struct {
			URL    string            `json:"url"`
			Method string            `json:"method"`
			Vars   map[string]string `json:"vars,omitempty"`
		} `json:"req,omitempty"`
		Expected struct {
			Body   []models.Transaction `json:"body"`
			Status int                  `json:"status"`
		} `json:"expected,omitempty"`
	}
	err := helpers.ReadJsonFile("internal/handlers/tests/handlerTransactions.json", &tests)
	if err != nil {
		t.Error(err)
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			req, err := http.NewRequest(tt.Req.Method, tt.Req.URL, nil)
			if err != nil {
				t.Error(err)
			}

			req = mux.SetURLVars(req, tt.Req.Vars)
			rr, err := helpers.ServeHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				handlers.HandleTransactions(env, w, r)
			}, req)
			if err != nil {
				t.Error(err)
			}

			// Check the response body is what we expect.
			if err != nil {
				t.Error(err)
			}

			response := strings.TrimSpace(rr.Body.String())
			if !assert.Equal(t, helpers.ObjectToJsonString(tt.Expected.Body), response) {
				t.Errorf("handler returned unexpected body: got %v want %v", response, tt.Expected.Body)
			}
		})
	}
}

func TestHandlerSummary(t *testing.T) {
	var tests []struct {
		Name string `json:"name"`
		Req  struct {
			URL    string            `json:"url"`
			Method string            `json:"method"`
			Vars   map[string]string `json:"vars,omitempty"`
		} `json:"req,omitempty"`
		Expected struct {
			Body   models.Summary `json:"body"`
			Status int            `json:"status"`
		} `json:"expected,omitempty"`
	}
	err := helpers.ReadJsonFile("internal/handlers/tests/handlerSummary.json", &tests)
	if err != nil {
		t.Error(err)
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			req, err := http.NewRequest(tt.Req.Method, tt.Req.URL, nil)
			if err != nil {
				t.Error(err)
			}

			req = mux.SetURLVars(req, tt.Req.Vars)
			rr, err := helpers.ServeHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				handlers.HandleSummary(env, w, r)
			}, req)
			if err != nil {
				t.Error(err)
			}

			// Check the response body is what we expect.
			if err != nil {
				t.Error(err)
			}

			response := strings.TrimSpace(rr.Body.String())
			if !assert.Equal(t, helpers.ObjectToJsonString(tt.Expected.Body), response) {
				t.Errorf("handler returned unexpected body: got %v want %v", response, tt.Expected.Body)
			}
		})
	}
}
