package iwf

import "github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"

type Registry interface {
	// AddWorkflow registers a workflow
	AddWorkflow(workflow ObjectWorkflow) error
	// AddWorkflows registers multiple workflows
	AddWorkflows(workflows ...ObjectWorkflow) error
	// GetAllRegisteredWorkflowTypes returns all the workflow types that have been registered
	GetAllRegisteredWorkflowTypes() []string

	// below are all for internal implementation
	getWorkflowStartingState(wfType string) WorkflowState
	getWorkflowStateDef(wfType string, id string) StateDef
	getWorkflowRPC(wfType string, rpcMethod string) CommunicationMethodDef
	getWorkflowSignalNameStore(wfType string) map[string]bool
	getWorkflowInternalChannelNameStore(wfType string) map[string]bool
	getWorkflowDataAttributesKeyStore(wfType string) map[string]bool
	getSearchAttributeTypeStore(wfType string) map[string]iwfidl.SearchAttributeValueType
}

func NewRegistry() Registry {
	return &registryImpl{
		workflowStore:            map[string]ObjectWorkflow{},
		workflowStartingState:    map[string]WorkflowState{},
		workflowStateStore:       map[string]map[string]StateDef{},
		signalNameStore:          map[string]map[string]bool{},
		internalChannelNameStore: map[string]map[string]bool{},
		dataAttrsKeyStore:        map[string]map[string]bool{},
		searchAttributeTypeStore: map[string]map[string]iwfidl.SearchAttributeValueType{},
		workflowRPCStore:         map[string]map[string]CommunicationMethodDef{},
	}
}
