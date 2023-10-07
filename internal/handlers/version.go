package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dobleub/transaction-history-backend/internal/models"
)

/*
 * HandleVersion
 * Simple version handler to check if the API is running
 * @param w http.ResponseWriter
 * @param r *http.Request
 * @return void
 *
 */
func HandleVersion(w http.ResponseWriter, r *http.Request) {
	version := models.Version{}
	version.Version = "1.0.1 - 2021-09-30 - transactions-history"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(version)
}

/*
 * HandleNotFound
 * Simple not found handler
 * @param w http.ResponseWriter
 * @param r *http.Request
 * @return void
 *
 * Sample request:
 * curl -X GET http://localhost:8080/other-request
 *
 */
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Not found", r.RequestURI)
	http.Error(w, fmt.Sprintf("Not found: %s", r.RequestURI), http.StatusNotFound)
}
