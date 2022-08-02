package txt

import (
	"encoding/json"
	"github.com/evheniikozlov/logger"
	"testing"
	"time"
)

func TestEncoder_Encode(t *testing.T) {
	params := []struct {
		level      string
		data       string
		timeLayout string
	}{
		{
			level:      logger.Trace,
			data:       "data",
			timeLayout: "2006-01-02 15:04:05",
		},
		{
			level:      logger.Fatal,
			data:       "data2",
			timeLayout: "2006-01-02",
		},
	}
	for _, param := range params {
		message := logger.Message{Level: param.level, Time: time.Now().Format(param.timeLayout), Data: param.data}
		expected, err := json.Marshal(message)
		if err != nil {
			t.Fatal(err)
		}
		actual, err := Encoder{}.Encode(message)
		if err != nil {
			t.Fatal(err)
		}
		if actual != string(expected) {
			t.Errorf("actual is %s but expected %s", actual, string(expected))
		}
	}
}
