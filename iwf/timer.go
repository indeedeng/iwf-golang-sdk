package iwf

import "time"

type TimerCommand interface {
	BaseCommand
	GetFiringTime() time.Time
}

func NewTimerCommand(firingTime time.Time) TimerCommand{
	return nil
}

type TimerStatus string

const (
	SCHEDULED TimerStatus = "SCHEDULED"
	FIRED     TimerStatus = "FIRED"
)

type TimerCommandResult interface {
	GetCommandId() string
	GetTimerStatus() TimerStatus
}
