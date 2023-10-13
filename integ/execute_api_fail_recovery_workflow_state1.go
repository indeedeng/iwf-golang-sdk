package integ

import (
	"errors"

	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type executeApiFailRecoveryWorkflowState1 struct{}

func (b executeApiFailRecoveryWorkflowState1) GetStateId() string {
	return "execute_api_fail_recovery_workflow_state1"
}

func (b executeApiFailRecoveryWorkflowState1) GetStateOptions() *iwfidl.WorkflowStateOptions {
	options := iwf.NewWorkflowStateOptionsExtension(nil).SetProceedOnExecuteFailure(executeApiFailRecoveryWorkflowState2{}, nil)
	options.ExecuteApiRetryPolicy = &iwfidl.RetryPolicy{
		InitialIntervalSeconds: iwfidl.PtrInt32(1),
		MaximumAttempts:        iwfidl.PtrInt32(1),
	}

	return options
}

func (b executeApiFailRecoveryWorkflowState1) WaitUntil(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	return iwf.EmptyCommandRequest(), nil
}

func (b executeApiFailRecoveryWorkflowState1) Execute(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	return nil, errors.New("error")
}
