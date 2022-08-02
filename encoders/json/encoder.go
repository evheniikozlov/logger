package json

import (
	"encoding/json"
	"github.com/evheniikozlov/logger"
)

type Encoder struct {
}

func (encoder Encoder) Encode(message logger.Message) (string, error) {
	messageJson, err := json.Marshal(message)
	if err != nil {
		return "", err
	}
	return string(messageJson), nil
}
