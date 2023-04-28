package iwf

import "github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"

type WorkflowStopOptions struct {
	StopType iwfidl.WorkflowStopType
	Reason   string
}
