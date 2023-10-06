package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/dobleub/transaction-history-backend/internal/errors"
	"github.com/dobleub/transaction-history-backend/internal/helpers"
	"github.com/stretchr/testify/assert"
)

func FindAndReadTestJsonFile(filename string) error {
	var tests []struct {
		Name     string         `json:"name"`
		Desc     string         `json:"desc"`
		Req      *http.Request  `json:"req"`
		Res      *http.Response `json:"res"`
		Expected *http.Response `json:"expected"`
	}

	err := helpers.ReadJsonFile(filename, tests)

	return err
}

func TestReadFile(t *testing.T) {
	t.Run("ReadFile: Found file", func(t *testing.T) {
		_, err := helpers.ReadFile(helpers.GetFilePrefix() + "internal/handlers/tests/handlerVersion.json")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("ReadFile: File not found", func(t *testing.T) {
		_, err := helpers.ReadFile("internal/handlers/tests/handlerVersion2.json")

		if err == nil || !assert.Contains(t, err.Error(), "no such file or directory") {
			t.Error(fmt.Errorf("Expected error to contain 'no such file or directory', got %v", err))
		}
	})
}

func TestReadJsonFile(t *testing.T) {
	t.Run("ReadJsonFile: Found file", func(t *testing.T) {
		err := FindAndReadTestJsonFile("internal/handlers/tests/handlerVersion.json")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("ReadJsonFile: File is not .json", func(t *testing.T) {
		err := FindAndReadTestJsonFile("data/transactions.csv")

		if err == nil || !assert.Equal(t, err.Error(), errors.ErrFileNotJSON) {
			t.Error(fmt.Errorf("Expected error %v', got %v", errors.ErrFileNotJSON, err))
		}
	})
}

func TestReadCSVFile(t *testing.T) {
	t.Run("ReadCSVFile: Found file", func(t *testing.T) {
		_, err := helpers.ReadCSVFile("data/transactions.csv")
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("ReadCSVFile: File is not .csv", func(t *testing.T) {
		_, err := helpers.ReadCSVFile("internal/handlers/tests/handlerVersion.json")

		if err == nil || !assert.Equal(t, err.Error(), errors.ErrFileNotCSV) {
			t.Error(fmt.Errorf("Expected error %v', got %v", errors.ErrFileNotCSV, err))
		}
	})
}
