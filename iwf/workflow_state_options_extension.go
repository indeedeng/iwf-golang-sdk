package iwf

import (
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf/ptr"
)

type workflowStateOptionsExtension struct {
	*iwfidl.WorkflowStateOptions
}

func NewWorkflowStateOptionsExtension(stateOptions *iwfidl.WorkflowStateOptions) workflowStateOptionsExtension {
	if stateOptions == nil {
		stateOptions = &iwfidl.WorkflowStateOptions{}
	}

	return workflowStateOptionsExtension{
		WorkflowStateOptions: stateOptions,
	}
}

func (b workflowStateOptionsExtension) GetStateOptions() *iwfidl.WorkflowStateOptions {
	return b.WorkflowStateOptions
}

func (b workflowStateOptionsExtension) SetProceedOnExecuteFailure(recoveryState WorkflowState, recoveryStateOptions *iwfidl.WorkflowStateOptions) *iwfidl.WorkflowStateOptions {
	if recoveryState != nil {
		b.ExecuteApiFailurePolicy = iwfidl.PROCEED_TO_CONFIGURED_STATE.Ptr()
		b.ExecuteApiFailureProceedStateId = ptr.Any(recoveryState.GetStateId())
	}

	if recoveryStateOptions != nil {
		b.ExecuteApiFailureProceedStateOptions = recoveryStateOptions
	}

	return b.WorkflowStateOptions
}
