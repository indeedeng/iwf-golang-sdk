package iwf

import "time"

type Command struct {
	commandId   string
	commandType CommandType
	// for CommandTypeSignalChannel and CommandTypeInterStateChannel
	channelName string
	// for CommandTypeTimer
	firingUnixTimestampSeconds int64
}

type CommandType string

const (
	CommandTypeSignalChannel     CommandType = "SignalChannel"
	CommandTypeTimer             CommandType = "Timer"
	CommandTypeInterStateChannel CommandType = "InterStateChannel"
)

func NewSignalCommand(commandId, channelName string) Command {
	return Command{
		commandId:   commandId,
		commandType: CommandTypeSignalChannel,
		channelName: channelName,
	}
}

func NewInterStateChannelCommand(commandId, channelName string) Command {
	return Command{
		commandId:   commandId,
		commandType: CommandTypeInterStateChannel,
		channelName: channelName,
	}
}

func NewTimerCommand(commandId string, firingTime time.Time) Command {
	return Command{
		commandId:                  commandId,
		commandType:                CommandTypeTimer,
		firingUnixTimestampSeconds: firingTime.Unix(),
	}
}

