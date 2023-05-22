package integ

import (
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type skipWaitUntilState1 struct {
	iwf.WorkflowStateDefaultsNoWaitUntil
}

func (b skipWaitUntilState1) Execute(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	var i int
	input.Get(&i)
	return iwf.SingleNextState(skipWaitUntilState2{}, i+1), nil
}
