package handlers

import (
	"net/http"

	"github.com/dobleub/transaction-history-backend/internal/config"
	"github.com/dobleub/transaction-history-backend/internal/handlers/email_templates"
	"github.com/dobleub/transaction-history-backend/internal/helpers"
	"github.com/dobleub/transaction-history-backend/internal/models"
	"github.com/dobleub/transaction-history-backend/pkg/email"
	"github.com/gorilla/mux"
)

/*
 * HandleSendEmail
 * Send email with summary of transactions for a user
 * @param w http.ResponseWriter
 * @param r *http.Request
 * @return void
 *
 */
func HandlerSendEmail(env *config.Config, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userid"]
	emailTo := vars["emailto"]

	user := models.User{
		UserId: helpers.StringToInt32(userId),
	}

	summary, err := user.GetSummary(env)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(summary.TransactionsPerMonth) == 0 {
		http.Error(w, "No transactions found", http.StatusNotFound)
		return
	}

	// Generate email body
	summaryEmailBody := summary.GetSummaryEmailData()
	emailBody := email_templates.GenerateSummaryEmailBody(summaryEmailBody)
	// Send email
	mailErr := email.SendEmail(&env.EmailConfig, "Stori Test<no-reply-info@godin.app>", emailTo, summaryEmailBody.Subject, emailBody, "")
	if mailErr != nil {
		http.Error(w, mailErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Email sent"))
}
