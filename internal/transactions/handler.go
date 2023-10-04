package transactions

import (
	"encoding/json"
	"net/http"
)

func HandleVersion(w http.ResponseWriter, r *http.Request) {
	version := struct {
		Version string `json:"version"`
	}{}
	version.Version = "1.0.1 - 2021-09-30 - transactions-history"
	json.NewEncoder(w).Encode(version)
}
