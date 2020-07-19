package log

import (
	"context"
	"encoding/json"
	log "log"

	"cloud.google.com/go/logging"
)

type logger struct {
}

// JaspergifyLogger is the default agora logger
type JaspergifyLogger interface {
	Debug(json []byte)
	Error(e error)
}

// Logger ...
var Logger *JaspergifyLogger

var cloudLogger *logging.Logger
var localLogger *log.Logger

func init() {
	cloudLogger = createGCloudLogger()
	localLogger = &log.Logger{}
}

func createGCloudLogger() *logging.Logger {
	loggingClient, err := logging.NewClient(context.Background(), "jaspergif")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return loggingClient.Logger("jaspergify-local")
}

// Debug ...
func Debug(entry []byte) {
	cloudLogger.Log(logging.Entry{
		Payload: json.RawMessage(entry),
	})
	indentedBytes, _ := json.MarshalIndent(entry, "", "    ")
	log.Println(string(indentedBytes))
}

// Error ...
func Error(e error) {
	localLogger.Println(e.Error())
	// cloudLogger.Log(logging.Entry{
	// 	e.Error()
	// })
}
