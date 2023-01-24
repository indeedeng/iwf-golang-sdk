package iwf

type StateDef struct {
	State WorkflowState
	// CanStartWorkflow decides whether the state can start a workflow
	CanStartWorkflow bool
}

func StartingStateDef(state WorkflowState) StateDef {
	return NewStartingState(state)
}

// Deprecated: use StartingStateDef instead to be more concise and readable
func NewStartingState(state WorkflowState) StateDef {
	return StateDef{
		State:            state,
		CanStartWorkflow: true,
	}
}

func NonStartingStateDef(state WorkflowState) StateDef {
	return NewNonStartingState(state)
}

// Deprecated: use NonStartingStateDef instead to be more concise and readable
func NewNonStartingState(state WorkflowState) StateDef {
	return StateDef{
		State:            state,
		CanStartWorkflow: false,
	}
}
