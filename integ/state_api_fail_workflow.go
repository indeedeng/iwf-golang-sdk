package integ

import "github.com/indeedeng/iwf-golang-sdk/iwf"

type stateApiFailWorkflow struct {
	iwf.DefaultWorkflowType
	iwf.EmptyCommunicationSchema
	iwf.EmptyPersistenceSchema
}

func (b stateApiFailWorkflow) GetStates() []iwf.StateDef {
	return []iwf.StateDef{
		iwf.StartingStateDef(&stateApiFailWorkflowState1{}),
	}
}
