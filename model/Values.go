package model

import "time"

// StringValue is awesome
type StringValue struct {
	Value string `json:"stringValue"`
}

// TimestampValue is awesome
type TimestampValue struct {
	Value time.Time `json:"timestampValue"`
}

// IntegerValue is awesome
type IntegerValue struct {
	Value int `json:"integerValue"`
}

// RequestStatusValue is awesome
type RequestStatusValue struct {
	Value RequestStatus `json:"stringValue"`
}
