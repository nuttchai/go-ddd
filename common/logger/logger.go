package logger

import (
	"log"
	"os"
)

type Logger struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	WarnLogger  *log.Logger
}

func (s *Logger) Info(message string, options ...any) {
	s.InfoLogger.Printf(message, options...)
}

func (s *Logger) Error(message string, options ...any) {
	s.ErrorLogger.Fatalf(message, options...)
}

func (s *Logger) Warn(message string, options ...any) {
	s.WarnLogger.Printf(message, options...)
}

func NewLogger(domain ...string) ILogger {
	domainPrefix := "APP"
	if len(domain) > 0 {
		domainPrefix = domain[0]
	}

	infoPrefix := domainPrefix + "_LOG: "
	errorPrefix := domainPrefix + "_ERROR_LOG: "
	warnPrefix := domainPrefix + "_WARN_LOG: "

	return &Logger{
		InfoLogger:  log.New(os.Stdout, infoPrefix, log.Ldate|log.Ltime),
		ErrorLogger: log.New(os.Stdout, errorPrefix, log.Ldate|log.Ltime),
		WarnLogger:  log.New(os.Stdout, warnPrefix, log.Ldate|log.Ltime),
	}
}
