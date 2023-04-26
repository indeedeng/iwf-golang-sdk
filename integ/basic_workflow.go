package integ

import "github.com/indeedeng/iwf-golang-sdk/iwf"

type basicWorkflow struct {
	iwf.DefaultWorkflowType
	iwf.EmptyPersistenceSchema
	iwf.EmptyCommunicationSchema
}

func (b basicWorkflow) GetWorkflowStates() []iwf.StateDef {
	return []iwf.StateDef{
		iwf.StartingStateDef(&basicWorkflowState1{}),
		iwf.NonStartingStateDef(&basicWorkflowState2{}),
	}
}
