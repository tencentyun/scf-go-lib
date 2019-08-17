package events

// CkafkaEvent represents a Ckafka event which consists of multiple records
type CkafkaEvent struct {
	Records []CkafkaRecord `json:"Records"`
}

// CkafkaRecord represents a Ckafka record
type CkafkaRecord struct {
	Message CkafkaMessage `json:"Ckafka"`
}

// CkafkaMessage represents a single Ckafka message
type CkafkaMessage struct {
	Topic     string `json:"topic"`
	Partition int64  `json:"partition"`
	Offset    int64  `json:"offset"`
	Key       string `json:"msgKey"`
	Body      string `json:"msgBody"`
}
