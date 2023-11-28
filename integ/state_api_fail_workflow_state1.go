package integ

import (
	"fmt"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type stateApiFailWorkflowState1 struct {
	iwf.DefaultStateId
}

func (b stateApiFailWorkflowState1) WaitUntil(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	return nil, fmt.Errorf("test api failing")
}

func (b stateApiFailWorkflowState1) Execute(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	return iwf.ForceFailWorkflow("a failing message"), nil
}

func (b stateApiFailWorkflowState1) GetStateOptions() *iwf.StateOptions {
	return &iwf.StateOptions{
		WaitUntilApiRetryPolicy: &iwfidl.RetryPolicy{
			MaximumAttempts: iwfidl.PtrInt32(1),
		},
	}
}
