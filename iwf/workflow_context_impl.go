package iwf

import "context"

type workflowContextImpl struct {
	context.Context
	workflowId                    string
	workflowRunId                 string
	stateExecutionId              string
	workflowStartTimestampSeconds int64
}

func newWorkflowContext(ctx context.Context, workflowId string, workflowRunId string, stateExecutionId string, workflowStartTimestampSeconds int64) WorkflowContext {
	return &workflowContextImpl{
		Context:                       ctx,
		workflowId:                    workflowId,
		workflowRunId:                 workflowRunId,
		stateExecutionId:              stateExecutionId,
		workflowStartTimestampSeconds: workflowStartTimestampSeconds,
	}
}

func (w workflowContextImpl) GetWorkflowId() string {
	return w.workflowId
}

func (w workflowContextImpl) GetWorkflowStartTimestampSeconds() int64 {
	return w.workflowStartTimestampSeconds
}

func (w workflowContextImpl) GetStateExecutionId() string {
	return w.stateExecutionId
}

func (w workflowContextImpl) GetWorkflowRunId() string {
	return w.workflowRunId
}
