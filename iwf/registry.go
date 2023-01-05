package iwf

import "github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"

type Registry interface {
	// AddWorkflow registers a workflow
	AddWorkflow(workflow Workflow) error
	// AddWorkflows registers multiple workflows
	AddWorkflows(workflows ...Workflow) error
	// GetAllRegisteredWorkflowTypes returns all the workflow types that have been registered
	GetAllRegisteredWorkflowTypes() []string

	// below are all for internal implementation
	getWorkflowStateDef(wfType string, id string) StateDef
	getWorkflowSignalNameStore(wfType string) map[string]bool
	getWorkflowInterStateChannelNameStore(wfType string) map[string]bool
	getWorkflowDataObjectKeyStore(wfType string) map[string]bool
	getSearchAttributeTypeStore(wfType string) map[string]iwfidl.SearchAttributeValueType
}

func NewRegistry() Registry {
	return &registryImpl{
		workflowStore:              map[string]Workflow{},
		workflowStateStore:         map[string]map[string]StateDef{},
		signalNameStore:            map[string]map[string]bool{},
		interStateChannelNameStore: map[string]map[string]bool{},
		dataObjectKeyStore:         map[string]map[string]bool{},
		searchAttributeTypeStore:   map[string]map[string]iwfidl.SearchAttributeValueType{},
	}
}
