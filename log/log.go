package log

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/logging"
)

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
	localLogger = log.New(os.Stdout, "[Local]: ", 0)
	cloudLogger = createGCloudLogger()
}

func createGCloudLogger() *logging.Logger {
	loggingClient, err := logging.NewClient(context.Background(), "jaspergif")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	logName := "jaspergify-log"
	return loggingClient.Logger(logName)
}

// Debug is
func Debug(entry string) {
	localLogger.Println(entry)
	cloudLogger.StandardLogger(logging.Debug).Println(entry)
}

// Error ...
func Error(e error) {
	log.Println(e)
	localLogger.Println(e)
	cloudLogger.StandardLogger(logging.Error).Println(e)
}
