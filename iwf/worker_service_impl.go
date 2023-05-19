package iwf

import (
	"context"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
)

type workerServiceImpl struct {
	registry Registry
	options  WorkerOptions
}

func (w *workerServiceImpl) HandleWorkflowWorkerRPC(ctx context.Context, request iwfidl.WorkflowWorkerRpcRequest) (resp *iwfidl.WorkflowWorkerRpcResponse, retErr error) {
	defer func() { captureStateExecutionError(recover(), &retErr) }()

	wfType := request.GetWorkflowType()
	methodDef := w.registry.getWorkflowRPC(wfType, request.RpcName)
	input := NewObject(request.Input, w.options.ObjectEncoder)
	reqContext := request.GetContext()
	wfCtx := newWorkflowContext(
		ctx, reqContext.GetWorkflowId(), reqContext.GetWorkflowRunId(), reqContext.GetStateExecutionId(), reqContext.GetWorkflowStartedTimestamp(),
		int(reqContext.GetAttempt()), reqContext.GetFirstAttemptTimestamp())

	pers, err := newPersistence(w.options.ObjectEncoder, w.registry.getWorkflowDataAttributesKeyStore(wfType), w.registry.getSearchAttributeTypeStore(wfType), request.DataAttributes, request.SearchAttributes, nil)
	if err != nil {
		return nil, err
	}
	comm := newCommunication(w.options.ObjectEncoder, w.registry.getWorkflowInternalChannelNameStore(wfType))

	output, err := methodDef.RPC(wfCtx, input, pers, comm)
	if err != nil {
		return nil, err
	}

	outputEncoded, err := w.options.ObjectEncoder.Encode(output)
	resp = &iwfidl.WorkflowWorkerRpcResponse{
		Output: outputEncoded,
	}

	publishings := toPublishing(comm.GetToPublishInternalChannel())
	if len(publishings) > 0 {
		resp.PublishToInterStateChannel = publishings
	}

	if len(comm.GetToTriggerStateMovements()) > 0 {
		stateDecision := StateDecision{
			NextStates: comm.GetToTriggerStateMovements(),
		}
		idlDecision, err := toIdlDecision(&stateDecision, wfType, w.registry, w.options.ObjectEncoder)
		if err != nil {
			return nil, err
		}

		resp.StateDecision = &iwfidl.StateDecision{
			NextStates: idlDecision.NextStates,
		}
	}

	dataObjectsToReturn, stateLocalToReturn, recordedEvents, upsertSearchAttributes := pers.GetToReturn()
	if len(dataObjectsToReturn) > 0 {
		resp.UpsertDataAttributes = dataObjectsToReturn
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

func (w *workerServiceImpl) HandleWorkflowStateWaitUntil(ctx context.Context, request iwfidl.WorkflowStateWaitUntilRequest) (resp *iwfidl.WorkflowStateWaitUntilResponse, retErr error) {
	defer func() { captureStateExecutionError(recover(), &retErr) }()

	wfType := request.GetWorkflowType()
	stateDef := w.registry.getWorkflowStateDef(wfType, request.GetWorkflowStateId())
	input := NewObject(request.StateInput, w.options.ObjectEncoder)
	reqContext := request.GetContext()
	wfCtx := newWorkflowContext(
		ctx, reqContext.GetWorkflowId(), reqContext.GetWorkflowRunId(), reqContext.GetStateExecutionId(), reqContext.GetWorkflowStartedTimestamp(),
		int(reqContext.GetAttempt()), reqContext.GetFirstAttemptTimestamp())

	pers, err := newPersistence(w.options.ObjectEncoder, w.registry.getWorkflowDataAttributesKeyStore(wfType), w.registry.getSearchAttributeTypeStore(wfType), request.DataObjects, request.SearchAttributes, nil)
	if err != nil {
		return nil, err
	}
	comm := newCommunication(w.options.ObjectEncoder, w.registry.getWorkflowInternalChannelNameStore(wfType))
	commandRequest, err := stateDef.State.WaitUntil(wfCtx, input, pers, comm)
	if err != nil {
		return nil, err
	}

	err = canNotRequestAndPublishTheSameInterStateChannel(comm.GetToPublishInternalChannel(), commandRequest)
	if err != nil {
		return nil, err
	}

	idlCommandRequest, err := toIdlCommandRequest(commandRequest)
	if err != nil {
		return nil, err
	}
	publishings := toPublishing(comm.GetToPublishInternalChannel())
	resp = &iwfidl.WorkflowStateWaitUntilResponse{
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
			if cr.CommandType == CommandTypeInternalChannel {
				chName := cr.InternalChannelCommand.ChannelName
				_, ok := channelToPublish[chName]
				if ok {
					return NewWorkflowDefinitionErrorFmt("Cannot publish and request for the same interStateChannel: %v", chName)
				}
			}
		}
	}
	return nil
}

func (w *workerServiceImpl) HandleWorkflowStateExecute(ctx context.Context, request iwfidl.WorkflowStateExecuteRequest) (resp *iwfidl.WorkflowStateExecuteResponse, retErr error) {
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
	pers, err := newPersistence(w.options.ObjectEncoder, w.registry.getWorkflowDataAttributesKeyStore(wfType), w.registry.getSearchAttributeTypeStore(wfType), request.DataObjects, request.SearchAttributes, request.GetStateLocals())
	if err != nil {
		return nil, err
	}
	comm := newCommunication(w.options.ObjectEncoder, w.registry.getWorkflowInternalChannelNameStore(wfType))
	decision, err := stateDef.State.Execute(wfCtx, input, commandResults, pers, comm)
	if err != nil {
		return nil, err
	}
	idlDecision, err := toIdlDecision(decision, wfType, w.registry, w.options.ObjectEncoder)
	resp = &iwfidl.WorkflowStateExecuteResponse{
		StateDecision: idlDecision,
	}
	publishings := toPublishing(comm.GetToPublishInternalChannel())
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
