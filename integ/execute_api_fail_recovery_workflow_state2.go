package integ

import (
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type executeApiFailRecoveryWorkflowState2 struct {
	iwf.WorkflowStateDefaultsNoWaitUntil
}

func (b executeApiFailRecoveryWorkflowState2) Execute(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	return iwf.GracefulCompleteWorkflow("this is workflow state 2"), nil
}
