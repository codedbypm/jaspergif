// Package entry contains an HTTP Cloud Function.
package entry

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/codedbypm/jaspergify/log"
	"github.com/codedbypm/jaspergify/model"
)

// Entry is the new amazing thing
func Entry(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Error: not found", http.StatusNotFound)
		return
	}

	var gifInfo struct {
		URL string `json:"url"`
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		http.Error(w, "Error: bad request - invalid request", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(bytes, &gifInfo); err != nil {
		log.Error(err)
		http.Error(w, "Error: bad request - invalid body", http.StatusBadRequest)
		return
	}

	gifURL, err := url.Parse(gifInfo.URL)
	pathComponents := strings.Split(gifURL.Path, "/")
	if err != nil {
		log.Error(err)
		http.Error(w, "Error: bad request - missing required 'cid' query item", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "jaspergif")
	if err != nil {
		log.Error(err)
		http.Error(w, "Error: internal error - could not create Firestore client", http.StatusInternalServerError)
		return
	}

	request := model.Request{
		GiphyIdentifier: pathComponents[2],
		Timestamp:       time.Now(),
		Status:          model.Received,
	}

	_, _, err = client.Collection("requests").Add(ctx, request)
	if err != nil {
		log.Error(err)
		http.Error(w, "Error: internal error - could not create gif request entry in Firestore", http.StatusInternalServerError)
		return
	}
}
