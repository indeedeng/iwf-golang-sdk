package iwf

import (
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
)

type StateOptions struct {
	WaitUntilApiSearchAttributesLoadingPolicy *iwfidl.PersistenceLoadingPolicy
	ExecuteApiSearchAttributesLoadingPolicy   *iwfidl.PersistenceLoadingPolicy
	SearchAttributesLoadingPolicy             *iwfidl.PersistenceLoadingPolicy
	WaitUntilApiDataAttributesLoadingPolicy   *iwfidl.PersistenceLoadingPolicy
	ExecuteApiDataAttributesLoadingPolicy     *iwfidl.PersistenceLoadingPolicy
	DataAttributesLoadingPolicy               *iwfidl.PersistenceLoadingPolicy
	WaitUntilApiTimeoutSeconds                *int32
	ExecuteApiTimeoutSeconds                  *int32
	WaitUntilApiRetryPolicy                   *iwfidl.RetryPolicy
	ExecuteApiRetryPolicy                     *iwfidl.RetryPolicy
	WaitUntilApiFailurePolicy                 *iwfidl.WaitUntilApiFailurePolicy
	ExecuteApiFailureProceedState             WorkflowState
}
