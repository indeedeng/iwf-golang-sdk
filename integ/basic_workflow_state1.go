package integ

import (
	"github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"
	"github.com/iworkflowio/iwf-golang-sdk/iwf"
)

type basicWorkflowState1 struct{}

const basicWorkflowState1Id = "basicWorkflowState1"

func (b basicWorkflowState1) GetStateId() string {
	return basicWorkflowState1Id
}

func (b basicWorkflowState1) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	return iwf.EmptyCommandRequest(), nil
}

func (b basicWorkflowState1) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	var i int
	err := input.Get(&i)
	if err != nil {
		return nil, err
	}
	return iwf.SingleNextState(basicWorkflowState2Id, i+1), nil
}

func (b basicWorkflowState1) GetStateOptions() *iwfidl.WorkflowStateOptions {
	return nil
}
