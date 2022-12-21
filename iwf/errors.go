package iwf

import (
	"fmt"
	"net/http"
)

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

func NewWorkflowDefinitionFmtError(tpl string, arg ...interface{}) error {
	return &WorkflowDefinitionError{
		msg: fmt.Sprintf(tpl, arg...),
	}
}

type InternalServiceError struct {
	Message      string
	Status       int
	HttpResponse http.Response
}

func NewInternalServiceError(message string, httpResponse http.Response) error {
	return &InternalServiceError{
		Message:      message,
		Status:       httpResponse.StatusCode,
		HttpResponse: httpResponse,
	}
}

func (i InternalServiceError) Error() string {
	return fmt.Sprintf("error message:%v, statusCode: %v", i.Message, i.Status)
}

