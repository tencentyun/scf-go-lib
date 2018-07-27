package scf

import (
	"time"
)

// COSEvent received as list
type COSEvent struct {
	Records []COSEventRecord `json:"Records"`
}

//COSEventRecord contains Event info and  COS info
type COSEventRecord struct {
	Event		      COSEventEntity      `json:"event"`
	COS               COSEntity           `json:"cos"`
}


type COSEventEntity struct {
	EventName         string         	     `json:"eventName"`
	EventVersion      string         	     `json:"eventVersion"`
	EventSource       string        	     `json:"eventSource"`
	EventTime         time.Time     	     `json:"eventTime"`

	EventQueue        string        	     `json:"eventQueue"`
	RequestParameters COSRequestParameters 	 `json:"requestParameters"`
	ReservedInfo      string 				 `json:"reservedInfo"`
	RequestID		  int64					 `json:"reqid"`
}

type COSRequestParameters struct {
	SourceIP 	string 				`json:"requestSourceIP"`
	Headers 	map[string]string   `json:"requestHeaders"`
}


type COSEntity struct {
	SchemaVersion   string    `json:"cosSchemaVersion"`
	NotificationID  string    `json:"cosNotificationId"`
	Bucket          COSBucket `json:"bucket"`
	Object          COSObject `json:"object"`
}

type COSBucket struct {
	Name          string         `json:"name"`
	AppID         string         `json:"appid"`
	Region        string         `json:"region"`
}

type COSObject struct {
	Key           string 			`json:"key"`
	Size          int64  			`json:"size"`
	URL           string 			`json:"url"`
	Metadata      map[string]string `json:"meta"`
	VID           string            `json:"vid"`
}
