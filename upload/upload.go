// Package upload contains an HTTP Cloud Function.
package upload

import (
	"context"

	"github.com/codedbypm/jaspergify/model"
)

// OnFetchGif is the new awesome thing
func OnFetchGif(ctx context.Context, e model.FirestoreGifEvent) error {

	// var gifURL = e.Value.Fields.URL.Value

	return nil
}
