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

	// TODO persistence
	comm := newCommunication(w.options.ObjectEncoder, w.registry.getWorkflowInterStateChannelNameStore(request.GetWorkflowType()))
	commandRequest, err := stateDef.State.Start(wfCtx, input, nil, comm)
	if err != nil {
		return nil, err
	}

	err = canNotRequestAndPublishTheSameInterStateChannel(comm.getToPublishInterStateChannel(), commandRequest)
	if err != nil {
		return nil, err
	}

	idlCommandRequest, err := toIdlCommandRequest(commandRequest)
	if err != nil {
		return nil, err
	}
	publishings := toPublishing(comm.getToPublishInterStateChannel())
	resp = &iwfidl.WorkflowStateStartResponse{
		CommandRequest: idlCommandRequest,
	}
	if len(publishings) > 0 {
		resp.PublishToInterStateChannel = publishings
	}
	return resp, nil
}

func toPublishing(channels map[string][]iwfidl.EncodedObject) []iwfidl.InterStateChannelPublishing {
	var res []iwfidl.InterStateChannelPublishing
	for name, l := range channels {
		for _, v := range l {
			res = append(res, iwfidl.InterStateChannelPublishing{
				ChannelName: name,
				Value:       &v,
			})
		}
	}
	return res
}

func canNotRequestAndPublishTheSameInterStateChannel(channelToPublish map[string][]iwfidl.EncodedObject, commandRequest *CommandRequest) error {
	if len(channelToPublish) > 0 && commandRequest != nil {
		for _, cr := range commandRequest.Commands {
			if cr.CommandType == CommandTypeInterStateChannel {
				chName := cr.InterStateChannelCommand.ChannelName
				_, ok := channelToPublish[chName]
				if ok {
					return NewWorkflowDefinitionFmtError("Cannot publish and request for the same interStateChannel: %v", chName)
				}
			}
		}
	}
	return nil
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
	// TODO persistence
	comm := newCommunication(w.options.ObjectEncoder, w.registry.getWorkflowInterStateChannelNameStore(request.GetWorkflowType()))
	decision, err := stateDef.State.Decide(wfCtx, input, commandResults, nil, comm)
	if err != nil {
		return nil, err
	}
	idlDecision, err := toIdlDecision(decision, request.GetWorkflowType(), w.registry, w.options.ObjectEncoder)
	resp = &iwfidl.WorkflowStateDecideResponse{
		StateDecision: idlDecision,
	}
	publishings := toPublishing(comm.getToPublishInterStateChannel())
	if len(publishings) > 0 {
		resp.PublishToInterStateChannel = publishings
	}
	return resp, nil
}
