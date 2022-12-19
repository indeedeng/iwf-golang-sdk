package iwf

type CommandResults interface {
	GetAllSignalCommandResults() []SignalCommandResult
	GetSignalValueByIndex(idx int) interface{}
	GetSignalValueById(id string) interface{}
	GetSignalCommandResultByIndex(idx int) SignalCommandResult
	GetSignalCommandResultById(id string) SignalCommandResult
	GetAllTimerCommandResults() []TimerCommandResult
}
