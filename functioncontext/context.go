package functioncontext

import (
	"context"
	"os"
	"strconv"
)

// FunctionName the name of the current Function
var FunctionName string

// MemoryLimitInMB is the configured memory limit for the current instance of the Function
var MemoryLimitInMB int

// FunctionVersion is the published version of the current instance of the Function
var FunctionVersion string

func init() {

	FunctionName = os.Getenv("FUNCTION_NAME")
	if limit, err := strconv.Atoi(os.Getenv("FUNCTION_MEMORY_SIZE")); err != nil {
		MemoryLimitInMB = 0
	} else {
		MemoryLimitInMB = limit
	}
	FunctionVersion = os.Getenv("FUNCTION_VERSION")
}

// ClientApplication is metadata about the calling application.
type ClientApplication struct {
	InstallationID string `json:"installation_id"`
	AppTitle       string `json:"app_title"`
	AppVersionCode string `json:"app_version_code"`
	AppPackageName string `json:"app_package_name"`
}

// ClientContext is information about the client application passed by the calling application.
type ClientContext struct {
	Client ClientApplication
	Env    map[string]string `json:"env"`
	Custom map[string]string `json:"custom"`
}

// FunctionContext is the set of metadata that is passed for every Invoke.
type FunctionContext struct {
	RequestID             string
	InvokedFunctionUnique string
	ClientContext         ClientContext
	Namespace             string
	FunctionName          string
	FunctionVersion       string
	MemoryLimitInMb       int32
	TimeLimitInMs         int32
        Environment           map[string]string
}

// An unexported type to be used as the key for types in this package.
// This prevents collisions with keys defined in other packages.
type key struct{}

// The key for a FunctionContext in Contexts.
// Users of this package must use functioncontext.NewContext and functioncontext.FromContext
// instead of using this key directly.
var contextKey = &key{}

// NewContext returns a new Context that carries value lc.
func NewContext(parent context.Context, lc *FunctionContext) context.Context {
	return context.WithValue(parent, contextKey, lc)
}

// FromContext returns the FunctionContext value stored in ctx, if any.
func FromContext(ctx context.Context) (*FunctionContext, bool) {
	lc, ok := ctx.Value(contextKey).(*FunctionContext)
	return lc, ok
}
