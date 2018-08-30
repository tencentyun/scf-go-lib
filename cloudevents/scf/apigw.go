package scf

// APIGatewayProxyRequest contains data coming from the API Gateway proxy in integration way
type APIGatewayProxyRequest struct {
	Path                  string                        `json:"path"`       // The url path be called
	Query                 string                        `json:"query"`      // Query string in request
	HTTPMethod            string                        `json:"httpMethod"` // HTTP method
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
	ServiceName string                    `json:"serviceName"` // Service name in API Gateway Service
	Path        string                    `json:"path"`        // Path define in API of Service
	HTTPMethod  string                    `json:"httpMethod"`  // Method define in API of Service
	RequestID   string                    `json:"requestId"`   //
	Stage       string                    `json:"stage"`       // The stage be called
	Identity    APIGatewayRequestIdentity `json:"identity"`    //
	SourceIP    string                    `json:"sourceIp"`    // Source IP of caller
}

// APIGatewayRequestIdentity contains identity information for the request caller.
type APIGatewayRequestIdentity struct {
	SecretID string `json:"secretId,omitempty"`
}
