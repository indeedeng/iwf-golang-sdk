package iwf

type CommandRequest struct {
	Commands           []Command
	DeciderTriggerType DeciderTriggerType
}

type DeciderTriggerType string

const (
	// AllCommandCompleted will wait for all commands are completed.
	AllCommandCompleted DeciderTriggerType = "AllCommandCompleted"
	// AnyCommandCompleted will wait for any command to be completed
	AnyCommandCompleted DeciderTriggerType = "AnyCommandCompleted"
)
