package iwf

import (
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf/ptr"
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

func toIdlStateOptions(skipWaitUntil bool, stateOptions *StateOptions) *iwfidl.WorkflowStateOptions {
	if stateOptions == nil {
		stateOptions = &StateOptions{}
	}

	idlStOptions := &iwfidl.WorkflowStateOptions{
		SearchAttributesLoadingPolicy: stateOptions.SearchAttributesLoadingPolicy,
		DataAttributesLoadingPolicy:   stateOptions.DataAttributesLoadingPolicy,
		WaitUntilApiTimeoutSeconds:    stateOptions.WaitUntilApiTimeoutSeconds,
		ExecuteApiTimeoutSeconds:      stateOptions.ExecuteApiTimeoutSeconds,
		WaitUntilApiRetryPolicy:       stateOptions.WaitUntilApiRetryPolicy,
		ExecuteApiRetryPolicy:         stateOptions.ExecuteApiRetryPolicy,
		WaitUntilApiFailurePolicy:     stateOptions.WaitUntilApiFailurePolicy,
	}

	if skipWaitUntil {
		idlStOptions.SkipWaitUntil = ptr.Any(true)
	}

	if stateOptions.ExecuteApiFailureProceedState != nil {
		idlStOptions.ExecuteApiFailurePolicy = iwfidl.PROCEED_TO_CONFIGURED_STATE.Ptr()
		idlStOptions.ExecuteApiFailureProceedStateId = ptr.Any(GetFinalWorkflowStateId(stateOptions.ExecuteApiFailureProceedState))

		proceedStateOptions := stateOptions.ExecuteApiFailureProceedState.GetStateOptions()
		if proceedStateOptions.ExecuteApiFailureProceedState != nil {
			panic("nested failure handling/recovery is not supported: ExecuteApiFailureProceedState cannot have ExecuteApiFailureProceedState")
		}
		idlStOptions.ExecuteApiFailureProceedStateOptions = toIdlStateOptions(ShouldSkipWaitUntilAPI(stateOptions.ExecuteApiFailureProceedState), proceedStateOptions)
	}

	return idlStOptions
}
