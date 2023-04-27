package integ

import (
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type interStateWorkflowState0 struct {
	iwf.DefaultStateIdAndOptions
}

func (b interStateWorkflowState0) WaitUntil(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	return iwf.EmptyCommandRequest(), nil
}

func (b interStateWorkflowState0) Execute(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	return iwf.MultiNextStates(interStateWorkflowState1{}, interStateWorkflowState2{}), nil
}
