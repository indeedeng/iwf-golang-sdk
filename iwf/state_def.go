package iwf

type StateDef struct {
	State WorkflowState
	// CanStartWorkflow decides whether the state can start a workflow
	CanStartWorkflow bool
}

func StartingStateDef(state WorkflowState) StateDef {
	return StateDef{
		State:            state,
		CanStartWorkflow: true,
	}
}

func NonStartingStateDef(state WorkflowState) StateDef {
	return StateDef{
		State:            state,
		CanStartWorkflow: false,
	}
}