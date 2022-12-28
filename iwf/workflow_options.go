package iwf

import "github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"

type WorkflowOptions struct {
	WorkflowIdReusePolicy   *iwfidl.WorkflowIDReusePolicy
	WorkflowCronSchedule    *string
	WorkflowRetryPolicy     *iwfidl.RetryPolicy
	StartStateOptions       *iwfidl.WorkflowStateOptions
	InitialSearchAttributes []iwfidl.SearchAttribute
}