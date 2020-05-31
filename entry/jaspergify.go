// Package entry contains an HTTP Cloud Function.
package entry

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Entry is the new amazing thing
func Entry(w http.ResponseWriter, r *http.Request) {

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
		http.Error(w, "Error: bad request - invalid request", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(bytes, &requestBody); err != nil {
		http.Error(w, "Error: bad request - invalid body", http.StatusBadRequest)
		return
	}

	urlComponents, err := url.Parse("https://media3.giphy.com/media/xUOxf4i3pCY6WwxBOo/giphy.gif?cid=6104955e701de1ad9435afe4c20fa29dd9ef371f6b36f183&rid=giphy.gif")
	if err != nil {
		http.Error(w, "Error: bad request - invalid url", http.StatusBadRequest)
		return
	}

	urlComponents.Query().Get("cid")
}
