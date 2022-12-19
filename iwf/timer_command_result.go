package iwf


type TimerStatus string

const (
	SCHEDULED TimerStatus = "SCHEDULED"
	FIRED     TimerStatus = "FIRED"
)

type TimerCommandResult interface {
	GetCommandId() string
	GetTimerStatus() TimerStatus
}
