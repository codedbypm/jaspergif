package model

// GiphyGif model the respone received from Giphy
type GiphyGif struct {
	Identifier string `firestore:"identifier"`
	Mp4URL     string `firestore:"mp4URL"`
	Size       int    `firestore:"size"`
}
