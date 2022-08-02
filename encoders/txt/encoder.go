package txt

import (
	"fmt"
	"github.com/evheniikozlov/logger"
)

type Encoder struct {
	TimeFormat string
}

func (encoder Encoder) Encode(message logger.Message) (string, error) {
	return fmt.Sprintf("%s %s %s", message.Level, message.Time.Format(encoder.TimeFormat), message.Data), nil
}
