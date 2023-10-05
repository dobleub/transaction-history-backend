package helpers

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/dobleub/transaction-history-backend/internal/transactions"
)

func ServeHandlerFunc(handlerFunc http.HandlerFunc, req *http.Request) (*httptest.ResponseRecorder, error) {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(transactions.HandleVersion)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		return nil, fmt.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	return rr, nil
}
