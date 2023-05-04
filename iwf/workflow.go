package iwf

import (
	"reflect"
	"strings"
)

// ObjectWorkflow is the interface to define a workflow definition.
// ObjectWorkflow is a top level concept in iWF. Any object that is long-lasting(at least a few seconds) can be modeled as an "ObjectWorkflow".
type ObjectWorkflow interface {
	// GetWorkflowStates defines the states of the workflow. A state represents a step of the workflow state machine.
	// A state can execute some commands (signal/timer) and wait for result
	// See more details in the WorkflowState interface.
	// It can return an empty list, meaning no states.
	// There can be at most one startingState in the list.
	// If there is no startingState or with the default empty state list, the workflow
	// will not start any state execution after workflow stated. Application can still
	// use RPC to invoke new state execution in the future.
	GetWorkflowStates() []StateDef

	// GetPersistenceSchema defines all the persistence fields for this workflow, this includes:
	//  1. Data objects
	//  2. Search attributes
	//
	// Data objects can be read/upsert in WorkflowState WaitUntil/Execute API
	// Data objects  can also be read by getDataObjects API by external applications using {@link Client}
	//
	// Search attributes can be read/upsert in WorkflowState WaitUntil/Execute API
	// Search attributes can also be read by GetSearchAttributes Client API by external applications.
	// External applications can also use "SearchWorkflow" API to find workflows by SQL-like query
	GetPersistenceSchema() []PersistenceFieldDef

	// GetCommunicationSchema defines all the communication methods for this workflow, this includes
	// 1. Signal channel
	// 2. Interstate channel
	//
	// Signal channel is for external applications to send signal to workflow execution.
	// ObjectWorkflow execution can listen on the signal in the WorkflowState WaitUntil API and receive in
	// the WorkflowState Execute API
	//
	// InterStateChannel is for synchronization communications between WorkflowStates.
	// E.g. WorkflowStateA will continue after receiving a value from WorkflowStateB
	///
	GetCommunicationSchema() []CommunicationMethodDef

	// GetWorkflowType Define the workflowType of this workflow definition.
	// See GetFinalWorkflowType for default value when return empty string.
	// It's the package + struct name of the workflow instance and ignores the import paths and aliases.
	// e.g. if the workflow is from myStruct{} under mywf package, the simple name is just "mywf.myStruct". Underneath, it's from reflect.TypeOf(wf).String().
	//
	// Usually using default value is enough. Unless cases like:
	// 1. To avoid type name conflicts because the GetFinalWorkflowType is not long enough
	// 2. In case of dynamic workflow implementation, return customized values instead of using empty string
	GetWorkflowType() string
}

// RPC is the signature of an RPC of workflow, which will be registered as RPCMethod under CommunicationSchema
type RPC func(ctx WorkflowContext, input Object, persistence Persistence, communication Communication, outputPtr interface{}) error

// GetFinalWorkflowType returns the workflow type that will be registered and used as IwfWorkflowType
// if the workflow is from &myStruct{} or myStruct{} under mywf package, the method returns "mywf.myStruct"
func GetFinalWorkflowType(wf ObjectWorkflow) string {
	wfType := wf.GetWorkflowType()
	if wfType == "" {
		simpleType := getSimpleTypeNameFromReflect(wf)
		return simpleType
	}
	return wfType
}

func getSimpleTypeNameFromReflect(obj interface{}) string {
	rt := reflect.TypeOf(obj)
	rtStr := strings.TrimLeft(rt.String(), "*")
	return rtStr
}


// DefaultWorkflowType is a convenient struct to put into your workflow implementation to save the boilerplate code. Eg:
// type myStateImpl struct{
//     DefaultWorkflowType
// }
type DefaultWorkflowType struct{}

func (d DefaultWorkflowType) GetWorkflowType() string {
	return ""
}

// EmptyPersistenceSchema is a convenient struct to put into your workflow implementation to save the boilerplate code. Eg:
// type myStateImpl struct{
//     EmptyPersistenceSchema
// }
type EmptyPersistenceSchema struct{}

func (d EmptyPersistenceSchema) GetPersistenceSchema() []PersistenceFieldDef {
	return nil
}

// EmptyCommunicationSchema is a convenient struct to put into your workflow implementation to save the boilerplate code. Eg:
// type myStateImpl struct{
//     EmptyCommunicationSchema
// }
type EmptyCommunicationSchema struct{}

func (d EmptyCommunicationSchema) GetCommunicationSchema() []CommunicationMethodDef {
	return nil
}