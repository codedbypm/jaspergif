package main

import (
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/codedbypm/jaspergif/entry"
	"github.com/codedbypm/jaspergif/fetch"
	"github.com/codedbypm/jaspergif/upload"
)

func main() {
	funcframework.RegisterHTTPFunction("/entry", entry.Entry)
	funcframework.RegisterEventFunction("/fetch", fetch.OnCreateRequest)
	funcframework.RegisterEventFunction("/upload", upload.OnFetchGif)

	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	if err := funcframework.Start(port); err != nil {
		log.Println(err)
	}
}
