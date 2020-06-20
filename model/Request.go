package model

import "time"

// Request models a request to Jaspergify a gif
type Request struct {
	GiphyIdentifier string        `firestore:"giphyId"`
	Timestamp       time.Time     `firestore:"time"`
	Status          RequestStatus `firestore:"status"`
}

// RequestStatus is the status of a Jaspergify request
type RequestStatus int

const (
	// Received means the request has been received
	Received RequestStatus = iota
	// Ongoing means the request started but not yet finished
	Ongoing
	// Done means the request has been executed
	Done
)

func (s RequestStatus) String() string {
	return [...]string{"Received", "Ongoing", "Done"}[s]
}
