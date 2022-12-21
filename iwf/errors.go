package iwf

import "fmt"

type WorkflowDefinitionError struct {
	msg string
}

func (w WorkflowDefinitionError) Error() string {
	return fmt.Sprintf("WorkflowDefinitionError: %s", w.msg)
}

func NewWorkflowDefinitionError(msg string) error {
	return &WorkflowDefinitionError{
		msg: msg,
	}
}
