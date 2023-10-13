package integ

import "github.com/indeedeng/iwf-golang-sdk/iwf"

type executeApiFailRecoveryWorkflow struct {
	iwf.DefaultWorkflowType
	iwf.EmptyPersistenceSchema
	iwf.EmptyCommunicationSchema
}

func (b executeApiFailRecoveryWorkflow) GetWorkflowStates() []iwf.StateDef {
	return []iwf.StateDef{
		iwf.StartingStateDef(&executeApiFailRecoveryWorkflowState1{}),
		iwf.NonStartingStateDef(&executeApiFailRecoveryWorkflowState2{}),
	}
}
