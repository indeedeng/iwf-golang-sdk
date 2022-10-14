package iwf

type StateDef interface {
	GetState() WorkflowState
	IsStartable() bool
}

type WorkflowState interface {
	GetStateId() string
	GetInputType() NewTypePtr
	/**
	 * Optional configuration to adjust the state behaviors
	 * Default options should work well for most cases
	 */
	GetStateOptions() StateOptions

	/**
	 * Implement this method to execute the commands set up for this state.
	 *
	 * @param input            the state input which is deserialized by dataConverter with {@link #getInputType}
	 * @param queryAttributes  the query attributes that can be used as readOnly
	 * @param searchAttributes the search attributes that can be used as readOnly
	 * @return the requested commands for this step
	 * NOTE: it's readonly here for simplifying the implementation(execute can be reverted in some edge cases),
	 *       We could change to support R+W if necessary.
	 */
	Execute(ctx WorkflowContext, input interface{}, searchAttributes SearchAttributesRO, queryAttributes QueryAttributesRO) (CommandRequest, error)

	/**
	 * Implement this method to decide what to do next when requested commands are ready
	 *
	 * @param input            the state input which is deserialized by dataConverter with {@link #getInputType}
	 * @param commandResults   the results of the command that executed by {@link #execute}
	 * @param queryAttributes  the query attributes that can be used as Read+Write
	 * @param searchAttributes the search attributes that can be used as Read+Write
	 * @return the decision of what to do next(e.g. transition to next states)
	 */
	Decide(ctx WorkflowContext, input interface{}, commandResults CommandResults, searchAttributes SearchAttributesRW, queryAttributes QueryAttributesRW) (StateDecision, error)
}

type StateOptions interface {
	// optional, can just return nil to use the default policy
	GetSearchAttributesLoadingPolicy() AttributeLoadingPolicy
	// optional, can just return nil to use the default policy
	GetQueryAttributesLoadingPolicy() AttributeLoadingPolicy
	/**
	 * when using RequestAnyCommandsCompleted or RequestAnyCommandsClosed
	 * there could be some unfinished commands left to this state. This policy decided whether and how to carry over those unfinished command to
	 * future states. Default to NONE which means no carry over.
	 */
	GetCommandCarryOverPolicy() CommandCarryOverPolicy
}

type StateDecision interface {
	WaitForMoreCommandResults() bool
	GetNextStates() []StateMovement
	GetUpsertSearchAttributes()
	GetUpsertQueryAttributes()
}

func NewStateDecision(movement ...StateMovement)StateDecision{
	return nil
}

func WaitForMoreCommandResults()StateDecision{
	return nil
}

type StateMovement interface {
	GetNextStateId() string
	GetNextStateInput() interface{}
}

func NewStateMovement(nextStateId string) StateMovement{
	return nil
}

func NewStateMovementWithInput(nextStateId string, nextStateInput interface{}) StateMovement{
	return nil
}

type builtInStateMovement struct {
	id string
}

func (m *builtInStateMovement) GetNextStateId() string {
	return m.id
}

func (m *builtInStateMovement) GetNextStateInput() interface{} {
	return nil
}

func completingWorkflowMovement() StateMovement {
	return &builtInStateMovement{
		id: "_SYS_COMPLETING_WORKFLOW",
	}
}

func failingWorkflowMovement() StateMovement {
	return &builtInStateMovement{
		id: "_SYS_FAILING_WORKFLOW",
	}

}

type builtInStateDecision struct {
	movement StateMovement
}

func (b builtInStateDecision) WaitForMoreCommandResults() bool {
	return false
}

func (b builtInStateDecision) GetNextStates() []StateMovement {
	return nil
}

func (b builtInStateDecision) GetUpsertSearchAttributes() {
	return
}

func (b builtInStateDecision) GetUpsertQueryAttributes() {
	return
}

func CompletingWorkflow() StateDecision {
	return &builtInStateDecision{
		movement: completingWorkflowMovement(),
	}
}

func FailingWorkflow() StateDecision {
	return &builtInStateDecision{
		movement: failingWorkflowMovement(),
	}
}

func NewStateDef(state WorkflowState, startable bool) StateDef{
	return nil
}