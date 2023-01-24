package integ

import (
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type basicWorkflowState1 struct {
	iwf.DefaultStateIdAndOptions
}

func (b basicWorkflowState1) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	if ctx.GetAttempt() <= 0 {
		panic("attempt should be greater than zero")
	}
	if ctx.GetFirstAttemptTimestampSeconds() <= 0 {
		panic("GetFirstAttemptTimestampSeconds should be greater than zero")
	}
	return iwf.EmptyCommandRequest(), nil
}

func (b basicWorkflowState1) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	if ctx.GetAttempt() <= 0 {
		panic("attempt should be greater than zero")
	}
	if ctx.GetFirstAttemptTimestampSeconds() <= 0 {
		panic("GetFirstAttemptTimestampSeconds should be greater than zero")
	}
	var i int
	input.Get(&i)
	return iwf.SingleNextState(basicWorkflowState2{}, i+1), nil
}
