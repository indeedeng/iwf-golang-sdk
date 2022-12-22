package iwf

import "github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"

type WorkflowState interface {
	GetStateId() string

	// Start is the method to execute the commands set up for this state.
	//
	// @param ctx              the context info of this API invocation, like workflow start time, workflowId, etc
	// @param input            the state input
	// @param Persistence      the API for 1) data objects, 2) search attributes and 3) stateLocals 4) recordEvent
	//                         DataObjects and SearchAttributes are defined by {@link Workflow} interface.
	//                         StateLocals are for passing data within the state execution from this start API to {@link #decide} API
	//                         RecordEvent is for storing some tracking info(e.g. RPC call input/output) when executing the API.
	//                         Note that any write API will be recorded to server after the whole start API response is accepted.
	// @param Communication    the API right now only for publishing value to interstate channel
	//                         Note that any write API will be recorded to server after the whole start API response is accepted.
	// @return the requested commands for this state
	///
	Start(ctx WorkflowContext, input Object, persistence Persistence, communication Communication) (*CommandRequest, error)

	// Decide is the method to decide what to do next when requested commands are ready
	//
	// @param ctx              the context info of this API invocation, like workflow start time, workflowId, etc
	// @param input            the state input
	// @param CommandResults   the results of the command that executed by Start
	// @param Persistence      the API for 1) data objects, 2) search attributes and 3) stateLocals 4) recordEvent
	//                         DataObjects and SearchAttributes are defined by {@link Workflow} interface.
	//                         StateLocals are for passing data within the state execution from this start API to {@link #decide} API
	//                         RecordEvent is for storing some tracking info(e.g. RPC call input/output) when executing the API.
	//                         Note that any write API will be recorded to server after the whole start API response is accepted.
	// @param Communication    the API right now only for publishing value to interstate channel
	//                         Note that any write API will be recorded to server after the whole start API response is accepted.
	// @return the decision of what to do next(e.g. transition to next states)
	Decide(ctx WorkflowContext, input Object, commandResults CommandResults, persistence Persistence, communication Communication) (*StateDecision, error)

	// GetStateOptions can just return nil to use the default Options
	// StateOptions is optional configuration to adjust the state behaviors
	GetStateOptions() *iwfidl.WorkflowStateOptions
}
