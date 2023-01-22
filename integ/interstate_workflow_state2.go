package integ

import (
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type interStateWorkflowState2 struct{}

const interStateWorkflowState2Id = "interStateWorkflowState2"

func (b interStateWorkflowState2) GetStateId() string {
	return interStateWorkflowState2Id
}

func (b interStateWorkflowState2) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	communication.PublishInterstateChannel(interStateChannel2, 2)
	return iwf.EmptyCommandRequest(), nil
}

func (b interStateWorkflowState2) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	return iwf.DeadEnd, nil
}

func (b interStateWorkflowState2) GetStateOptions() *iwfidl.WorkflowStateOptions {
	return nil
}
