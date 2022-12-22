package iwf

type StateDef struct {
	State WorkflowState
	// CanStartWorkflow decides whether the state can start a workflow
	CanStartWorkflow bool
}

func NewStartingState(state WorkflowState) StateDef {
	return StateDef{
		State:            state,
		CanStartWorkflow: true,
	}
}

func NewNonStartingState(state WorkflowState) StateDef {
	return StateDef{
		State:            state,
		CanStartWorkflow: false,
	}
}
