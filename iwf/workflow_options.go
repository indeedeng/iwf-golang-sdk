package iwf

import "github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"

type WorkflowOptions struct {
	WorkflowIdReusePolicy *iwfidl.WorkflowIDReusePolicy
	WorkflowCronSchedule  *string
	WorkflowRetryPolicy   *iwfidl.RetryPolicy
	StartStateOptions     *iwfidl.WorkflowStateOptions
	// InitialSearchAttributes set the initial search attributes to start a workflow
	// key is search attribute key, value much match with PersistenceSchema of the workflow definition
	InitialSearchAttributes map[string]interface{}
}