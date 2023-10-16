package integ

import "github.com/indeedeng/iwf-golang-sdk/iwf"

type executeApiFailRecoveryWorkflow struct {
	iwf.WorkflowDefaults
}

func (b executeApiFailRecoveryWorkflow) GetWorkflowStates() []iwf.StateDef {
	return []iwf.StateDef{
		iwf.StartingStateDef(&executeApiFailRecoveryWorkflowState1{}),
		iwf.NonStartingStateDef(&executeApiFailRecoveryWorkflowState2{}),
	}
}
