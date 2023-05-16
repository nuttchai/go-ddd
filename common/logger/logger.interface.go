package logger

type ILogger interface {
	Info(message string, options ...any)
	Warn(message string, options ...any)
	Error(message string, options ...any)
}
