package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dobleub/transaction-history-backend/internal/helpers"
	"github.com/dobleub/transaction-history-backend/internal/models"
	"github.com/gorilla/mux"
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
 * HandleTransactions
 * @param w http.ResponseWriter
 * @param r *http.Request
 * @return void
 *
 * Sample request:
 * curl -X GET http://localhost:8080/transactions/user/1
 *
 * Process:
 * 1. Read CSV file
 * 2. Loop through CSV file
 * 3. If user id matches, append to transactions slice
 */
func HandleTransactions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userid"]

	user := models.User{
		UserId: helpers.StringToInt32(userId),
	}

	transactions, err := user.GetTransactions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(transactions) == 0 {
		http.Error(w, "No transactions found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

/*
 * HandleSummary
 * @param w http.ResponseWriter
 * @param r *http.Request
 * @return void
 *
 * Sample request:
 * curl -X GET http://localhost:8080/summary/user/1
 *
 * Process:
 * 1. Get transactions for user from HandleTransactions
 * 3. Calc summary
 * 	- Total balance: sum of all transactions
 * 	- Total income: sum of all income transactions
 * 	- Total expense: sum of all expense transactions
 * 	- Transactions per month: number of transactions per month
 * 	- Average credit amount: average amount of all credit transactions maked by +amount
 * 	- Average debit amount: average amount of all debit transactions maked by -amount
 * 5. Return summary and transactions
 */

func HandleSummary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userid"]

	user := models.User{
		UserId: helpers.StringToInt32(userId),
	}

	summary, err := user.GetSummary()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(summary.TransactionsPerMonth) == 0 {
		http.Error(w, "No transactions found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}
