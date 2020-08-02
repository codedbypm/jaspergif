// Package decode contains an HTTP Cloud Function.
package decode

import (
	"context"

	"github.com/codedbypm/jaspergify/log"
	"github.com/codedbypm/jaspergify/model"
	"gocv.io/x/gocv"
)

// OnFetchGif is the new awesome thing
func OnFetchGif(ctx context.Context, e model.FirestoreGifEvent) error {

	var gifURL = e.Value.Fields.URL.Value

	mp4, err := gocv.VideoCaptureFile(gifURL)
	defer mp4.Close()
	if !err {
		log.Error(err)
		return err
	}

	// create an empty n-dimensional matrix.
	// These matrix will be used to store the images we read from our camera later.
	img := gocv.NewMat()
	defer img.Close()

	for {
		if ok := mp4.Read(&img); !ok || img.Empty() {
			log.Debug("Unable to read from the webcam")
			continue
		}

	}
	return nil
}
