package logger

import (
	"log"
	"os"
)

type key int
const LoggerKey key = 0

// Logger is a middleware handler that logs the request as it goes in and the response as it goes out.
type Logger struct {
	// Logger is the log.Logger instance used to log messages with the Logger middleware
	Logger *log.Logger
	isInit bool
}

// NewStaticLogger returns a logger instance to be used for the server
var l Logger

func GetStaticLogger() *Logger {
	if !l.isInit {
		l = Logger{
			Logger: log.New(os.Stdout, "[snowball.im - ] ", 0),
			isInit: true,
		}
	}
	return &l
}
