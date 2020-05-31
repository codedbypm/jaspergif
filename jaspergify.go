// Package jaspergify contains an HTTP Cloud Function.
package jaspergify

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Run is the new amazing thing
func Run(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		http.Error(w, "Error: not found", http.StatusNotFound)
		return
	}

	var requestBody struct {
		Identifier string `json:"identifier"`
		URL        string `json:"url"`
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		http.Error(w, "Could not read POST request body", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(bytes, &requestBody); err != nil {
		log.Print(err)
		http.Error(w, "Could not decode POST request body", http.StatusBadRequest)
		return
	}

}
