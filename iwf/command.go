package iwf

import "time"

type (
	CommandType string

	Command struct {
		CommandId              string
		CommandType            CommandType
		TimerCommand           *TimerCommand
		SignalCommand          *SignalCommand
		InternalChannelCommand *InternalChannelCommand
	}

	TimerCommand struct {
		DurationSeconds int64
	}

	SignalCommand struct {
		ChannelName string
	}

	InternalChannelCommand struct {
		ChannelName string
	}
)

const (
	CommandTypeSignalChannel   CommandType = "SignalChannel"
	CommandTypeTimer           CommandType = "Timer"
	CommandTypeInternalChannel CommandType = "InternalChannel"
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

func NewInternalChannelCommand(commandId, channelName string) Command {
	return Command{
		CommandId:   commandId,
		CommandType: CommandTypeInternalChannel,
		InternalChannelCommand: &InternalChannelCommand{
			ChannelName: channelName,
		},
	}
}

// Deprecated: Use NewTimerCommandByDuration instead.
func NewTimerCommand(commandId string, firingTime time.Time) Command {
	durationSeconds := int64(time.Until(firingTime).Seconds())
	if durationSeconds < 0 {
		panic("Firing time is set in the past")
	}

	return Command{
		CommandId:   commandId,
		CommandType: CommandTypeTimer,
		TimerCommand: &TimerCommand{
			DurationSeconds: durationSeconds,
		},
	}
}

func NewTimerCommandByDuration(commandId string, durationSeconds time.Duration) Command {
	return Command{
		CommandId:   commandId,
		CommandType: CommandTypeTimer,
		TimerCommand: &TimerCommand{
			DurationSeconds: int64(durationSeconds.Seconds()),
		},
	}
}
