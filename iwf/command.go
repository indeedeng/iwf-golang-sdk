package iwf

import "time"

type (
	CommandType string

	Command struct {
		CommandId                string
		CommandType              CommandType
		TimerCommand             *TimerCommand
		SignalCommand            *SignalCommand
		InterStateChannelCommand *InterStateChannelCommand
	}

	TimerCommand struct {
		FiringUnixTimestampSeconds int64
	}

	SignalCommand struct {
		ChannelName string
	}

	InterStateChannelCommand struct {
		ChannelName string
	}
)

const (
	CommandTypeSignalChannel     CommandType = "SignalChannel"
	CommandTypeTimer             CommandType = "Timer"
	CommandTypeInterStateChannel CommandType = "InterStateChannel"
)

func NewSignalCommand(commandId, channelName string) Command {
	return Command{
		CommandId:   commandId,
		CommandType: CommandTypeSignalChannel,
		SignalCommand: &SignalCommand{
			ChannelName: channelName,
		},
	}
}

func NewInterStateChannelCommand(commandId, channelName string) Command {
	return Command{
		CommandId:   commandId,
		CommandType: CommandTypeInterStateChannel,
		InterStateChannelCommand: &InterStateChannelCommand{
			ChannelName: channelName,
		},
	}
}

func NewTimerCommand(commandId string, firingTime time.Time) Command {
	return Command{
		CommandId:   commandId,
		CommandType: CommandTypeTimer,
		TimerCommand: &TimerCommand{
			FiringUnixTimestampSeconds: firingTime.Unix(),
		},
	}
}

