package models

import (
	"fmt"

	"github.com/dobleub/transaction-history-backend/internal/helpers"
)

type Summary struct {
	TotalBalance         float64                         `json:"totalBalance"`
	TotalIncome          float64                         `json:"totalIncome"`
	TotalExpense         float64                         `json:"totalExpense"`
	TransactionsPerMonth map[string]TransactionsPerMonth `json:"transactionsPerMonth"`
	AverageCreditAmount  float64                         `json:"averageCreditAmount"`
	AverageDebitAmount   float64                         `json:"averageDebitAmount"`
}

type TransactionsPerMonth struct {
	Amount float64 `json:"amount"`
	Total  int32   `json:"total"`
}

type GeneralDescrition struct {
	Desc  string `json:"desc"`
	Value string `json:"value"`
}

type SummaryEmailBody struct {
	Subject      string              `json:"subject"`
	Summary      []GeneralDescrition `json:"summary"`
	Transactions []GeneralDescrition `json:"transactions"`
}

func (s *Summary) GetSummaryEmailData() SummaryEmailBody {
	summaryEmailBody := SummaryEmailBody{
		Subject: "Summary",
		Summary: []GeneralDescrition{
			{
				Desc:  "Total Balance",
				Value: fmt.Sprintf("%v", helpers.FormatMoney(s.TotalBalance, 2)),
			},
			{
				Desc:  "Total Credit",
				Value: fmt.Sprintf("%v", helpers.FormatMoney(s.TotalIncome, 2)),
			},
			{
				Desc:  "Total Debit",
				Value: fmt.Sprintf("%v", helpers.FormatMoney(s.TotalExpense, 2)),
			},
			{
				Desc:  "Average Credit Amount",
				Value: fmt.Sprintf("%v", helpers.FormatMoney(s.AverageCreditAmount, 2)),
			},
			{
				Desc:  "Average Debit Amount",
				Value: fmt.Sprintf("%v", helpers.FormatMoney(s.AverageDebitAmount, 2)),
			},
		},
		Transactions: []GeneralDescrition{},
	}

	for month, transactionsPerMonth := range s.TransactionsPerMonth {
		summaryEmailBody.Transactions = append(summaryEmailBody.Transactions, GeneralDescrition{
			Desc:  fmt.Sprintf("Transactions in %v", helpers.Capitalize(month)),
			Value: fmt.Sprintf("%v", transactionsPerMonth.Total),
		})
	}

	return summaryEmailBody
}
