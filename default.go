package logger

import (
	"fmt"
	"io"
	"time"
)

type DefaultLogger struct {
	Encoder    Encoder
	Output     io.Writer
	TimeFormat string
}

func (logger DefaultLogger) Trace(data any) error {
	return logger.Log(Trace, data)
}

func (logger DefaultLogger) Debug(data any) error {
	return logger.Log(Debug, data)
}

func (logger DefaultLogger) Info(data any) error {
	return logger.Log(Info, data)
}

func (logger DefaultLogger) Warn(data any) error {
	return logger.Log(Warn, data)
}

func (logger DefaultLogger) Error(data any) error {
	return logger.Log(Error, data)
}

func (logger DefaultLogger) Fatal(data any) error {
	return logger.Log(Fatal, data)
}

func (logger DefaultLogger) Log(level string, data any) error {
	encodedMessage, err := logger.Encoder.Encode(Message{Level: level, Time: time.Now().Format(logger.TimeFormat), Data: data})
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(logger.Output, encodedMessage)
	if err != nil {
		return err
	}
	return nil
}
