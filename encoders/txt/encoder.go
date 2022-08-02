package txt

import (
	"fmt"
	"github.com/evheniikozlov/logger"
)

type Encoder struct {
}

func (encoder Encoder) Encode(message logger.Message) (string, error) {
	return fmt.Sprintf("%s %s %s", message.Level, message.Time, message.Data), nil
}
