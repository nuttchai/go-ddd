package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	Info  *log.Logger
	Error *log.Logger
}

func NewLogger(domain ...string) ILogger {
	domainPrefix := "APP"
	if len(domain) > 0 {
		domainPrefix = domain[0]
	}

	infoPrefix := domainPrefix + "_LOG: "
	errorPrefix := domainPrefix + "_ERROR_LOG: "
	return &Logger{
		Info:  log.New(os.Stdout, infoPrefix, log.Ldate|log.Ltime),
		Error: log.New(os.Stdout, errorPrefix, log.Ldate|log.Ltime),
	}
}

func (s *Logger) Log(messages ...any) {
	var logMsg string
	for index, message := range messages {
		if index == 0 {
			logMsg += fmt.Sprintf("%v", message)
			continue
		}
		logMsg += fmt.Sprintf(" %v", message)
	}

	s.Info.Println(logMsg)
}

func (s *Logger) Logf(message string, options ...any) {
	s.Info.Printf(message, options...)
}

func (s *Logger) Fatal(messages ...any) {
	var fatalMsg string
	for index, message := range messages {
		if index == 0 {
			fatalMsg += fmt.Sprintf("%v", message)
			continue
		}
		fatalMsg += fmt.Sprintf(" %v", message)
	}

	s.Error.Fatal(fatalMsg)
}

func (s *Logger) Fatalf(message string, options ...any) {
	s.Error.Fatalf(message, options...)
}
