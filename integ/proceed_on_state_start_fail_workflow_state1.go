package integ

import (
	"errors"

	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type proceedOnStateStartFailWorkflowState1 struct {
	output string
}

func (b *proceedOnStateStartFailWorkflowState1) GetStateId() string {
	return "proceed_on_state_start_fail_workflow_state1"
}

func (b *proceedOnStateStartFailWorkflowState1) WaitUntil(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	var i string
	input.Get(&i)
	b.output = i + "_state1_start"
	return nil, errors.New("")
}

func (b *proceedOnStateStartFailWorkflowState1) Execute(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	b.output += "_state1_decide"
	return iwf.SingleNextState(&proceedOnStateStartFailWorkflowState2{}, b.output), nil
}

func (b *proceedOnStateStartFailWorkflowState1) GetStateOptions() *iwfidl.WorkflowStateOptions {
	return &iwfidl.WorkflowStateOptions{
		WaitUntilApiRetryPolicy: &iwfidl.RetryPolicy{
			InitialIntervalSeconds: iwfidl.PtrInt32(1),
			MaximumAttempts:        iwfidl.PtrInt32(2),
		},
		WaitUntilApiFailurePolicy: iwfidl.PROCEED_ON_FAILURE.Ptr(),
	}
}
