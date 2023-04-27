package integ

import (
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type proceedOnStateStartFailWorkflowState2 struct {
	iwf.DefaultStateIdAndOptions
	output string
}

func (b *proceedOnStateStartFailWorkflowState2) WaitUntil(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	var i string
	input.Get(&i)
	b.output = i + "_state2_start"
	return iwf.EmptyCommandRequest(), nil
}

func (b *proceedOnStateStartFailWorkflowState2) Execute(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	b.output += "_state2_decide"
	return iwf.GracefulCompleteWorkflow(b.output), nil
}
