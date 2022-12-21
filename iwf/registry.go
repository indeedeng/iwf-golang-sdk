package iwf

type Registry interface {
	// AddWorkflow registers a workflow
	AddWorkflow(workflow Workflow) error
	// GetWorkflowType returns the workflow type that will be registered
	GetWorkflowType(workflow Workflow) string
	// GetAllWorkflowTypes returns all the workflow types that have been registered
	GetAllWorkflowTypes() []string
}
