package model

import "time"

// FirestoreGifEvent holds Firestore fields.
type FirestoreGifEvent struct {
	Value FirestoreGifValue `json:"value"`
}

// FirestoreGifValue holds Firestore fields.
type FirestoreGifValue struct {
	CreateTime time.Time    `json:"createTime"`
	Fields     FirestoreGif `json:"fields"`
	Name       string       `json:"name"`
	UpdateTime time.Time    `json:"updateTime"`
}

// FirestoreGif is awesome
type FirestoreGif struct {
	URL        StringValue `json:"mp4URL"`
	Identifier StringValue `json:"identifier"`
}
