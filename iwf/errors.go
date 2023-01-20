package iwf

import (
	"fmt"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"log"
	"net/http"
	"runtime/debug"
)

type InvalidArgumentError struct {
	msg string
}

func (w InvalidArgumentError) Error() string {
	return fmt.Sprintf("WorkflowDefinitionError: %s", w.msg)
}

func NewInvalidArgumentError(msg string) error {
	return &InvalidArgumentError{
		msg: msg,
	}
}

func NewInvalidArgumentErrorFmt(tpl string, arg ...interface{}) error {
	return &WorkflowDefinitionError{
		msg: fmt.Sprintf(tpl, arg...),
	}
}

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

func NewWorkflowDefinitionErrorFmt(tpl string, arg ...interface{}) error {
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

func newInternalError(format string, args ...interface{}) error {
	return &InternalServiceError{
		Message: fmt.Sprintf(format, args...),
	}
}

func (i InternalServiceError) Error() string {
	return fmt.Sprintf("error message:%v, statusCode: %v", i.Message, i.Status)
}

type StateExecutionError struct {
	OriginalError error
	StackTrace    string
}

func newStateExecutionError(err error, stackTrace string) error {
	return &StateExecutionError{
		OriginalError: err,
		StackTrace:    stackTrace,
	}
}

func (i StateExecutionError) Error() string {
	return fmt.Sprintf("error message:%v, stacktrace: %v", i.OriginalError, i.StackTrace)
}

// for skipping the logging in testing code
var skipCaptureErrorLogging = false

// MUST be the result from calling recover, which MUST be done in a single level deep
// deferred function. The usual way of calling this is:
// - defer func() { captureStateExecutionError(recover(), logger, &err) }()
func captureStateExecutionError(errPanic interface{}, retError *error) {
	if errPanic != nil || *retError != nil {
		st := string(debug.Stack())

		var err error
		panicError, ok := errPanic.(error)
		if errPanic != nil {
			if ok && panicError != nil {
				err = newStateExecutionError(panicError, st)
			} else {
				err = newStateExecutionError(fmt.Errorf("errPanic is not an error %v", errPanic), st)
			}
		} else {
			err = newStateExecutionError(*retError, st)
		}

		if !skipCaptureErrorLogging && errPanic != nil {
			log.Printf("panic is captured: %v", errPanic)
		}
		*retError = err
	}
}

// GetOpenApiErrorDetailedMessage retrieve the API error body into a string to be human-readable
func GetOpenApiErrorDetailedMessage(err error) string {
	oerr, ok := err.(*iwfidl.GenericOpenAPIError)
	if !ok {
		return "not an OpenAPI Generic Error type"
	}
	return string(oerr.Body())
}