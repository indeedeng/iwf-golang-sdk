package iwf

import (
	"context"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
)

type workerServiceImpl struct {
	registry Registry
	options  WorkerOptions
}

func (w *workerServiceImpl) HandleWorkflowStateStart(ctx context.Context, request iwfidl.WorkflowStateStartRequest) (resp *iwfidl.WorkflowStateStartResponse, retErr error) {
	defer func() { captureStateExecutionError(recover(), &retErr) }()

	wfType := request.GetWorkflowType()
	stateDef := w.registry.getWorkflowStateDef(wfType, request.GetWorkflowStateId())
	input := NewObject(request.StateInput, w.options.ObjectEncoder)
	reqContext := request.GetContext()
	wfCtx := newWorkflowContext(
		ctx, reqContext.GetWorkflowId(), reqContext.GetWorkflowRunId(), reqContext.GetStateExecutionId(), reqContext.GetWorkflowStartedTimestamp(),
		int(reqContext.GetAttempt()), reqContext.GetFirstAttemptTimestamp())

	pers, err := newPersistence(w.options.ObjectEncoder, w.registry.getWorkflowDataObjectKeyStore(wfType), w.registry.getSearchAttributeTypeStore(wfType), request.DataObjects, request.SearchAttributes, nil)
	if err != nil {
		return nil, err
	}
	comm := newCommunication(w.options.ObjectEncoder, w.registry.getWorkflowInterStateChannelNameStore(wfType))
	commandRequest, err := stateDef.State.Start(wfCtx, input, pers, comm)
	if err != nil {
		return nil, err
	}

	err = canNotRequestAndPublishTheSameInterStateChannel(comm.GetToPublishInterStateChannel(), commandRequest)
	if err != nil {
		return nil, err
	}

	idlCommandRequest, err := toIdlCommandRequest(commandRequest)
	if err != nil {
		return nil, err
	}
	publishings := toPublishing(comm.GetToPublishInterStateChannel())
	resp = &iwfidl.WorkflowStateStartResponse{
		CommandRequest: idlCommandRequest,
	}
	if len(publishings) > 0 {
		resp.PublishToInterStateChannel = publishings
	}
	dataObjectsToReturn, stateLocalToReturn, recordedEvents, upsertSearchAttributes := pers.GetToReturn()
	if len(dataObjectsToReturn) > 0 {
		resp.UpsertDataObjects = dataObjectsToReturn
	}
	if len(stateLocalToReturn) > 0 {
		resp.UpsertStateLocals = stateLocalToReturn
	}
	if len(recordedEvents) > 0 {
		resp.RecordEvents = recordedEvents
	}
	if len(upsertSearchAttributes) > 0 {
		resp.UpsertSearchAttributes = upsertSearchAttributes
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
					return NewWorkflowDefinitionErrorFmt("Cannot publish and request for the same interStateChannel: %v", chName)
				}
			}
		}
	}
	return nil
}

func (w *workerServiceImpl) HandleWorkflowStateDecide(ctx context.Context, request iwfidl.WorkflowStateDecideRequest) (resp *iwfidl.WorkflowStateDecideResponse, retErr error) {
	defer func() { captureStateExecutionError(recover(), &retErr) }()

	wfType := request.GetWorkflowType()
	stateDef := w.registry.getWorkflowStateDef(wfType, request.GetWorkflowStateId())
	input := NewObject(request.StateInput, w.options.ObjectEncoder)
	reqContext := request.GetContext()
	wfCtx := newWorkflowContext(
		ctx, reqContext.GetWorkflowId(), reqContext.GetWorkflowRunId(), reqContext.GetStateExecutionId(), reqContext.GetWorkflowStartedTimestamp(),
		int(reqContext.GetAttempt()), reqContext.GetFirstAttemptTimestamp())

	commandResults, err := fromIdlCommandResults(request.CommandResults, w.options.ObjectEncoder)
	if err != nil {
		return nil, err
	}
	pers, err := newPersistence(w.options.ObjectEncoder, w.registry.getWorkflowDataObjectKeyStore(wfType), w.registry.getSearchAttributeTypeStore(wfType), request.DataObjects, request.SearchAttributes, request.GetStateLocals())
	if err != nil {
		return nil, err
	}
	comm := newCommunication(w.options.ObjectEncoder, w.registry.getWorkflowInterStateChannelNameStore(wfType))
	decision, err := stateDef.State.Decide(wfCtx, input, commandResults, pers, comm)
	if err != nil {
		return nil, err
	}
	idlDecision, err := toIdlDecision(decision, wfType, w.registry, w.options.ObjectEncoder)
	resp = &iwfidl.WorkflowStateDecideResponse{
		StateDecision: idlDecision,
	}
	publishings := toPublishing(comm.GetToPublishInterStateChannel())
	if len(publishings) > 0 {
		resp.PublishToInterStateChannel = publishings
	}
	dataObjectsToReturn, stateLocalToReturn, recordedEvents, upsertSearchAttributes := pers.GetToReturn()
	if len(dataObjectsToReturn) > 0 {
		resp.UpsertDataObjects = dataObjectsToReturn
	}
	if len(stateLocalToReturn) > 0 {
		resp.UpsertStateLocals = stateLocalToReturn
	}
	if len(recordedEvents) > 0 {
		resp.RecordEvents = recordedEvents
	}
	if len(upsertSearchAttributes) > 0 {
		resp.UpsertSearchAttributes = upsertSearchAttributes
	}
	return resp, nil
}
