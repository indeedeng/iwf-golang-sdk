package integ

import (
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type executeApiFailRecoveryWorkflowState2 struct {
	iwf.DefaultStateOptions
}

func (b executeApiFailRecoveryWorkflowState2) GetStateId() string {
	return "execute_api_fail_recovery_workflow_state2"
}

func (b executeApiFailRecoveryWorkflowState2) WaitUntil(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	return iwf.EmptyCommandRequest(), nil
}

func (b executeApiFailRecoveryWorkflowState2) Execute(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	return iwf.GracefulCompleteWorkflow("this is workflow state 2"), nil
}
