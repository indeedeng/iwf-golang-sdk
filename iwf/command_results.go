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

func (c CommandResults) GetTimerCommandResultById(id string) *TimerCommandResult {
	for _, cmd := range c.Timers {
		if cmd.CommandId == id {
			return &cmd
		}
	}
	return nil
}

func (c CommandResults) GetSignalCommandResultById(id string) *SignalCommandResult {
	for _, cmd := range c.Signals {
		if cmd.CommandId == id {
			return &cmd
		}
	}
	return nil
}

func (c CommandResults) GetInterStateChannelCommandResultById(id string) *InterStateChannelCommandResult {
	for _, cmd := range c.InterStateChannelCommands {
		if cmd.CommandId == id {
			return &cmd
		}
	}
	return nil
}

func (c CommandResults) GetSignalCommandResultByChannel(channelName string) *SignalCommandResult {
	for _, cmd := range c.Signals {
		if cmd.ChannelName == channelName {
			return &cmd
		}
	}
	return nil
}

func (c CommandResults) GetInterStateChannelCommandResultByChannel(channelName string) *InterStateChannelCommandResult {
	for _, cmd := range c.InterStateChannelCommands {
		if cmd.ChannelName == channelName {
			return &cmd
		}
	}
	return nil
}