package iwf

import (
	"reflect"
	"strings"
)

// Workflow is the interface to define a workflow definition.
// Most of the time, the implementation only needs to return static value for each method.
// For a dynamic workflow definition, the implementation can return different values based on different constructor inputs.
type Workflow interface {
	// GetStates defines the states of the workflow. A state represents a step of the workflow state machine.
	// A state can execute some commands (signal/timer) and wait for result
	// See more details in the WorkflowState interface.
	GetStates() []StateDef

	// GetPersistenceSchema defines all the persistence fields for this workflow, this includes:
	//  1. Data objects
	//  2. Search attributes
	//
	// Data objects can be read/upsert in WorkflowState Start/Decide API
	// Data objects  can also be read by getDataObjects API by external applications using {@link Client}
	//
	// Search attributes can be read/upsert in WorkflowState Start/Decide API
	// Search attributes can also be read by GetSearchAttributes Client API by external applications.
	// External applications can also use "SearchWorkflow" API to find workflows by SQL-like query
	GetPersistenceSchema() []PersistenceFieldDef

	// GetCommunicationSchema defines all the communication methods for this workflow, this includes
	// 1. Signal channel
	// 2. Interstate channel
	//
	// Signal channel is for external applications to send signal to workflow execution.
	// Workflow execution can listen on the signal in the WorkflowState Start API and receive in
	// the WorkflowState Decide API
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

// SetLegacyUseStarPrefixInWorkflowTypeForPointerStruct will GetFinalWorkflowType to use "*" as prefix in the workflow type, if the struct is a pointer
// e.g. &myStruct{} will return "*mywf.myStruct"
// this is only for being compatible for workflows running on old SDK versions
func SetLegacyUseStarPrefixInWorkflowTypeForPointerStruct(legacyWorkflows ...Workflow) {
	for _, wf := range legacyWorkflows {
		simpleType := getSimpleTypeNameFromReflect(wf)
		legacyUseStarPrefixInWorkflowTypeForPointerStruct[simpleType] = true
	}
}

var legacyUseStarPrefixInWorkflowTypeForPointerStruct map[string]bool

// GetFinalWorkflowType returns the workflow type that will be registered and used as IwfWorkflowType
// if the workflow is from &myStruct{} or myStruct{} under mywf package, the method returns "mywf.myStruct"
// if SetLegacyUseStarPrefixForPointerStruct, then &myStruct{} will return "*mywf.myStruct"
func GetFinalWorkflowType(wf Workflow) string {
	wfType := wf.GetWorkflowType()
	if wfType == "" {
		legacyType := getLegacyTypeNameFromReflect(wf)
		simpleType := getSimpleTypeNameFromReflect(wf)
		if legacyUseStarPrefixInWorkflowTypeForPointerStruct[simpleType] {
			return legacyType
		}
		return simpleType
	}
	return wfType
}

func getSimpleTypeNameFromReflect(obj interface{}) string {
	rt := reflect.TypeOf(obj)
	rtStr := strings.TrimLeft(rt.String(), "*")
	return rtStr
}

func getLegacyTypeNameFromReflect(obj interface{}) string {
	rt := reflect.TypeOf(obj)
	return rt.String()
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

func (d DefaultWorkflowType) GetPersistenceSchema() []PersistenceFieldDef {
	return nil
}

// EmptyCommunicationSchema is a convenient struct to put into your workflow implementation to save the boilerplate code. Eg:
// type myStateImpl struct{
//     EmptyCommunicationSchema
// }
type EmptyCommunicationSchema struct{}

func (d EmptyPersistenceSchema) GetCommunicationSchema() []CommunicationMethodDef {
	return nil
}