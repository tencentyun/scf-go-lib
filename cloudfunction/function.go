package cloudfunction

import (
	"context"
	"encoding/json"
	"reflect"
	"time"
        "os"
	"github.com/tencentyun/scf-go-lib/cloudfunction/messages"
	"github.com/tencentyun/scf-go-lib/functioncontext"
)

type Function struct {
	handler Handler
}

func (fn *Function) Ping(req *messages.PingRequest, response *messages.PingResponse) error {
	*response = messages.PingResponse{}
	return nil
}

func (fn *Function) Invoke(req *messages.InvokeRequest, response *messages.InvokeResponse) error {
	defer func() {
		if err := recover(); err != nil {
			panicInfo := getPanicInfo(err)
			response.Error = &messages.InvokeResponse_Error{
				Message:    panicInfo.Message,
				Type:       getErrorType(err),
				StackTrace: panicInfo.StackTrace,
				ShouldExit: true,
			}
		}
	}()

	deadline := time.Unix(req.Deadline.Seconds, req.Deadline.Nanos).UTC()
	invokeContext, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	lc := &functioncontext.FunctionContext{
		RequestID:             req.RequestId,
		InvokedFunctionUnique: req.InvokedFunctionUnique,
		Namespace:             req.Namespace,
		FunctionName:          req.FunctionName,
		FunctionVersion:       req.FunctionVersion,
		MemoryLimitInMb:       req.MemoryLimitInMb,
		TimeLimitInMs:         req.TimeLimitInMs,
	}

	if len(req.ClientContext) > 0 {
		if err := json.Unmarshal(req.ClientContext, &lc.ClientContext); err != nil {
			response.Error = functionErrorResponse(err)
			return nil
		}
	}

        if len(req.Environment) > 0 {
		if err := json.Unmarshal([]byte(req.Environment), &lc.Environment); err != nil {
			response.Error = functionErrorResponse(err)
			return nil
		}
                for key, value := range lc.Environment {
                    os.Setenv(key, value)
                }
	}

	invokeContext = functioncontext.NewContext(invokeContext, lc)

	payload, err := fn.handler.Invoke(invokeContext, req.Payload)
	if err != nil {
		response.Error = functionErrorResponse(err)
		return nil
	}

	response.Payload = payload

	return nil
}

func getErrorType(err interface{}) string {
	errorType := reflect.TypeOf(err)
	if errorType.Kind() == reflect.Ptr {
		return errorType.Elem().Name()
	}
	return errorType.Name()
}

func functionErrorResponse(invokeError error) *messages.InvokeResponse_Error {
	var errorName string
	if errorType := reflect.TypeOf(invokeError); errorType.Kind() == reflect.Ptr {
		errorName = errorType.Elem().Name()
	} else {
		errorName = errorType.Name()
	}
	return &messages.InvokeResponse_Error{
		Message: invokeError.Error(),
		Type:    errorName,
	}
}
