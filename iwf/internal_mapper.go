package iwf

import (
	"github.com/indeedeng/iwf-golang-sdk/iwf/ptr"
	"strings"

	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
)

func fromIdlCommandResults(results *iwfidl.CommandResults, encoder ObjectEncoder) (CommandResults, error) {
	if results == nil {
		return CommandResults{}, nil
	}
	var timerResults []TimerCommandResult
	var signalResults []SignalCommandResult
	var internalChannelResults []InternalChannelCommandResult
	for _, t := range results.TimerResults {
		timerResult := TimerCommandResult{
			CommandId: t.CommandId,
			Status:    t.TimerStatus,
		}
		timerResults = append(timerResults, timerResult)
	}
	for _, t := range results.SignalResults {
		signalResult := SignalCommandResult{
			CommandId:   t.CommandId,
			ChannelName: t.SignalChannelName,
			Status:      t.SignalRequestStatus,
			SignalValue: NewObject(t.SignalValue, encoder),
		}
		signalResults = append(signalResults, signalResult)
	}
	for _, t := range results.InterStateChannelResults {
		interStateChannelResult := InternalChannelCommandResult{
			CommandId:   t.CommandId,
			ChannelName: t.ChannelName,
			Status:      t.RequestStatus,
			Value:       NewObject(t.Value, encoder),
		}
		internalChannelResults = append(internalChannelResults, interStateChannelResult)
	}
	return CommandResults{
		Timers:                     timerResults,
		Signals:                    signalResults,
		InternalChannelCommands:    internalChannelResults,
		StateWaitUntilApiSucceeded: results.StateStartApiSucceeded,
	}, nil
}

func toIdlCommandRequest(commandRequest *CommandRequest) (*iwfidl.CommandRequest, error) {
	var timerCmds []iwfidl.TimerCommand
	var signalCmds []iwfidl.SignalCommand
	var interStateCmds []iwfidl.InterStateChannelCommand
	for _, t := range commandRequest.Commands {
		commandId := t.CommandId
		if t.CommandType == CommandTypeTimer {
			timerCmd := iwfidl.TimerCommand{
				CommandId:       &commandId,
				DurationSeconds: t.TimerCommand.DurationSeconds,
			}
			timerCmds = append(timerCmds, timerCmd)
		}
		if t.CommandType == CommandTypeSignalChannel {
			signalCmd := iwfidl.SignalCommand{
				CommandId:         &commandId,
				SignalChannelName: t.SignalCommand.ChannelName,
			}
			signalCmds = append(signalCmds, signalCmd)
		}
		if t.CommandType == CommandTypeInternalChannel {
			interstateChannelCmd := iwfidl.InterStateChannelCommand{
				CommandId:   &commandId,
				ChannelName: t.InternalChannelCommand.ChannelName,
			}
			interStateCmds = append(interStateCmds, interstateChannelCmd)
		}
	}

	idlCmdReq := &iwfidl.CommandRequest{
		CommandWaitingType: commandRequest.CommandWaitingType,
	}
	if len(timerCmds) > 0 {
		idlCmdReq.TimerCommands = timerCmds
	}
	if len(signalCmds) > 0 {
		idlCmdReq.SignalCommands = signalCmds
	}
	if len(interStateCmds) > 0 {
		idlCmdReq.InterStateChannelCommands = interStateCmds
	}
	if len(commandRequest.CommandCombinations) > 0 {
		idlCmdReq.CommandCombinations = commandRequest.CommandCombinations
	}
	return idlCmdReq, nil
}

func toIdlDecision(from *StateDecision, wfType string, registry Registry, encoder ObjectEncoder) (*iwfidl.StateDecision, error) {
	var mvs []iwfidl.StateMovement
	for _, fromMv := range from.NextStates {
		input, err := encoder.Encode(fromMv.NextStateInput)
		if err != nil {
			return nil, err
		}
		var options *iwfidl.WorkflowStateOptions
		if !strings.HasPrefix(fromMv.NextStateId, ReservedStateIdPrefix) {
			stateDef := registry.getWorkflowStateDef(wfType, fromMv.NextStateId)
			options = toIdlStateOptions(ShouldSkipWaitUntilAPI(stateDef.State), stateDef.State.GetStateOptions())
		}
		mv := iwfidl.StateMovement{
			StateId:      fromMv.NextStateId,
			StateInput:   input,
			StateOptions: options,
		}
		mvs = append(mvs, mv)
	}
	return &iwfidl.StateDecision{
		NextStates: mvs,
	}, nil
}

func toIdlStateOptions(skipWaitUntil bool, stateOptions *StateOptions) *iwfidl.WorkflowStateOptions {
	if stateOptions == nil {
		stateOptions = &StateOptions{}
	}

	idlStOptions := &iwfidl.WorkflowStateOptions{
		WaitUntilApiSearchAttributesLoadingPolicy: stateOptions.WaitUntilApiSearchAttributesLoadingPolicy,
		ExecuteApiSearchAttributesLoadingPolicy:   stateOptions.ExecuteApiSearchAttributesLoadingPolicy,
		WaitUntilApiDataAttributesLoadingPolicy:   stateOptions.WaitUntilApiDataAttributesLoadingPolicy,
		ExecuteApiDataAttributesLoadingPolicy:     stateOptions.ExecuteApiDataAttributesLoadingPolicy,
		SearchAttributesLoadingPolicy:             stateOptions.SearchAttributesLoadingPolicy,
		DataAttributesLoadingPolicy:               stateOptions.DataAttributesLoadingPolicy,
		WaitUntilApiTimeoutSeconds:                stateOptions.WaitUntilApiTimeoutSeconds,
		ExecuteApiTimeoutSeconds:                  stateOptions.ExecuteApiTimeoutSeconds,
		WaitUntilApiRetryPolicy:                   stateOptions.WaitUntilApiRetryPolicy,
		ExecuteApiRetryPolicy:                     stateOptions.ExecuteApiRetryPolicy,
		WaitUntilApiFailurePolicy:                 stateOptions.WaitUntilApiFailurePolicy,
	}

	if skipWaitUntil {
		idlStOptions.SkipWaitUntil = ptr.Any(true)
	}

	if stateOptions.ExecuteApiFailureProceedState != nil {
		idlStOptions.ExecuteApiFailurePolicy = iwfidl.PROCEED_TO_CONFIGURED_STATE.Ptr()
		idlStOptions.ExecuteApiFailureProceedStateId = ptr.Any(GetFinalWorkflowStateId(stateOptions.ExecuteApiFailureProceedState))

		proceedStateOptions := stateOptions.ExecuteApiFailureProceedState.GetStateOptions()
		if proceedStateOptions != nil && proceedStateOptions.ExecuteApiFailureProceedState != nil {
			panic("nested failure handling/recovery is not supported: ExecuteApiFailureProceedState cannot have ExecuteApiFailureProceedState")
		}
		idlStOptions.ExecuteApiFailureProceedStateOptions =
			toIdlStateOptions(ShouldSkipWaitUntilAPI(stateOptions.ExecuteApiFailureProceedState), proceedStateOptions)
	}

	return idlStOptions
}
