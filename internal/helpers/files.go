package helpers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dobleub/transaction-history-backend/internal/errors"
)

func ReadFile(filePath string) ([]byte, error) {
	byteValue, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return byteValue, nil
}

func ReadJsonFile(filePath string, obj interface{}) error {
	if !strings.Contains(filePath, ".json") {
		return fmt.Errorf(errors.ErrFileNotJSON)
	}

	filePrefix := GetFilePrefix()
	byteValue, err := ReadFile(filePrefix + filePath)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(byteValue, &obj); err != nil {
		return err
	}

	return nil
}

func ReadCSVFile(filePath string) ([][]string, error) {
	if !strings.Contains(filePath, ".csv") {
		return nil, fmt.Errorf(errors.ErrFileNotCSV)
	}

	filePrefix := GetFilePrefix()
	byteValue, err := ReadFile(filePrefix + filePath)
	if err != nil {
		return nil, err
	}

	// csvLines := [][]string{}
	// csvLines = append(csvLines, strings.Split(string(byteValue), "\n"))
	csvLines := strings.Split(string(byteValue), "\n")
	csvLines2 := [][]string{}
	for _, line := range csvLines {
		csvLines2 = append(csvLines2, strings.Split(line, ","))
	}

	return csvLines2, nil
}

func GetFilePrefix() string {
	wd, _ := os.Getwd()
	base := filepath.Base(wd)
	currDir := wd
	i := 0
	for base != "transactions-history" && base != "src" && base != "/" {
		// fmt.Println(currDir, base)
		currDir = filepath.Dir(currDir)
		base = filepath.Base(currDir)
		i++
	}
	if i > 0 {
		filePrefix := strings.Repeat("../", i)
		if base == "/" {
			filePrefix += "src/"
		}
		return filePrefix
	}
	return "./"
}
