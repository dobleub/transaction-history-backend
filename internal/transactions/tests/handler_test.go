package tests

import (
	"io"
	"net/http"
	"testing"

	"github.com/dobleub/transaction-history-backend/internal/helpers"
	"github.com/dobleub/transaction-history-backend/internal/transactions"
)

// HandleVersion is a test function for the transactions package
func TestHandlerVersion(t *testing.T) {
	var tests []struct {
		Name     string         `json:"name"`
		Desc     string         `json:"desc"`
		Req      *http.Request  `json:"req"`
		Res      *http.Response `json:"res"`
		Expected *http.Response `json:"expected"`
	}

	err := helpers.ReadJsonFile("internal/transactions/tests/handlerVersion.json", tests)
	if err != nil {
		t.Error(err)
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Error(err)
			}

			rr, err := helpers.ServeHandlerFunc(transactions.HandleVersion, req)
			if err != nil {
				t.Error(err)
			}

			// Check the response body is what we expect.
			expected, err := io.ReadAll(tt.Expected.Body)
			if err != nil {
				t.Error(err)
			}
			if rr.Body.String() != string(expected) {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tt.Expected.Body)
			}
		})
	}
}
