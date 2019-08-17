package events_test

import (
	"encoding/json"
	"testing"

	"github.com/tencentyun/scf-go-lib/events"
)

func TestCKafka(t *testing.T) {
	data := `{"Records": [{"Ckafka": {"msgBody": "hello world", "msgKey": "ckafka-test-key", "offset": 7, "partition": 0, "topic": "ritchiechen-ckafka-test"}}]}`
	var event events.CkafkaEvent
	err := json.Unmarshal([]byte(data), &event)
	if err != nil {
		t.Fatal(err)
	}
}
