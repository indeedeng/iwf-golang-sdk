package iwf

import (
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
)

type RPCOptions struct {
	// default timeout is provided by iwf-server (5s)
	TimeoutSeconds *int
	// default is ALL_WITHOUT_LOCKING
	DataAttributesLoadingPolicy *iwfidl.PersistenceLoadingPolicy
	// default is ALL_WITHOUT_LOCKING
	SearchAttributesLoadingPolicy *iwfidl.PersistenceLoadingPolicy
}
