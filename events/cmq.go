package events

import (
	"encoding/json"
	"fmt"
	"strings"
)

// CMQEvent represents a CMQ event which consists of multiple records
type CMQEvent struct {
	Records []CMQRecord `json:"Records"`
}

// CMQRecord represents a CMQ record
type CMQRecord struct {
	Message CMQMessage `json:"CMQ"`
}

// CMQMessage represents a single CMQ message
type CMQMessage struct {
	ID               string  `json:"msgId"`
	Body             string  `json:"msgBody"`
	Tags             CMQTags `json:"msgTag"`
	PublishTime      string  `json:"publishTime"`
	RequestID        string  `json:"requestId"`
	SubscriptionName string  `json:"subscriptionName"`
	TopicName        string  `json:"topicName"`
	TopicOwner       int64   `json:"topicOwner"`
	Type             string  `json:"type"`
}

// CMQTags represents CMQ routing tags
type CMQTags []string

// UnmarshalJSON implements the json.Unmarshaller interface,
// it handles the JSON String/Array properly
func (tags *CMQTags) UnmarshalJSON(data []byte) error {
	switch data[0] {
	case '"':
		// `"abc"`
		l := len(data)
		d := data[1 : l-1]
		// s = `123`
		s := string(d)
		*tags = CMQTags(strings.Split(s, ", "))
		return nil
	case '[':
		// `["abc", "xyz"]`
		t := []string{}
		err := json.Unmarshal(data, &t)
		if err != nil {
			return err
		}
		*tags = CMQTags(t)
		return nil
	default:
		return fmt.Errorf("unexpected token: `%q`", data[0])
	}
}
