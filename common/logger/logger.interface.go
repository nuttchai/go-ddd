package logger

type ILogger interface {
	Log(messages ...any)
	Logf(message string, options ...any)
	Fatal(messages ...any)
	Fatalf(message string, options ...any)
}
