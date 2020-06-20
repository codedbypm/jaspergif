package fetch

import (
	"context"
	"strings"
	"time"

	"github.com/codedbypm/jaspergify/model"
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

	bytes, err = ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "Error: bad response - The response received by Giphy could not read", http.StatusInternalServerError)
		return
	}

	var temnpGif map[string]interface{}
	err = json.Unmarshal(bytes, &temnpGif)
	if err != nil {
		http.Error(w, "Error: bad response - The response received by Giphy could not decoded", http.StatusInternalServerError)
		return
	}

	data := temnpGif["data"].(map[string]interface{})
	id := data["id"].(string)
	images := data["images"].(map[string]interface{})
	original := images["original"].(map[string]interface{})
	frames := original["frames"].(string)
	framesInt, err := strconv.Atoi(frames)
	if err != nil {
		http.Error(w, "Error: bad response - Invalid number of frames from Giphy", http.StatusInternalServerError)
		return
	}
	originalMp4 := images["original_mp4"].(map[string]interface{})
	mp4URL := originalMp4["mp4"].(string)
	size := originalMp4["mp4_size"].(string)
	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		http.Error(w, "Error: bad response - Invalid size of mp4 from Giphy", http.StatusInternalServerError)
		return
	}

	gif := model.GiphyGif{
		Identifier: id,
		Mp4URL:     mp4URL,
		Size:       sizeInt,
		Frames:     framesInt,
	}

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "jaspergif")
	if err != nil {
		http.Error(w, "Error: internal error - could not create Firestore client", http.StatusInternalServerError)
		return
	}

	_, _, err = client.Collection("giphys").Add(ctx, gif)
	if err != nil {
		http.Error(w, "Error: internal error - could not create gif request entry in Firestore", http.StatusInternalServerError)
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

	pubsubTopic := pubsubClient.Topic("new-giphy")

	gifData, err := json.Marshal(gif)
	if err != nil {
		http.Error(w, "Error: internal error - could not marshal GiphyGif instance", http.StatusInternalServerError)
		return
	}

	pubResult := pubsubTopic.Publish(r.Context(), &pubsub.Message{Data: gifData})
	if _, err := pubResult.Get(r.Context()); err != nil {
		http.Error(w, "Error: internal error - could not publish Pub/Sub topic 'new-giphy'", http.StatusInternalServerError)
		return
	}
}
