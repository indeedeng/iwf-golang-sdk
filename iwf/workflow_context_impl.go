package iwf

import "context"

type workflowContextImpl struct {
	context.Context
	workflowId                    string
	workflowRunId                 string
	stateExecutionId              string
	workflowStartTimestampSeconds int64
	attempt                       int
	firstAttemptTimestampSeconds  int64
}

func newWorkflowContext(
	ctx context.Context, workflowId string, workflowRunId string, stateExecutionId string, workflowStartTimestampSeconds int64,
	attempt int, firstAttemptTimestampSeconds int64,
) WorkflowContext {
	return &workflowContextImpl{
		Context:                       ctx,
		workflowId:                    workflowId,
		workflowRunId:                 workflowRunId,
		stateExecutionId:              stateExecutionId,
		workflowStartTimestampSeconds: workflowStartTimestampSeconds,
		attempt:                       attempt,
		firstAttemptTimestampSeconds:  firstAttemptTimestampSeconds,
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

func (w workflowContextImpl) GetFirstAttemptTimestampSeconds() int64 {
	return w.firstAttemptTimestampSeconds
}

func (w workflowContextImpl) GetAttempt() int {
	return w.attempt
}