package fetch

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/functions/metadata"
)

// FirestoreEvent is the payload of a Firestore event.
type FirestoreEvent struct {
	OldValue   FirestoreValue `json:"oldValue"`
	Value      FirestoreValue `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// FirestoreValue holds Firestore fields.
type FirestoreValue struct {
	CreateTime time.Time   `json:"createTime"`
	Fields     interface{} `json:"fields"`
	Name       string      `json:"name"`
	UpdateTime time.Time   `json:"updateTime"`
}

// OnFirestoreWrite is the new awesome thing
func OnFirestoreWrite(ctx context.Context, e FirestoreEvent) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	log.Printf("Function triggered by change to: %v", meta.Resource)
	log.Printf("Old value: %+v", e.OldValue)
	log.Printf("New value: %+v", e.Value)
	return nil

	// collection := pathParts[0]
	// doc := strings.Join(pathParts[1:], "/")
	// curValue := e.Value.Fields.Original.StringValue

	// giphyIdentifier := payload["GiphyIdentifier"].(string)

	// // Create Giphy request
	// res, err := http.Get("https://api.giphy.com/v1/gifs/" + giphyIdentifier + "?api_key=QuCgTOvpRJlHx6QMtNCYTqfL5Efj0vgT")
	// if err != nil {
	// 	return err
	// }

	// if res.StatusCode != http.StatusOK {
	// 	return errors.New("Error: not found - The request gif could not be found on api.giphy.com")
	// }

	// bytes, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	return err
	// }

	// var temnpGif map[string]interface{}
	// err = json.Unmarshal(bytes, &temnpGif)
	// if err != nil {
	// 	return err
	// }

	// data := temnpGif["data"].(map[string]interface{})
	// id := data["id"].(string)
	// images := data["images"].(map[string]interface{})
	// original := images["original"].(map[string]interface{})
	// frames := original["frames"].(string)
	// framesInt, err := strconv.Atoi(frames)
	// if err != nil {
	// 	return err
	// }
	// originalMp4 := images["original_mp4"].(map[string]interface{})
	// mp4URL := originalMp4["mp4"].(string)
	// size := originalMp4["mp4_size"].(string)
	// sizeInt, err := strconv.Atoi(size)
	// if err != nil {
	// 	return err
	// }

	// gif := model.GiphyGif{
	// 	Identifier: id,
	// 	Mp4URL:     mp4URL,
	// 	Size:       sizeInt,
	// 	Frames:     framesInt,
	// }

	// newCtx := context.Background()
	// client, err := firestore.NewClient(newCtx, "jaspergif")
	// if err != nil {
	// 	return err
	// }

	// _, _, err = client.Collection("giphys").Add(newCtx, gif)
	// if err != nil {
	// 	return err
	// }

	// // PubSubMessage is the payload of a Pub/Sub event.
	// // See https://cloud.google.com/functions/docs/calling/pubsub.
	// type PubSubMessage struct {
	// 	Data []byte `json:"data"`
	// }

	// // Create Pub/Sub client
	// pubsubClient, err := pubsub.NewClient(newCtx, "jaspergif")
	// if err != nil {
	// 	return err
	// }

	// pubsubTopic := pubsubClient.Topic("new-giphy")

	// gifData, err := json.Marshal(gif)
	// if err != nil {
	// 	return err
	// }

	// pubResult := pubsubTopic.Publish(r.Context(), &pubsub.Message{Data: gifData})
	// if _, err := pubResult.Get(r.Context()); err != nil {
	// 	return err
	// }
}
