// Package upload contains an HTTP Cloud Function.
package upload

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
	"github.com/codedbypm/jaspergify/log"
	"github.com/codedbypm/jaspergify/model"
)

// OnFetchGif is the new awesome thing
func OnFetchGif(ctx context.Context, e model.FirestoreGifEvent) error {

	var gifURL = e.Value.Fields.URL.Value
	var gifID = e.Value.Fields.Identifier.Value

	// Get the data
	resp, err := http.Get(gifURL)
	if err != nil {
		log.Error(err)
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return err
	}

	return uploadFile(data, "jaspergify-gifs-mp4", gifID)
}

// uploadFile uploads an object.
func uploadFile(data []byte, bucketName, objectName string) error {
	// object := "object-name"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	writer := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)
	if _, err := writer.Write(data); err != nil {
		log.Error(err)
		return err
	}
	if err := writer.Close(); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
