package iwf

import "github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"

type UnregisteredWorkflowOptions struct {
	WorkflowIdReusePolicy     *iwfidl.IDReusePolicy
	WorkflowCronSchedule      *string
	WorkflowStartDelaySeconds *int32
	WorkflowRetryPolicy       *iwfidl.WorkflowRetryPolicy
	StartStateOptions         *iwfidl.WorkflowStateOptions
	InitialSearchAttributes   []iwfidl.SearchAttribute
}
