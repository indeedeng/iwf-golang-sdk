package iwf

import "context"

type WorkflowContext interface {
	context.Context
	GetWorkflowId() string
	GetWorkflowStartTimestampSeconds() int64
	GetStateExecutionId() string
	GetWorkflowRunId() string
	// GetFirstAttemptTimestampSeconds returns the start time of the first attempt of the API call. It's from ScheduledTimestamp of Cadence/Temporal activity.GetInfo
	// require server version 1.2.2+, return 0 if server version is lower
	GetFirstAttemptTimestampSeconds() int64
	// GetAttempt returns an attempt number, which starts from 1, and increased by 1 for every retry if retry policy is specified. It's from Attempt of Cadence/Temporal activity.GetInfo
	// require server version 1.2.2+, return 0 if server version is lower
	GetAttempt() int
}
