// Package entry contains an HTTP Cloud Function.
package entry

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"cloud.google.com/go/firestore"
)

type Gif struct {
	Identifier string
}

// Entry is the new amazing thing
func Entry(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		http.Error(w, "Error: not found", http.StatusNotFound)
		return
	}

	var gifInfo struct {
		URL string `json:"url"`
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

	gifURLComponents, err := url.Parse(gifInfo.URL)
	if err != nil {
		http.Error(w, "Error: bad request - invalid URL", http.StatusBadRequest)
		return
	}

	query, err := url.ParseQuery(gifURLComponents.RawQuery)
	if err != nil {
		http.Error(w, "Error: bad request - missing required 'cid' query item", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "jaspergif")
	if err != nil {
		http.Error(w, "Error: internal error - could not create Firestore client", http.StatusInternalServerError)
		return
	}

	write, err := client.Collection("gifs").NewDoc().Create(ctx, Gif{
		Identifier: query.Get("cid"),
	})
	if err != nil {
		http.Error(w, "Error: internal error - could not create GIF entry in Firestore", http.StatusInternalServerError)
		return
	}

	fmt.Println(write)
}
