package iwf

import "github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"

type CommandRequest struct {
	Commands            []Command
	CommandCombinations []iwfidl.CommandCombination
	DeciderTriggerType  iwfidl.DeciderTriggerType
}

// EmptyCommandRequest will jump to decide stage immediately.
func EmptyCommandRequest() *CommandRequest {
	return &CommandRequest{
		// it doesn't matter what type is for empty commands
		// TODO: server doesn't seem requiring it, maybe we can remove it
		DeciderTriggerType: iwfidl.ALL_COMMAND_COMPLETED,
	}
}

// AnyCommandCompletedRequest will wait for all the commands to complete
func AnyCommandCompletedRequest(commands ...Command) *CommandRequest {
	return &CommandRequest{
		Commands:           commands,
		DeciderTriggerType: iwfidl.ANY_COMMAND_COMPLETED,
	}
}

// AllCommandsCompletedRequest will wait for any the commands to complete
func AllCommandsCompletedRequest(commands ...Command) *CommandRequest {
	return &CommandRequest{
		Commands:           commands,
		DeciderTriggerType: iwfidl.ALL_COMMAND_COMPLETED,
	}
}

// AnyCommandCombinationsCompletedRequest will wait for any combination to complete.
// Using this requires every command has a commandId when created.
// Functionally this one can cover both forAllCommandCompleted, forAnyCommandCompleted. So the other two are like "shortcuts" of it.
func AnyCommandCombinationsCompletedRequest(listsOfCommandIds [][]string, commands ...Command) *CommandRequest {
	for _, cmd := range commands {
		if cmd.CommandId == "" {
			panic("commandId must be provided for using ANY_COMMAND_COMBINATION_COMPLETED")
		}
	}
	
	var comList []iwfidl.CommandCombination
	for _, commandIds := range listsOfCommandIds {
		com := iwfidl.CommandCombination{
			CommandIds: commandIds,
		}
		comList = append(comList, com)
	}
	return &CommandRequest{
		Commands:            commands,
		CommandCombinations: comList,
		DeciderTriggerType:  iwfidl.ANY_COMMAND_COMBINATION_COMPLETED,
	}
}
