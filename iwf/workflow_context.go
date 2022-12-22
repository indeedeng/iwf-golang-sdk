package iwf

import "context"

type WorkflowContext interface {
	context.Context
	GetWorkflowId() string
	GetWorkflowStartTimestampSeconds() int64
	GetStateExecutionId() string
	GetWorkflowRunId() string
}