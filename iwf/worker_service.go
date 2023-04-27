package iwf

import (
	"context"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf/ptr"
)

const (
	WorkflowStateStartApi  = "/api/v1/workflowState/start"
	WorkflowStateDecideApi = "/api/v1/workflowState/decide"
)

type WorkerService interface {
	HandleWorkflowStateWaitUntil(ctx context.Context, request iwfidl.WorkflowStateWaitUntilRequest) (*iwfidl.WorkflowStateWaitUntilResponse, error)
	HandleWorkflowStateExecute(ctx context.Context, request iwfidl.WorkflowStateExecuteRequest) (*iwfidl.WorkflowStateExecuteResponse, error)
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