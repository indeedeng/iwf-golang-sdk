package iwf

import "github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"

type CommandRequest struct {
	Commands           []Command
	DeciderTriggerType iwfidl.DeciderTriggerType
}
