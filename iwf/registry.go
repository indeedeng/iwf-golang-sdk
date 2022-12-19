package iwf

type Registry interface {
	AddWorkflow(workflow Workflow)
}
