package helpers

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ReadJsonFile(filePath string, obj interface{}) error {
	filePrefix := getFilePrefix()

	jsonFile, err := os.Open(filePrefix + filePath)
	if err != nil {
		return err
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer func(jsonFile *os.File) {
		_ = jsonFile.Close()
	}(jsonFile)

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(byteValue, &obj); err != nil {
		return err
	}

	return nil
}

func getFilePrefix() string {
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
