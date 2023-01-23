package iwf

import (
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"reflect"
	"strings"
)

type WorkflowState interface {
	// GetStateId defines the StateId of this workflow state definition.
	// the StateId is being used for WorkerService to choose the right WorkflowState to execute Start/Decide APIs
	// See GetDefaultWorkflowStateId for default value when return empty string.
	// It's the package + struct name of the workflow instance and ignores the import paths and aliases.
	// e.g. if the workflow is from myStruct{} under mywf package, the simple name is just "mywf.myStruct". Underneath, it's from reflect.TypeOf(wf).String().
	//
	// Usually using default value is enough. Unless cases like:
	// 1. You rename the workflowState struct but there is some in-flight state execution still using the old StateId
	// 2. To avoid type name conflicts because the GetDefaultWorkflowStateId is not long enough
	// 3. In case of dynamic workflow state implementation, return customized values instead of using empty string
	GetStateId() string

	// Start is the method to execute the commands set up for this state.
	//
	//  ctx              the context info of this API invocation, like workflow start time, workflowId, etc
	//  input            the state input
	//  Persistence      the API for 1) data objects, 2) search attributes and 3) stateLocals 4) recordEvent
	//                         DataObjects and SearchAttributes are defined by Workflow interface.
	//                         StateLocals are for passing data within the state execution from this start API to {@link #decide} API
	//                         RecordEvent is for storing some tracking info(e.g. RPC call input/output) when executing the API.
	//                         Note that any write API will be recorded to server after the whole start API response is accepted.
	//  Communication    the API right now only for publishing value to interstate channel
	//                         Note that any write API will be recorded to server after the whole start API response is accepted.
	// @return the requested commands for this state
	///
	Start(ctx WorkflowContext, input Object, persistence Persistence, communication Communication) (*CommandRequest, error)

	// Decide is the method to decide what to do next when requested commands are ready
	//
	//  ctx              the context info of this API invocation, like workflow start time, workflowId, etc
	//  input            the state input
	//  CommandResults   the results of the command that executed by Start
	//  Persistence      the API for 1) data objects, 2) search attributes and 3) stateLocals 4) recordEvent
	//                         DataObjects and SearchAttributes are defined by Workflow interface.
	//                         StateLocals are for passing data within the state execution from this start API to {@link #decide} API
	//                         RecordEvent is for storing some tracking info(e.g. RPC call input/output) when executing the API.
	//                         Note that any write API will be recorded to server after the whole start API response is accepted.
	//  Communication    the API right now only for publishing value to interstate channel
	//                         Note that any write API will be recorded to server after the whole start API response is accepted.
	// @return the decision of what to do next(e.g. transition to next states)
	Decide(ctx WorkflowContext, input Object, commandResults CommandResults, persistence Persistence, communication Communication) (*StateDecision, error)

	// GetStateOptions can just return nil to use the default Options
	// StateOptions is optional configuration to adjust the state behaviors
	GetStateOptions() *iwfidl.WorkflowStateOptions
}

// GetFinalWorkflowStateId returns the stateId that will be registered and used
// if the workflowState is from myStruct{} under mywf package, the method returns "mywf.myStruct"
func GetFinalWorkflowStateId(workflowState WorkflowState) string {
	sid := workflowState.GetStateId()
	if sid == "" {
		rt := reflect.TypeOf(workflowState)
		rtStr := strings.TrimLeft(rt.String(), "*")
		return rtStr
	}
	return sid
}

// DefaultStateIdAndOptions is a convenient struct to put into your state implementation to save the boilerplate code. Eg:
// type myStateImpl struct{
//     DefaultStateIdAndOptions
// }
type DefaultStateIdAndOptions struct {
	DefaultStateId
	DefaultStateOptions
}

type DefaultStateId struct{}

func (d DefaultStateId) GetStateId() string {
	return ""
}

type DefaultStateOptions struct{}

func (d DefaultStateOptions) GetStateOptions() *iwfidl.WorkflowStateOptions {
	return nil
}