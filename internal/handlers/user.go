package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dobleub/transaction-history-backend/internal/config"
	"github.com/dobleub/transaction-history-backend/internal/models"
)

/*
 * HandleUsers
 * Returns all users with it's total balance
 * @param w http.ResponseWriter
 * @param r *http.Request
 * @return void
 *
 * Sample request:
 * curl -X GET http://localhost:8080/users
 *
 * Process:
 * 1. Read CSV file
 * 2. Loop through CSV file
 * 3. Create user object
 * 4. Add user to users array
 */
func HandleUsers(env *config.Config, w http.ResponseWriter, r *http.Request) {

	users, err := models.GetUsers(env)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
