package iwf

import (
	"context"
	"github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"
	"github.com/iworkflowio/iwf-golang-sdk/iwf/ptr"
)

const (
	WorkflowStateStartApi  = "/api/v1/workflowState/start"
	WorkflowStateDecideApi = "/api/v1/workflowState/decide"
)

type WorkerService interface {
	HandleWorkflowStateStart(ctx context.Context, request iwfidl.WorkflowStateStartRequest) (*iwfidl.WorkflowStateStartResponse, error)
	HandleWorkflowStateDecide(ctx context.Context, request iwfidl.WorkflowStateDecideRequest) (*iwfidl.WorkflowStateDecideResponse, error)
}

func NewWorkerService(registry Registry, options *WorkerOptions) WorkerService {
	if options == nil {
		options = ptr.Any(GetDefaultWorkerOptions())
	}
	return &workerServiceImpl{
		registry: registry.(*registryImpl),
		options:  *options,
	}
}