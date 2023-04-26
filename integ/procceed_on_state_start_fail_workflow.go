package integ

import "github.com/indeedeng/iwf-golang-sdk/iwf"

type proceedOnStateStartFailWorkflow struct {
	iwf.DefaultWorkflowType
	iwf.EmptyPersistenceSchema
	iwf.EmptyCommunicationSchema
}

func (b proceedOnStateStartFailWorkflow) GetWorkflowStates() []iwf.StateDef {
	return []iwf.StateDef{
		iwf.StartingStateDef(&proceedOnStateStartFailWorkflowState1{output: ""}),
		iwf.NonStartingStateDef(&proceedOnStateStartFailWorkflowState2{output: ""}),
	}
}
