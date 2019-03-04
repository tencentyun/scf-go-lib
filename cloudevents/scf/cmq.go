package scf


// CMQEvent received as list
type CMQEvent struct {
	Records []CMQEventRecord `json:"Records"`
}


//CMQEventRecord contains CMQ Event info
type CMQEventRecord struct {
	CMQEventEntity    CMQEventEntity      `json:"CMQ"`
}


type CMQEventEntity struct {
	Type    			string     	`json:"type"`
	TopicOwner 			int64 		`json:"topicOwner"`
	TopicName 			string 		`json:"topicName"`
	SubscriptionName 	string 		`json:"subscriptionName"`
	PublishTime 		string 		`json:"publishTime"`
	MsgID			 	string 		`json:"msgId"`
	RequestID	 		string 		`json:"requestId"`
	MsgBody			 	string 		`json:"msgBody"`
	MsgTag		 		string 	`json:"msgTag"`
}

