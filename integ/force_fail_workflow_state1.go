package integ

import (
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type forceFailWorkflowState1 struct {
	iwf.DefaultStateId
	iwf.DefaultStateOptions
}

func (b forceFailWorkflowState1) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	return iwf.EmptyCommandRequest(), nil
}

func (b forceFailWorkflowState1) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	return iwf.ForceFailWorkflow("a failing message"), nil
}
