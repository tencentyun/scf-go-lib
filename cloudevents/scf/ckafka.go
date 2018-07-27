package scf


// CkafkaEvent received as list
type CkafkaEvent struct {
	Records []CkafkaEventRecord `json:"Records"`
}


//CkafkaEventRecord contains Ckafka Event info
type CkafkaEventRecord struct {
	CkafkaEventEntity    CkafkaEventEntity      `json:"Ckafka"`
}


type CkafkaEventEntity struct {
	Topic    			string     	`json:"topic"`
	Partition 			int64 		`json:"partition"`
	Offset 				int64 		`json:"offset"`
	MsgKey 				string 		`json:"msgKey"`
	MsgBody			 	string 		`json:"msgBody"`
}



