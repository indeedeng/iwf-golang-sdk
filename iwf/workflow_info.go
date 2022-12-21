package iwf

import "github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"

type WorkflowInfo struct {
	Status       iwfidl.WorkflowStatus
	CurrentRunId string
}
