package iwf

import (
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
)

type StateOptions struct {
	// apply for both waitUntil and execute API
	DataAttributesLoadingPolicy   *iwfidl.PersistenceLoadingPolicy
	SearchAttributesLoadingPolicy *iwfidl.PersistenceLoadingPolicy
	// below are wait_until API specific options:
	WaitUntilApiTimeoutSeconds                *int32
	WaitUntilApiRetryPolicy                   *iwfidl.RetryPolicy
	WaitUntilApiFailurePolicy                 *iwfidl.WaitUntilApiFailurePolicy
	WaitUntilApiDataAttributesLoadingPolicy   *iwfidl.PersistenceLoadingPolicy
	WaitUntilApiSearchAttributesLoadingPolicy *iwfidl.PersistenceLoadingPolicy
	// below are execute API specific options:
	ExecuteApiTimeoutSeconds                *int32
	ExecuteApiRetryPolicy                   *iwfidl.RetryPolicy
	ExecuteApiFailureProceedState           WorkflowState
	ExecuteApiDataAttributesLoadingPolicy   *iwfidl.PersistenceLoadingPolicy
	ExecuteApiSearchAttributesLoadingPolicy *iwfidl.PersistenceLoadingPolicy
}
