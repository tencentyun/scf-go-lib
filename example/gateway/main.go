package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"./httprouter"

	"github.com/tencentyun/scf-go-lib/cloudevents/scf"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

func newJsonResponse() scf.APIGatewayProxyResponse {
	var headers map[string]string = make(map[string]string)
	headers["Content-Type"] = "application/json"
	return scf.APIGatewayProxyResponse{Headers: headers, IsBase64Encoded: false}
}

func hello(req scf.APIGatewayProxyRequest) (scf.APIGatewayProxyResponse, error) {
	var headers map[string]string = make(map[string]string)
	headers["Content-Type"] = "application/json"
	body := fmt.Sprintf("Hello world")
	response := scf.APIGatewayProxyResponse{StatusCode: 200, Headers: headers, Body: body, IsBase64Encoded: false}
	return response, nil
}

func world(req scf.APIGatewayProxyRequest) (scf.APIGatewayProxyResponse, error) {
	name := req.PathParameters["name"]
	response := newJsonResponse()
	var body map[string]string = make(map[string]string)
	body["body"] = fmt.Sprintf("hello, %s", name)
	json_body, _ := json.Marshal(body)
	response.Body = string(json_body)
	response.StatusCode = http.StatusCreated
	return response, nil
}

func runServer(ctx context.Context, req scf.APIGatewayProxyRequest) (scf.APIGatewayProxyResponse, error) {
	router := httprouter.New()
	router.GET("/", hello)
	router.GET("/hello/:name", world)

	return router.ServeHTTP(ctx, req)
}

func main() {
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(runServer)
}
