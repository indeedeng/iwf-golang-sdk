package iwf

import "github.com/golang/protobuf/ptypes/duration"

type ActivityTypeDef interface {
	GetActivityType() string
	GetActivityOutputType() NewTypePtr
}

func NewActivityDef(activityType string, outputType NewTypePtr) ActivityTypeDef{
	return nil
}

type ActivityOptions interface {
	GetActivityCommandId() string
	GetStartToCloseTimeout() duration.Duration
	// TODO: other optional configs: 1. retryOption, 2. tasklist 3, other detailed timeouts(e.g. startToClose, heartbeat)
}

type ActivityStatus string

const (
	OPEN      ActivityStatus = "OPEN"
	COMPLETED ActivityStatus = "COMPLETED"
	FAILED    ActivityStatus = "FAILED"
	TIMEOUT   ActivityStatus = "TIMEOUT"
)

type ActivityCommand interface {
	BaseCommand
	GetActivityType() string
	GetActivityInput() interface{}
	GetActivityOptions() ActivityOptions
}

func NewActivityCommand(activityType string, input ...interface{}) ActivityCommand{
	return nil
}

type ActivityTimeoutType string

const (
	SCHEDULED_TO_START ActivityTimeoutType = "SCHEDULED_TO_START"
	START_TO_CLOSE     ActivityTimeoutType = "START_TO_CLOSE"
	SCHEDULE_TO_CLOSE  ActivityTimeoutType = "SCHEDULE_TO_CLOSE"
	HEARTBEAT          ActivityTimeoutType = "HEARTBEAT"
)

type ActivityCommandResult interface {
	GetActivityType() string
	GetActivityCommandId() string
	GetActivityOutput(output interface{})
	GetActivityStatus() ActivityStatus
	GetActivityTimeoutType() ActivityTimeoutType
}
