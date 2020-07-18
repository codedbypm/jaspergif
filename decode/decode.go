// Package decode contains an HTTP Cloud Function.
package decode

import (
	"encoding/json"
	"image/gif"
	"io/ioutil"
	"net/http"
)

// Decode is the new amazing thing
func Decode(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Error: not found", http.StatusNotFound)
		return
	}

	var gifMp4 struct {
		url  string `json:"mp4URL"`
		size int    `json:"size"`
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error: bad request - invalid request", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(bytes, &gifMp4); err != nil {
		http.Error(w, "Error: bad request - invalid body", http.StatusBadRequest)
		return
	}

	response, err := http.Get(gifMp4.url)
	defer response.Body.Close()

	gif, err := gif.DecodeAll(response.Body)
	if err != nil {
		http.Error(w, "Error: bad response - could not download mp4 file", http.StatusBadRequest)
		return
	}

	print(gif)
}
