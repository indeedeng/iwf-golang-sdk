package integ

import (
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type basicWorkflowState2 struct{}

const basicWorkflowState2Id = "basicWorkflowState2"

func (b basicWorkflowState2) GetStateId() string {
	return basicWorkflowState2Id
}

func (b basicWorkflowState2) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	return iwf.EmptyCommandRequest(), nil
}

func (b basicWorkflowState2) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	var i int
	err := input.Get(&i)
	if err != nil {
		return nil, err
	}
	return iwf.GracefulCompleteWorkflow(i + 1), nil
}

func (b basicWorkflowState2) GetStateOptions() *iwfidl.WorkflowStateOptions {
	return nil
}
