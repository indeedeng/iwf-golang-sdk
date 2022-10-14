package iwf

type SignalMethodDef interface {
	GetName() string
	GetValueType() NewTypePtr
}

func NewSignalMethodDef(signalName string, valueType NewTypePtr) SignalMethodDef{
	return nil
}

type SignalCommand interface {
	BaseCommand
	GetName() string
}

func NewSignalCommand(name string) SignalCommand{
	return nil
}

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
