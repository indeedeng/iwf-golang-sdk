package iwf

import "github.com/cadence-oss/iwf-golang-sdk/iwf/command"

type StateDef struct {
	State WorkflowState
	// CanStartWorkflow decides whether the state can start a workflow
	CanStartWorkflow bool
}

type WorkflowState interface {
	GetStateId() string

	// Start is the method to execute the commands set up for this state.
	//
	// @param ctx              the context info of this API invocation, like workflow start time, workflowId, etc
	// @param input            the state input which is deserialized by ObjectEncoder
	// @param Persistence      the API for 1) data objects, 2) search attributes and 3) stateLocals 4) recordEvent
	//                         DataObjects and SearchAttributes are defined by {@link Workflow} interface.
	//                         StateLocals are for passing data within the state execution from this start API to {@link #decide} API
	//                         RecordEvent is for storing some tracking info(e.g. RPC call input/output) when executing the API.
	//                         Note that any write API will be recorded to server after the whole start API response is accepted.
	// @param Communication    the API right now only for publishing value to interstate channel
	//                         Note that any write API will be recorded to server after the whole start API response is accepted.
	// @return the requested commands for this state
	///
	Start(ctx WorkflowContext, input interface{}, persistence Persistence, communication Communication) (CommandRequest, error)

	// Decide is the method to decide what to do next when requested commands are ready
	//
	// @param ctx              the context info of this API invocation, like workflow start time, workflowId, etc
	// @param input            the state input which is deserialized by ObjectEncoder
	// @param CommandResults   the results of the command that executed by Start
	// @param Persistence      the API for 1) data objects, 2) search attributes and 3) stateLocals 4) recordEvent
	//                         DataObjects and SearchAttributes are defined by {@link Workflow} interface.
	//                         StateLocals are for passing data within the state execution from this start API to {@link #decide} API
	//                         RecordEvent is for storing some tracking info(e.g. RPC call input/output) when executing the API.
	//                         Note that any write API will be recorded to server after the whole start API response is accepted.
	// @param Communication    the API right now only for publishing value to interstate channel
	//                         Note that any write API will be recorded to server after the whole start API response is accepted.
	// @return the decision of what to do next(e.g. transition to next states)
	Decide(ctx WorkflowContext, input interface{}, commandResults CommandResults, persistence Persistence, communication Communication) (StateDecision, error)

	// GetStateOptions can just return nil to use the default Options
	// StateOptions is optional configuration to adjust the state behaviors
	GetStateOptions() *StateOptions
}

type StateOptions struct {
	// SearchAttributesLoadingPolicy decides what and how search attributes will be loaded into this start execution
	SearchAttributesLoadingPolicy PersistenceLoadingPolicy
	// DataObjectsLoadingPolicy decides what and how data object will be loaded into this start execution
	DataObjectsLoadingPolicy PersistenceLoadingPolicy
	// CommandCarryOverPolicy decides how to carry over remaining commands into next states
	CommandCarryOverPolicy command.CarryOverPolicy
}
