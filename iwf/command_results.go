package iwf

import "github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"

type (
	CommandResults struct {
		Timers                    []TimerCommandResult
		Signals                   []SignalCommandResult
		InterStateChannelCommands []InterStateChannelCommandResult
	}

	SignalCommandResult struct {
		CommandId   string
		ChannelName string
		SignalValue Object
		Status      iwfidl.ChannelRequestStatus
	}

	TimerCommandResult struct {
		CommandId string
		Status    iwfidl.TimerStatus
	}

	InterStateChannelCommandResult struct {
		CommandId   string
		ChannelName string
		Value       Object
		Status      iwfidl.ChannelRequestStatus
	}
)
