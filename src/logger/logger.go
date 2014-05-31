package logger

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/context"
)

type key int
const LoggerKey key = 0

// Logger is a middleware handler that logs the request as it goes in and the response as it goes out.
type Logger struct {
	// Logger is the log.Logger instance used to log messages with the Logger middleware
	Logger *log.Logger
}

// NewLogger returns a new Logger instance
func NewLogger() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "[snowball.im - ] ", 0),
	}
}

func GetLogger(r *http.Request) *Logger {
	rv := context.Get(r, LoggerKey)
	if rv == nil {
		rv = NewLogger()
		context.Set(r, LoggerKey, rv)
	}
	return rv.(*Logger)
}

func (l *Logger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, LoggerKey, l)
	next(rw, r)
}


