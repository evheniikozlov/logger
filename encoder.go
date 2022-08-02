package logger

import "time"

type Message struct {
	Level string
	Time  time.Time
	Data  any
}

type Encoder interface {
	Encode(message Message) (string, error)
}
