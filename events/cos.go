package events

// COSEvent represents a COS event which consists of multiple records
type COSEvent struct {
	Records []COSRecord `json:"Records"`
}

// COSRecord represents a COS record
type COSRecord struct {
	Object COSEntity        `json:"cos"`
	Event  COSEventMetadata `json:"event"`
}

// COSEntity represents a COS object in a specific bucket
type COSEntity struct {
	NotificationID string    `json:"cosNotificationId"`
	SchemaVersion  string    `json:"cosSchemaVersion"`
	Bucket         COSBucket `json:"cosBucket"`
	Object         COSObject `json:"cosObject"`
}

// COSBucket represents a COS bucket
type COSBucket struct {
	AppID  string `json:"appid"`
	Region string `json:"region"`
	Name   string `json:"name"`
}

// COSObject represents a COS object
type COSObject struct {
	Name     string            `json:"key"`
	Size     int64             `json:"size"`
	URL      string            `json:"url"`
	Metadata map[string]string `json:"meta"`
	VID      string            `json:"vid"`
}

// COSEventMetadata provides information about the event which creates a COS record
type COSEventMetadata struct {
	RequestID  int64                `json:"reqid"`
	Name       string               `json:"eventName"`
	Queue      string               `json:"eventQueue"`
	Source     string               `json:"eventSource"`
	Timestamp  int64                `json:"eventTime"`
	Version    string               `json:"eventVersion"`
	Parameters COSRequestParameters `json:"requestParameters"`
	Reserved   string               `json:"reservedInfo"`
}

// COSRequestParameters represents the request parameters
type COSRequestParameters struct {
	SourceIP string            `json:"requestSourceIP"`
	Headers  map[string]string `json:"requestHeaders"`
}
