package fetch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Fetch is the next new thing
func Fetch(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
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

	// Create Giphy request
	res, err := http.Get("https://api.giphy.com/v1/gifs/" + requestBody.Identifier + "?api_key=QuCgTOvpRJlHx6QMtNCYTqfL5Efj0vgT")
	if err != nil {
		http.Error(w, "Error: bad request - invalid request for api.giphy.com", http.StatusBadRequest)
		return
	}

	if res.StatusCode != http.StatusOK {
		http.Error(w, "Error: not found - The request gif could not be found on api.giphy.com", http.StatusNotFound)
		return
	}
}
