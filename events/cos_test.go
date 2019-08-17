package events_test

import (
	"encoding/json"
	"testing"

	"github.com/tencentyun/scf-go-lib/events"
)

func TestCOS(t *testing.T) {
	data := `{"Records": [{"cos": {"cosBucket": {"appid": "123456", "name": "ritchiechen-cos-test", "region": "bj"}, "cosNotificationId": "xxx", "cosObject": {"key": "/123456/ritchiechen-cos-test/foobar.zip", "meta": {"Content-Type": "multipart/form-data; boundary=---------------------------b45c43a0b3", "x-cos-request-id": "NWQ1N2IzMGFfMzE0MzIyMDlfNDlhZl8xMzhkYTU4"}, "size": 9755203, "url": "http://ritchiechen-cos-test-123456.cosbj.myqcloud.com/foobar.zip", "vid": "15945fbff94a6476647d1e08b64ed9f71566028555"}, "cosSchemaVersion": "1.0"}, "event": {"eventName": "cos:ObjectCreated:*", "eventQueue": "qcs:0:lambda:bj:appid/123456:ritchiechen-cos-test.ritchiechen-cos-test.$LATEST", "eventSource": "qcs::cos", "eventTime": 1566028555, "eventVersion": "1.0", "reqid": 566750263, "requestParameters": {"requestHeaders": {"Authorization": ""}, "requestSourceIP": "8.8.8.8"}, "reservedInfo": ""}}]}`
	var event events.COSEvent
	err := json.Unmarshal([]byte(data), &event)
	if err != nil {
		t.Fatal(err)
	}
}
