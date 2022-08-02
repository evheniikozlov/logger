package txt

import (
	"fmt"
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
		now := time.Now()
		expected := fmt.Sprintf("%s %s %s", param.level, now.Format(param.timeLayout), param.data)
		actual, err := Encoder{}.Encode(logger.Message{Level: param.level, Time: now.Format(param.timeLayout), Data: param.data})
		if err != nil {
			t.Fatal(err)
		}
		if actual != expected {
			t.Errorf("actual is %s but expected %s", actual, expected)
		}
	}
}
