package integ

import (
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type interStateWorkflowState0 struct{}

const interStateWorkflowState0Id = "interStateWorkflowState0"

func (b interStateWorkflowState0) GetStateId() string {
	return interStateWorkflowState0Id
}

func (b interStateWorkflowState0) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	return iwf.EmptyCommandRequest(), nil
}

func (b interStateWorkflowState0) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	return iwf.MultiNextStatesByStateIds(interStateWorkflowState1Id, interStateWorkflowState2Id), nil
}

func (b interStateWorkflowState0) GetStateOptions() *iwfidl.WorkflowStateOptions {
	return nil
}
