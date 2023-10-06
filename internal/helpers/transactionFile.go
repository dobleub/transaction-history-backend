package helpers

import (
	"strings"

	"github.com/dobleub/transaction-history-backend/internal/config"
	"github.com/dobleub/transaction-history-backend/pkg/aws"
)

func GetFileFromAWS(env *config.AWSConfig, file string) ([]byte, error) {
	bytesBuffer, err := aws.DownloadObject(env, file)
	if err != nil {
		return nil, err
	}

	bytes := bytesBuffer.Bytes()

	return bytes, nil
}

func DownloadCSVFileFromAWS(env *config.AWSConfig, file string) ([][]string, error) {
	bytes, err := GetFileFromAWS(env, file)
	if err != nil {
		return nil, err
	}

	csvLines := strings.Split(string(bytes), "\n")
	csvLines2 := [][]string{}
	for _, line := range csvLines {
		csvLines2 = append(csvLines2, strings.Split(line, ","))
	}

	return csvLines2, nil
}
