package logger

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
)

const separator = "$"

type TestEncoder struct {
	TimeFormat string
}

func (encoder TestEncoder) Encode(message Message) (string, error) {
	return fmt.Sprintf("%s%s%s%s%s", message.Level, separator, message.Time.Format(encoder.TimeFormat), separator, message.Data), nil
}

func TestDefaultLogger_Trace(t *testing.T) {
	testDefaultLoggerMethodByLevel(t, Trace)
}

func TestDefaultLogger_Debug(t *testing.T) {
	testDefaultLoggerMethodByLevel(t, Debug)
}

func TestDefaultLogger_Info(t *testing.T) {
	testDefaultLoggerMethodByLevel(t, Info)
}

func TestDefaultLogger_Warn(t *testing.T) {
	testDefaultLoggerMethodByLevel(t, Warn)
}

func TestDefaultLogger_Error(t *testing.T) {
	testDefaultLoggerMethodByLevel(t, Error)
}

func TestDefaultLogger_Fatal(t *testing.T) {
	testDefaultLoggerMethodByLevel(t, Fatal)
}
func TestDefaultLogger_Log(t *testing.T) {
	testDefaultLoggerMethodByLevel(t, "Custom")
}

func testDefaultLoggerMethodByLevel(t *testing.T, level string) {
	params := []struct {
		data   string
		format string
	}{
		{
			data:   "data",
			format: "2006-01-02 15:04:05",
		},
		{
			data:   "data2",
			format: "2006-01-02",
		},
	}
	for _, param := range params {
		var output bytes.Buffer
		datetime := time.Now().Format(param.format)
		err := callLoggerMethodByLevel(DefaultLogger{Encoder: TestEncoder{TimeFormat: param.format}, Output: &output}, level, param.data)
		if err != nil {
			t.Fatal(err)
		}
		loggedLevel, loggedDatetime, loggedData := loggedMessageParts(output.String())
		if loggedLevel != level {
			t.Errorf("level is %s but expected %s", loggedLevel, level)
		}
		if loggedDatetime != datetime {
			t.Errorf("datetime is %s but expected %s", loggedDatetime, datetime)
		}
		if loggedData != param.data {
			t.Errorf("data is %s but expected %s", loggedData, param.data)
		}
	}
}

func loggedMessageParts(loggedMessage string) (string, string, string) {
	parts := strings.Split(loggedMessage, separator)
	return parts[0], parts[1], parts[2]
}

func callLoggerMethodByLevel(logger Logger, level, data string) error {
	switch level {
	case Trace:
		return logger.Trace(data)
	case Debug:
		return logger.Debug(data)
	case Info:
		return logger.Info(data)
	case Warn:
		return logger.Warn(data)
	case Error:
		return logger.Error(data)
	case Fatal:
		return logger.Fatal(data)
	default:
		return logger.Log(level, data)
	}
}
