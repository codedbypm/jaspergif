package model

import "time"

// JaspergifyRequest models a request to Jaspergify a gif
type JaspergifyRequest struct {
	GiphyIdentifier string                  `firestore:"giphyIdentifier"`
	Timestamp       time.Time               `firestore:"time"`
	Status          JaspergifyRequestStatus `firestore:"status"`
}

// JaspergifyRequestStatus is the status of a Jaspergify request
type JaspergifyRequestStatus int

const (
	// Received means the request has been received
	Received JaspergifyRequestStatus = iota
	// Ongoing means the request started but not yet finished
	Ongoing
	// Done means the request has been executed
	Done
)

func (s JaspergifyRequestStatus) String() string {
	return [...]string{"Received", "Ongoing", "Done"}[s]
}
