package fetch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Fetch(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Error(w, "Error: not found", http.StatusNotFound)
		return
	}

	var requestBody struct {
		Identifier string `json:"identifier"`
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error: bad request - invalid request", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(bytes, &requestBody); err != nil {
		http.Error(w, "Error: bad request - invalid body", http.StatusBadRequest)
		return
	}
}
