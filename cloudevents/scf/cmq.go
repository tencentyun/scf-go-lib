package scf

import (
	"encoding/json"
	"strings"
)

// CMQEvent received as list
type CMQEvent struct {
	Records []CMQEventRecord `json:"Records"`
}

//CMQEventRecord contains CMQ Event info
type CMQEventRecord struct {
	CMQEventEntity CMQEventEntity `json:"CMQ"`
}

type CMQEventEntity struct {
	Type             string   `json:"type"`
	TopicOwner       int64    `json:"topicOwner"`
	TopicName        string   `json:"topicName"`
	SubscriptionName string   `json:"subscriptionName"`
	PublishTime      string   `json:"publishTime"`
	MsgID            string   `json:"msgId"`
	RequestID        string   `json:"requestId"`
	MsgBody          string   `json:"msgBody"`
	MsgTag           []string `json:"-"`
}

// UnmarshalJSON implements the json.Unmarshaller interface
func (entity *CMQEventEntity) UnmarshalJSON(data []byte) error {
	type Alias CMQEventEntity
	type Entity struct {
		Alias
		Tags string `json:"msgTag"`
	}
	var e Entity
	err := json.Unmarshal(data, &e)
	if err != nil {
		return err
	}
	*entity = CMQEventEntity(e.Alias)
	entity.MsgTag = strings.Split(e.Tags, ", ")
	return nil
}
