package logger

type Message struct {
	Level string
	Time  string
	Data  any
}

type Encoder interface {
	Encode(message Message) (string, error)
}
