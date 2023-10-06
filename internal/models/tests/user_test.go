package tests

import (
	"context"
	"testing"

	"github.com/dobleub/transaction-history-backend/internal/config"
	"github.com/dobleub/transaction-history-backend/internal/models"
	"github.com/sethvargo/go-envconfig"
)

var env *config.Config

func TestMain(m *testing.M) {
	var tmpEnv config.Config
	envconfig.ProcessWith(context.Background(), &tmpEnv, envconfig.OsLookuper())

	env = &tmpEnv
	m.Run()
}

func TestGetUserId(t *testing.T) {
	t.Run("GetUserId", func(t *testing.T) {
		user := models.User{
			UserId: 1,
		}

		userId := user.GetUserId()

		if userId != 1 {
			t.Errorf("Expected 1, got %v", userId)
		}
	})

	t.Run("GetUserId: Not found", func(t *testing.T) {
		user := models.User{}

		userId := user.GetUserId()

		if userId != 0 {
			t.Errorf("Expected 0, got %v", userId)
		}
	})
}

func TestGetTransactions(t *testing.T) {
	t.Run("GetTransactions", func(t *testing.T) {
		user := models.User{
			UserId: 1,
		}

		transactions, err := user.GetTransactions(env)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if len(transactions) == 0 {
			t.Errorf("Expected more than 0, got %v", len(transactions))
		}
	})

	t.Run("GetTransactions: Not found", func(t *testing.T) {
		user := models.User{
			UserId: 999,
		}

		transactions, err := user.GetTransactions(env)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if len(transactions) != 0 {
			t.Errorf("Expected 0, got %v", len(transactions))
		}
	})
}

func TestGetTransactionsPerMonth(t *testing.T) {
	t.Run("GetTransactionsPerMonth", func(t *testing.T) {
		user := models.User{
			UserId: 1,
		}

		transactionsPerMonth, err := user.GetTransactionsPerMonth(env)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if len(transactionsPerMonth) == 0 {
			t.Errorf("Expected more than 0, got %v", len(transactionsPerMonth))
		}
	})

	t.Run("GetTransactionsPerMonth: Not found", func(t *testing.T) {
		user := models.User{
			UserId: 999,
		}

		transactionsPerMonth, err := user.GetTransactionsPerMonth(env)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if len(transactionsPerMonth) != 0 {
			t.Errorf("Expected 0, got %v", len(transactionsPerMonth))
		}
	})
}

func TestGetSummary(t *testing.T) {
	t.Run("GetSummary", func(t *testing.T) {
		user := models.User{
			UserId: 1,
		}

		summary, err := user.GetSummary(env)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if len(summary.TransactionsPerMonth) == 0 {
			t.Errorf("Expected more than 0, got %v", len(summary.TransactionsPerMonth))
		}
	})

	t.Run("GetSummary: Not found", func(t *testing.T) {
		user := models.User{
			UserId: 999,
		}

		summary, err := user.GetSummary(env)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if len(summary.TransactionsPerMonth) != 0 {
			t.Errorf("Expected 0, got %v", len(summary.TransactionsPerMonth))
		}
	})
}
