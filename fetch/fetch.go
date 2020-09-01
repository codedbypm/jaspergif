package fetch

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"cloud.google.com/go/firestore"
	goLog "github.com/codedbypm/golog"
	"github.com/codedbypm/jaspergif/model"
)

// OnCreateRequest is the new awesome thing
func OnCreateRequest(ctx context.Context, e model.FirestoreRequestEvent) error {

	var giphyIdentifier = e.Value.Fields.GiphyIdentifier.Value

	var url = "https://api.giphy.com/v1/gifs/" + giphyIdentifier + "?api_key=QuCgTOvpRJlHx6QMtNCYTqfL5Efj0vgT"

	// Create Giphy request
	res, err := http.Get(url)
	if err != nil {
		goLog.Error(err)
		return err
	}

	if res.StatusCode != http.StatusOK {
		err = errors.New("Error: not found - The request gif could not be found on api.giphy.com. (" + url + ")")
		goLog.Error(err)
		return err
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		goLog.Error(err)
		return err
	}

	var temnpGif map[string]interface{}
	err = json.Unmarshal(bytes, &temnpGif)
	if err != nil {
		goLog.Error(err)
		return err
	}

	data := temnpGif["data"].(map[string]interface{})
	id := data["id"].(string)
	images := data["images"].(map[string]interface{})
	original := images["original"].(map[string]interface{})
	frames := original["frames"].(string)
	framesInt, err := strconv.Atoi(frames)
	if err != nil {
		goLog.Error(err)
		return err
	}
	mp4URL := original["mp4"].(string)
	size := original["mp4_size"].(string)
	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		goLog.Error(err)
		return err
	}

	gif := model.GiphyGif{
		Identifier: id,
		Mp4URL:     mp4URL,
		Size:       sizeInt,
		Frames:     framesInt,
	}

	newCtx := context.Background()
	client, err := firestore.NewClient(newCtx, "jaspergif")
	if err != nil {
		goLog.Error(err)
		return err
	}

	_, _, err = client.Collection("giphys").Add(newCtx, gif)
	if err != nil {
		goLog.Error(err)
		return err
	}

	return nil
}
