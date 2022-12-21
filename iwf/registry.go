package iwf

import "github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"

type Registry interface {
	// AddWorkflow registers a workflow
	AddWorkflow(workflow Workflow) error
	// GetWorkflowType returns the workflow type that will be registered
	GetWorkflowType(workflow Workflow) string
	// GetAllWorkflowTypes returns all the workflow types that have been registered
	GetAllWorkflowTypes() []string
}

func NewRegistry() Registry {
	return &registry{
		workflowStore:              map[string]Workflow{},
		workflowStateStore:         map[string]map[string]StateDef{},
		signalNameStore:            map[string]map[string]bool{},
		interStateChannelNameStore: map[string]map[string]bool{},
		dataObjectKeyStore:         map[string]map[string]bool{},
		searchAttributeTypeStore:   map[string]map[string]iwfidl.SearchAttributeValueType{},
	}
}
