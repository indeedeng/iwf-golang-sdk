package iwf

type SignalStatus string

const (
	REQUESTED SignalStatus = "REQUESTED"
	RECEIVED  SignalStatus = "RECEIVED"
)

type SignalCommandResult interface {
	GetCommandId() string
	GetName() string
	GeValue() interface{}
	GetSignalStatus() SignalStatus
}
