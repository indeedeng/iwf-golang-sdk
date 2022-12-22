package iwf

import (
	"context"
	"github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"
)

type workerServiceImpl struct {
	registry Registry
	options  WorkerOptions
}

func (w *workerServiceImpl) HandleWorkflowStateStart(ctx context.Context, request iwfidl.WorkflowStateStartRequest) (resp *iwfidl.WorkflowStateStartResponse, retErr error) {
	defer func() { captureStateExecutionError(recover(), &retErr) }()

	stateDef := w.registry.getWorkflowStateDef(request.GetWorkflowType(), request.GetWorkflowStateId())
	input := NewObject(request.StateInput, w.options.ObjectEncoder)
	reqContext := request.GetContext()
	wfCtx := newWorkflowContext(ctx, reqContext.GetWorkflowId(), reqContext.GetWorkflowRunId(), reqContext.GetStateExecutionId(), reqContext.GetWorkflowStartedTimestamp())

	// TODO persistence, communication
	commandRequest, err := stateDef.State.Start(wfCtx, input, nil, nil)
	if err != nil {
		return nil, err
	}

	idlCommandRequest, err := toIdlCommandRequest(commandRequest)
	if err != nil {
		return nil, err
	}
	return &iwfidl.WorkflowStateStartResponse{
		CommandRequest: idlCommandRequest,
	}, nil
}

func (w *workerServiceImpl) HandleWorkflowStateDecide(ctx context.Context, request iwfidl.WorkflowStateDecideRequest) (resp *iwfidl.WorkflowStateDecideResponse, retErr error) {
	defer func() { captureStateExecutionError(recover(), &retErr) }()

	stateDef := w.registry.getWorkflowStateDef(request.GetWorkflowType(), request.GetWorkflowStateId())
	input := NewObject(request.StateInput, w.options.ObjectEncoder)
	reqContext := request.GetContext()
	wfCtx := newWorkflowContext(ctx, reqContext.GetWorkflowId(), reqContext.GetWorkflowRunId(), reqContext.GetStateExecutionId(), reqContext.GetWorkflowStartedTimestamp())

	commandResults, err := fromIdlCommandResults(request.CommandResults, w.options.ObjectEncoder)
	if err != nil {
		return nil, err
	}
	// TODO persistence, communication
	decision, err := stateDef.State.Decide(wfCtx, input, commandResults, nil, nil)
	if err != nil {
		return nil, err
	}
	idlDecision, err := toIdlDecision(decision, request.GetWorkflowType(), w.registry, w.options.ObjectEncoder)
	return &iwfidl.WorkflowStateDecideResponse{
		StateDecision: idlDecision,
	}, nil
}
