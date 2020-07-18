package main

import (
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/codedbypm/jaspergify/decode"
	"github.com/codedbypm/jaspergify/entry"
	"github.com/codedbypm/jaspergify/fetch"
)

func main() {
	funcframework.RegisterHTTPFunction("/enter", entry.Entry)
	funcframework.RegisterHTTPFunction("/fetch", fetch.Fetch)
	funcframework.RegisterHTTPFunction("/decode", decode.Decode)

	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
