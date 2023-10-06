package models

import (
	"strings"

	"github.com/dobleub/transaction-history-backend/internal/helpers"
)

type User struct {
	UserId int32 `json:"userId"`
}

func (u *User) GetUserId() int32 {
	return u.UserId
}

func (u *User) GetTransactions() ([]Transaction, error) {
	var transactions []Transaction

	csvFile, err := helpers.ReadCSVFile("data/transactions.csv")
	if err != nil {
		return transactions, err
	}

	for _, line := range csvFile {
		tmpUserId := helpers.StringToInt32(line[1])
		if tmpUserId != u.GetUserId() {
			continue
		}
		transaction := Transaction{
			TransactionId: helpers.StringToInt32(line[0]),
			UserId:        tmpUserId,
			Date:          line[2],
			Transaction:   helpers.StringToFloat64(line[3]),
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (u *User) GetTransactionsPerMonth() (map[string]TransactionsPerMonth, error) {
	transactions, err := u.GetTransactions()
	if err != nil {
		return nil, err
	}
	if len(transactions) == 0 {
		return nil, nil
	}

	transactionsPerMonth := make(map[string]TransactionsPerMonth)
	for _, transaction := range transactions {
		month := strings.ToLower(helpers.StringToDate(transaction.Date).Month().String())

		if _, ok := transactionsPerMonth[month]; !ok {
			transactionsPerMonth[month] = TransactionsPerMonth{
				Amount: transaction.Transaction,
				Total:  1,
			}
		} else {
			transactionsPerMonth[month] = TransactionsPerMonth{
				Amount: transactionsPerMonth[month].Amount + transaction.Transaction,
				Total:  transactionsPerMonth[month].Total + 1,
			}
		}
	}

	return transactionsPerMonth, nil
}

func (u *User) GetSummary() (Summary, error) {
	var summary Summary
	var incomeTransactions, expenseTransactions int

	transactions, err := u.GetTransactions()
	if err != nil {
		return summary, err
	}

	for _, transaction := range transactions {
		// Total balance
		summary.TotalBalance += transaction.Transaction

		// Total income
		if transaction.Transaction > 0 {
			incomeTransactions++
			summary.TotalIncome += transaction.Transaction
		}

		// Total expense
		if transaction.Transaction < 0 {
			expenseTransactions++
			summary.TotalExpense += transaction.Transaction
		}
	}

	// Average credit amount
	summary.AverageCreditAmount = summary.TotalIncome / float64(incomeTransactions)
	// Average debit amount
	summary.AverageDebitAmount = summary.TotalExpense / float64(expenseTransactions)

	// Transactions per month
	transactionsPerMonth, err := u.GetTransactionsPerMonth()
	if err != nil {
		return summary, err
	}

	summary.TransactionsPerMonth = transactionsPerMonth

	return summary, nil
}
