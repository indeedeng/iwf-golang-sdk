package iwf

import (
	"github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"
	"github.com/iworkflowio/iwf-golang-sdk/iwf/ptr"
)

const (
	WorkflowStateStartApi  = "/api/v1/workflowState/start"
	WorkflowStateDecideApi = "/api/v1/workflowState/decide"
)

type WorkerService interface {
	HandleWorkflowStateStart(request iwfidl.WorkflowStateStartRequest) (iwfidl.WorkflowStateStartResponse, error)
	HandleWorkflowStateDecide(request iwfidl.WorkflowStateDecideRequest) (iwfidl.WorkflowStateDecideResponse, error)
}

func NewWorkerService(registry Registry, options *WorkerOptions) WorkerService {
	if options == nil {
		options = ptr.Any(GetDefaultWorkerOptions())
	}
	return &workerServiceImpl{
		registry: registry,
		options:  *options,
	}
}