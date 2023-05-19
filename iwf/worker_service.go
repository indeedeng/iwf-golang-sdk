package iwf

import (
	"context"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf/ptr"
)

const (
	WorkflowStateWaitUntilApi = "/api/v1/workflowState/start"
	WorkflowStateExecuteApi   = "/api/v1/workflowState/decide"
	WorkflowWorkerRPCAPI      = "/api/v1/workflowWorker/rpc"
)

type WorkerService interface {
	HandleWorkflowStateWaitUntil(ctx context.Context, request iwfidl.WorkflowStateWaitUntilRequest) (*iwfidl.WorkflowStateWaitUntilResponse, error)
	HandleWorkflowStateExecute(ctx context.Context, request iwfidl.WorkflowStateExecuteRequest) (*iwfidl.WorkflowStateExecuteResponse, error)
	HandleWorkflowWorkerRPC(ctx context.Context, request iwfidl.WorkflowWorkerRpcRequest) (*iwfidl.WorkflowWorkerRpcResponse, error)
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
