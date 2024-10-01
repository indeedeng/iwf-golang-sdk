package iwf

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

func NewTimerCommand(commandId string, durationSeconds int64) Command {
	return Command{
		CommandId:   commandId,
		CommandType: CommandTypeTimer,
		TimerCommand: &TimerCommand{
			DurationSeconds: durationSeconds,
		},
	}
}
