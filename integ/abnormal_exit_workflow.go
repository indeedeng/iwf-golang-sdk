package integ

import "github.com/indeedeng/iwf-golang-sdk/iwf"

type abnormalExitWorkflow struct {
	iwf.DefaultWorkflowType
	iwf.EmptyPersistenceSchema
	iwf.EmptyCommunicationSchema
}

func (wf abnormalExitWorkflow) GetWorkflowStates() []iwf.StateDef {
	return []iwf.StateDef{
		iwf.StartingStateDef(&abnormalExitWorkflowState1{}),
	}
}