package iwf

import "github.com/iworkflowio/iwf-golang-sdk/iwf/internal"

type Registry interface {
	// AddWorkflow registers a workflow
	AddWorkflow(workflow Workflow) error
	// GetWorkflowType returns the workflow type that will be registered
	GetWorkflowType(workflow Workflow) string
	// GetAllWorkflowTypes returns all the workflow types that have been registered
	GetAllWorkflowTypes() []string
}

func NewRegistry() Registry {
	return internal.NewRegistry()
}