package iwf

type Workflow interface {
	/**
	 * defines the states of the workflow. A state represents a step of the workflow state machine.
	 * A state can execute some commands (activities/signal/timer) and wait for result
	 * See more details in the state definition.
	 */
	GetStates() []StateDef

	/**
	 * defines all the signal methods supported by this workflow.
	 */
	GetActivityTypes() []ActivityTypeDef

	/**
	 * defines all the signal methods supported by this workflow.
	 */
	GetSignalMethods() []SignalMethodDef
	/**
	 * defines all the search attributes supported by this workflow.
	 */
	GetSearchAttributes() []SearchAttributeDef
	/**
	 * defines all the query attributes supported by this workflow.
	 */
	GetQueryAttributes() []QueryAttributeDef
}