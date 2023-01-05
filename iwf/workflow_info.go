package iwf

import "github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"

type WorkflowInfo struct {
	Status       iwfidl.WorkflowStatus
	CurrentRunId string
}
