package example

import "github.com/indeedeng/iwf-golang-sdk/iwf"

func NewInitState() iwf.WorkflowState {
	return initState{}
}

type initState struct {
	iwf.WorkflowStateDefaults
}

const keyCustomer = "customer"

func (b initState) WaitUntil(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	var customer string
	input.Get(&customer)
	persistence.SetDataAttribute(keyCustomer, customer)
	return iwf.EmptyCommandRequest(), nil
}

func (b initState) Execute(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	return iwf.GracefulCompletingWorkflow, nil
}
