package messages

type PingRequest struct {
}

type PingResponse struct {
}

type InvokeRequest_Timestamp struct {
	Seconds int64
	Nanos   int64
}

type InvokeRequest struct {
	Payload               []byte
	RequestId             string
	TraceId               string
	Deadline              InvokeRequest_Timestamp
	InvokedFunctionUnique string
	ClientContext         []byte
	Namespace             string
	FunctionName          string
	FunctionVersion       string
	MemoryLimitInMb       int32
	TimeLimitInMs         int32
        Environment           string
}

type InvokeResponse struct {
	Payload []byte
	Error   *InvokeResponse_Error
}

type InvokeResponse_Error struct {
	Message    string
	Type       string
	StackTrace []*InvokeResponse_Error_StackFrame
	ShouldExit bool
}

type InvokeResponse_Error_StackFrame struct {
	Path  string `json:"path"`
	Line  int32  `json:"line"`
	Label string `json:"label"`
}
