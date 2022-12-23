package iwf

import (
	"github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"
	"strings"
)

func fromIdlCommandResults(results *iwfidl.CommandResults, encoder ObjectEncoder) (CommandResults, error) {
	if results == nil {
		return CommandResults{}, nil
	}
	var timerResults []TimerCommandResult
	var signalResults []SignalCommandResult
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
	// TODO interstatechannel
	return CommandResults{
		Timers:  timerResults,
		Signals: signalResults,
	}, nil
}

func toIdlCommandRequest(commandRequest *CommandRequest) (*iwfidl.CommandRequest, error) {
	// TODO interstateChannel commands
	var timerCmds []iwfidl.TimerCommand
	var signalCmds []iwfidl.SignalCommand
	for _, t := range commandRequest.Commands {
		if t.CommandType == CommandTypeTimer {
			timerCmd := iwfidl.TimerCommand{
				CommandId:                  t.CommandId,
				FiringUnixTimestampSeconds: t.TimerCommand.FiringUnixTimestampSeconds,
			}
			timerCmds = append(timerCmds, timerCmd)
		}
		if t.CommandType == CommandTypeSignalChannel {
			signalCmd := iwfidl.SignalCommand{
				CommandId:         t.CommandId,
				SignalChannelName: t.SignalCommand.ChannelName,
			}
			signalCmds = append(signalCmds, signalCmd)
		}
	}

	idlCmdReq := &iwfidl.CommandRequest{
		DeciderTriggerType: commandRequest.DeciderTriggerType,
	}
	if len(timerCmds) > 0 {
		idlCmdReq.TimerCommands = timerCmds
	}
	if len(signalCmds) > 0 {
		idlCmdReq.SignalCommands = signalCmds
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
			options = stateDef.State.GetStateOptions()
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
