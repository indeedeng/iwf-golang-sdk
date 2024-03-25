package iwf

import "github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"

type WorkflowOptions struct {
	WorkflowIdReusePolicy     *iwfidl.IDReusePolicy
	WorkflowCronSchedule      *string
	WorkflowStartDelaySeconds *int32
	WorkflowRetryPolicy       *iwfidl.WorkflowRetryPolicy
	// InitialSearchAttributes set the initial search attributes to start a workflow
	// key is search attribute key, value much match with PersistenceSchema of the workflow definition
	// For iwfidl.DATETIME , the value can be either time.Time or a string value in format of DateTimeFormat
	InitialSearchAttributes map[string]interface{}
}
