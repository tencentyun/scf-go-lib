package events_test

import (
	"encoding/json"
	"testing"

	"github.com/tencentyun/scf-go-lib/events"
)

func TestTimer(t *testing.T) {
	data := `{"Message": "", "Time": "2019-08-17T05:26:00Z", "TriggerName": "oneminute", "Type": "Timer"}`
	var event events.TimerEvent
	err := json.Unmarshal([]byte(data), &event)
	if err != nil {
		t.Fatal(err)
	}
}
