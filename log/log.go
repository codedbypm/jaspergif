package log

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/logging"
)

// JaspergifLogger is the default agora logger
type JaspergifLogger interface {
	Debug(json []byte)
	Error(e error)
}

// Logger ...
var Logger *JaspergifLogger

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

	logName := "jaspergif-log"
	return loggingClient.Logger(logName)
}

// Debug is
func Debug(v ...interface{}) {
	localLogger.Println(v)
	cloudLogger.StandardLogger(logging.Debug).Println(v)
}

// Error ...
func Error(e error) {
	localLogger.Println(e)
	cloudLogger.StandardLogger(logging.Error).Println(e)
}
