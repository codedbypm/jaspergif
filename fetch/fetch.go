package fetch

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/codedbypm/jaspergify/fetch/model"
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
	originalMp4 := images["original_mp4"].(map[string]interface{})
	mp4URL := originalMp4["mp4"].(string)
	size := originalMp4["mp4_size"].(string)
	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		http.Error(w, "Error: bad response - Invalid size of mp4 filpe from Giphy", http.StatusInternalServerError)
		return
	}

	gif := model.GiphyGif{
		Identifier: id,
		Mp4URL:     mp4URL,
		Size:       sizeInt,
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

}
