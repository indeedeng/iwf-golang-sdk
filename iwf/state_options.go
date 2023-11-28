package iwf

import (
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
)

type StateOptions struct {
	SearchAttributesLoadingPolicy *iwfidl.PersistenceLoadingPolicy
	DataAttributesLoadingPolicy   *iwfidl.PersistenceLoadingPolicy
	WaitUntilApiTimeoutSeconds    *int32
	ExecuteApiTimeoutSeconds      *int32
	WaitUntilApiRetryPolicy       *iwfidl.RetryPolicy
	ExecuteApiRetryPolicy         *iwfidl.RetryPolicy
	WaitUntilApiFailurePolicy     *iwfidl.WaitUntilApiFailurePolicy
	ExecuteApiFailureProceedState WorkflowState
}

func (s *StateOptions) SetProceedOnExecuteFailure(state WorkflowState) *StateOptions {
	s.ExecuteApiFailureProceedState = state
	return s
}