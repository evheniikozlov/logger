package logger

type Logger interface {
	Trace(data any) error
	Debug(data any) error
	Info(data any) error
	Warn(data any) error
	Error(data any) error
	Fatal(data any) error
	Log(level string, data any) error
}

const (
	Trace = "TRACE"
	Debug = "DEBUG"
	Info  = "INFO"
	Warn  = "WARN"
	Error = "ERROR"
	Fatal = "FATAL"
)
