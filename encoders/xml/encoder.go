package xml

import (
	"encoding/xml"
	"github.com/evheniikozlov/logger"
)

type Encoder struct {
}

func (encoder Encoder) Encode(message logger.Message) (string, error) {
	messageXml, err := xml.Marshal(message)
	if err != nil {
		return "", err
	}
	return string(messageXml), nil
}
