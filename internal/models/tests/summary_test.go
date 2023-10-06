package tests

import (
	"testing"

	"github.com/dobleub/transaction-history-backend/internal/models"
)

func TestGetGetSummaryEmailData(t *testing.T) {
	t.Run("GetGetSummaryEmailData", func(t *testing.T) {
		summary := models.Summary{
			TotalBalance:        1603.25,
			TotalIncome:         3664.35,
			TotalExpense:        -2061.1,
			AverageCreditAmount: 174.49285714285713,
			AverageDebitAmount:  -98.14761904761905,
			TransactionsPerMonth: map[string]models.TransactionsPerMonth{
				"april": {
					Amount: 3.1900000000000084,
					Total:  9,
				},
				"february": {
					Amount: 782.7,
					Total:  8,
				},
				"january": {
					Amount: -82.2,
					Total:  11,
				},
				"march": {
					Amount: 779.35,
					Total:  13,
				},
				"may": {
					Amount: 120.21,
					Total:  1,
				},
			},
		}

		summaryEmailBody := summary.GetSummaryEmailData()

		if summaryEmailBody.Subject != "Summary" {
			t.Errorf("Expected 'Summary', got %v", summaryEmailBody.Subject)
		}
		if len(summaryEmailBody.Summary) != 5 {
			t.Errorf("Expected 5, got %v", len(summaryEmailBody.Summary))
		}
		if len(summaryEmailBody.Transactions) != 5 {
			t.Errorf("Expected 5, got %v", len(summaryEmailBody.Transactions))
		}
	})
}
