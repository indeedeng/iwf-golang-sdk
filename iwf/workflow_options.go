package iwf

import "github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"

type WorkflowOptions struct {
	// WorkflowTimeoutSeconds is the only required field. It must be greater than zero
	WorkflowTimeoutSeconds int
	WorkflowIdReusePolicy  *iwfidl.WorkflowIDReusePolicy
	WorkflowCronSchedule   *string
	WorkflowRetryPolicy    *iwfidl.RetryPolicy
	StartStateOptions      *iwfidl.WorkflowStateOptions
}

func Minimum(timeoutSecs int) *WorkflowOptions {
	return &WorkflowOptions{
		WorkflowTimeoutSeconds: timeoutSecs,
	}
}
