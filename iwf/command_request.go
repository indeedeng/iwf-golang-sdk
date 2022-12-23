package iwf

import "github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"

type CommandRequest struct {
	Commands           []Command
	DeciderTriggerType iwfidl.DeciderTriggerType
}

func EmptyCommandRequest() *CommandRequest {
	return &CommandRequest{
		// this field is required to be not empty, but it doesn't matter what type is for empty commands
		DeciderTriggerType: iwfidl.ALL_COMMAND_COMPLETED,
	}
}

func AnyCommandCompletedRequest(commands ...Command) *CommandRequest {
	return &CommandRequest{
		Commands:           commands,
		DeciderTriggerType: iwfidl.ANY_COMMAND_COMPLETED,
	}
}

func AllCommandsCompletedRequest(commands ...Command) *CommandRequest {
	return &CommandRequest{
		Commands:           commands,
		DeciderTriggerType: iwfidl.ALL_COMMAND_COMPLETED,
	}
}
