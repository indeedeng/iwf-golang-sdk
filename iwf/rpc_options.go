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

	// Only used when workflow has enabled CachingDataAttributesByMemo (see PersistenceSchemaOptions)
	// By default, it's false for high throughput support. Flip to true to bypass the caching for a strong consistent read
	BypassCachingForStrongConsistency bool
}
