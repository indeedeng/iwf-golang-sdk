package iwf

import "github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"

type (
	CommandResults struct {
		Timers                  []TimerCommandResult
		Signals                 []SignalCommandResult
		InternalChannelCommands []InternalChannelCommandResult
		WaitUntilApiSucceeded   *bool
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

	InternalChannelCommandResult struct {
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

func (c CommandResults) GetInternalChannelCommandResultById(id string) *InternalChannelCommandResult {
	for _, cmd := range c.InternalChannelCommands {
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

func (c CommandResults) GetInternalChannelCommandResultByChannel(channelName string) *InternalChannelCommandResult {
	for _, cmd := range c.InternalChannelCommands {
		if cmd.ChannelName == channelName {
			return &cmd
		}
	}
	return nil
}

func (c CommandResults) GetWaitUntilApiSucceeded() *bool {
	return c.WaitUntilApiSucceeded
}
