// Package decode contains an HTTP Cloud Function.
package decode

import (
	"context"
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
		Mp4URL string `json:"mp4URL"`
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error: bad request - invalid request", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(bytes, &gifInfo); err != nil {
		http.Error(w, "Error: bad request - invalid body", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	downloadClient := http.Client()
	response, err := downloadClient.Get(gifMp4.Mp4URL)
	defer response.Body.Close()

	gif, err := gif.DecodeAll(response.Body)
}
