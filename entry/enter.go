// Package entry contains an HTTP Cloud Function.
package entry

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
)

type gif struct {
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

	gifURL, err := url.Parse(gifInfo.URL)
	pathComponents := strings.Split(gifURL.Path, "/")
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

	gif, err := client.Collection("gifs").NewDoc().Create(ctx, gif{
		Identifier: pathComponents[2],
	})
	if err != nil {
		http.Error(w, "Error: internal error - could not create GIF entry in Firestore", http.StatusInternalServerError)
		return
	}

	gifData, err := json.Marshal(gif)
	if err != nil {
		http.Error(w, "Error: internal error - could not marshal new gif entry", http.StatusInternalServerError)
		return
	}

	// PubSubMessage is the payload of a Pub/Sub event.
	// See https://cloud.google.com/functions/docs/calling/pubsub.
	type PubSubMessage struct {
		Data []byte `json:"data"`
	}

	// Create Pub/Sub client
	pubsubClient, err := pubsub.NewClient(ctx, "jaspergif")
	if err != nil {
		http.Error(w, "Error: internal error - could not create Pub/Sub Client", http.StatusInternalServerError)
		return
	}

	pubsubTopic := pubsubClient.Topic("new-gif-request")

	res := pubsubTopic.Publish(r.Context(), &pubsub.Message{Data: gifData})
	if _, err := res.Get(r.Context()); err != nil {
		http.Error(w, "Error: internal error - could not publish Pub/Sub topic", http.StatusInternalServerError)
		return
	}
}
