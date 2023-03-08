package iwf

import "github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"

type UnregisteredWorkflowOptions struct {
	WorkflowIdReusePolicy   *iwfidl.WorkflowIDReusePolicy
	WorkflowCronSchedule    *string
	WorkflowRetryPolicy     *iwfidl.WorkflowRetryPolicy
	StartStateOptions       *iwfidl.WorkflowStateOptions
	InitialSearchAttributes []iwfidl.SearchAttribute
}
