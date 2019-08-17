package events_test

import (
	"encoding/json"
	"testing"

	"github.com/tencentyun/scf-go-lib/events"
)

func TestCMQ(t *testing.T) {
	data := `{"Records": [{"CMQ": {"msgBody": "aaaaaaaaaa", "msgId": "13510798882111490", "msgTag": "aaaa, bbbbb, cccc", "publishTime": "2019-08-16T10:48:49Z", "requestId": "2758374289357404466", "subscriptionName": "ritchiechen-cmq-test", "topicName": "ritchiechen-cmq-test", "topicOwner": 123456, "type": "topic"}}]}`
	var event events.CMQEvent
	err := json.Unmarshal([]byte(data), &event)
	if err != nil {
		t.Fatal(err)
	}
}
