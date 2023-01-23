package integ

import (
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type interStateWorkflowState2 struct {
	iwf.DefaultStateIdAndOptions
}

func (b interStateWorkflowState2) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	communication.PublishInterstateChannel(interStateChannel2, 2)
	return iwf.EmptyCommandRequest(), nil
}

func (b interStateWorkflowState2) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	return iwf.DeadEnd, nil
}
