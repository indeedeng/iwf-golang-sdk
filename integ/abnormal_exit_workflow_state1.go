package integ

import (
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type abnormalExitWorkflowState1 struct {
	iwf.WorkflowStateDefaults
}

func (wf abnormalExitWorkflowState1) Execute(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	panic("abnormal exit state")
}

func (wf abnormalExitWorkflowState1) GetStateOptions() *iwf.StateOptions {
	options := &iwf.StateOptions{
		ExecuteApiRetryPolicy: &iwfidl.RetryPolicy{
			InitialIntervalSeconds: iwfidl.PtrInt32(1),
			MaximumAttempts:        iwfidl.PtrInt32(1),
		},
	}

	return options
}
