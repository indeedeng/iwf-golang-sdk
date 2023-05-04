package iwf

import (
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"reflect"
	"runtime"
	"strings"
)

type RPCOptions struct {
	// default timeout is provided by iwf-server (5s)
	TimeoutSeconds *int
	// default is ALL_WITHOUT_LOCKING
	DataAttributesLoadingPolicy *iwfidl.PersistenceLoadingPolicy
	// default is ALL_WITHOUT_LOCKING
	SearchAttributesLoadingPolicy *iwfidl.PersistenceLoadingPolicy
	// default is the
	RPCMethodName *string
}

// GetFinalRPCMethodName returns the final RPC method name
// by default will use the simple method name of the RPC method(without package name)
func GetFinalRPCMethodName(rpc RPC, rpcOptions *RPCOptions) string {
	if rpcOptions != nil && rpcOptions.RPCMethodName != nil {
		return *rpcOptions.RPCMethodName
	}
	fullName := runtime.FuncForPC(reflect.ValueOf(rpc).Pointer()).Name()
	strs := strings.Split(fullName, ".")
	return strs[len(strs)-1]
}
