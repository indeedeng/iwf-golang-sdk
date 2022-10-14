package iwf

import "context"

type WorkflowContext interface {
	context.Context
    GetWorkflowId() string
}
