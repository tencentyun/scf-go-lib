package scf

// APIGatewayProxyRequest contains data coming from the API Gateway proxy in integration way
type APIGatewayProxyRequest struct {
	Path                  string                        `json:"path"`        // The url path be called
	QueryString           map[string]string             `json:"queryString"` // Query string in request
	HTTPMethod            string                        `json:"httpMethod"`  // HTTP method
	Headers               map[string]string             `json:"headers"`
	QueryStringParameters map[string]string             `json:"queryStringParameters,omitempty"`
	PathParameters        map[string]string             `json:"pathParameters,omitempty"`
	HeaderParameters      map[string]string             `json:"headerParameters,omitempty"`
	StageVariables        map[string]string             `json:"stageVariables,omitempty"`
	RequestContext        APIGatewayProxyRequestContext `json:"requestContext"`
	Body                  string                        `json:"body,omitempty"`
	IsBase64Encoded       bool                          `json:"isBase64Encoded,omitempty"`
}

// APIGatewayProxyResponse contains the response to be returned to API Gateway to answer the request in integration way
type APIGatewayProxyResponse struct {
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	Body            string            `json:"body"`
	IsBase64Encoded bool              `json:"isBase64Encoded,omitempty"`
}

// APIGatewayProxyRequestContext contains the information of service and api config in api gateway
type APIGatewayProxyRequestContext struct {
	ServiceID       string                    `json:"serviceId"`
	Path            string                    `json:"path"`
	HTTPMethod      string                    `json:"httpMethod"`
	RequestID       string                    `json:"requestId"`
	Stage           string                    `json:"stage"`
	Identity        APIGatewayRequestIdentity `json:"identity"`
	SourceIP        string                    `json:"sourceIp"`
	WebsocketEnable bool                      `json:"websocketEnable,omitempty"`
}

// APIGatewayRequestIdentity contains identity information for the request caller.
type APIGatewayRequestIdentity struct {
	SecretID string `json:"secretId,omitempty"`
}

// APIGatewayWebSocketActionType defines the type of websocket action
type APIGatewayWebSocketActionType string

// APIGatewayWebSocketActionType define values
const (
	Connecting APIGatewayWebSocketActionType = "connecting"
	Closing    APIGatewayWebSocketActionType = "closing"
	DataSend   APIGatewayWebSocketActionType = "data send"
	DataRecv   APIGatewayWebSocketActionType = "data recv"
)

// APIGatewayWebSocketConnection contains websocket connecting info.
type APIGatewayWebSocketConnection struct {
	Action                 APIGatewayWebSocketActionType `json:"action"`
	SecConnectionID        string                        `json:"secConnectionID"`
	SecWebSocketProtocol   string                        `json:"secWebSocketProtocol"`
	SecWebSocketExtensions string                        `json:"secWebSocketExtensions"`
}

// APIGatewayWebSocketAction contains websocket send and recv data or action info.
type APIGatewayWebSocketAction struct {
	Action          APIGatewayWebSocketActionType `json:"action"`
	SecConnectionID string                        `json:"secConnectionID"`
	DataType        string                        `json:"dataType"`
	Data            string                        `json:"data"`
}

// APIGatewayWebSocketConnectionRequest contains connection info send to cloud function
type APIGatewayWebSocketConnectionRequest struct {
	RequestContext APIGatewayProxyRequestContext `json:"requestContext"`
	WebSocketConn  APIGatewayWebSocketConnection `json:"websocket"`
}

// APIGatewayWebSocketConnectionResponse contains info need to response to api gateway
type APIGatewayWebSocketConnectionResponse struct {
	ErrNumber     int                           `json:"errNo"`
	ErrMesg       string                        `json:"errMsg"`
	WebSocketConn APIGatewayWebSocketConnection `json:"websocket"`
}
