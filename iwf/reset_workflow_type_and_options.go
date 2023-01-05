package iwf

import (
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"time"
)

type ResetWorkflowTypeAndOptions struct {
	// ResetType is required
	ResetType iwfidl.WorkflowResetType
	// Reason is required
	Reason string
	// HistoryEventId is required for iwfidl.HISTORY_EVENT_ID resetType
	HistoryEventId *int32
	// HistoryEventTime is required for iwfidl.HISTORY_EVENT_TIME resetType
	HistoryEventTime *time.Time
	//  StateId is required for iwfidl.STATE_ID resetType
	StateId *string
	// StateExecutionId is required for iwfidl.STATE_EXECUTION_ID resetType
	StateExecutionId *string
	// SkipSignalReapply is optional, default to false on server, which will re-apply all signals
	SkipSignalReapply *bool
}

func ResetToBeginning(reason string) ResetWorkflowTypeAndOptions {
	return ResetWorkflowTypeAndOptions{
		ResetType: iwfidl.BEGINNING,
		Reason:    reason,
	}
}

func ResetToHistoryEventId(historyEventId int32, reason string) ResetWorkflowTypeAndOptions {
	return ResetWorkflowTypeAndOptions{
		ResetType:      iwfidl.HISTORY_EVENT_ID,
		Reason:         reason,
		HistoryEventId: &historyEventId,
	}
}

func ResetToHistoryEventTime(historyEventTime time.Time, reason string) ResetWorkflowTypeAndOptions {
	return ResetWorkflowTypeAndOptions{
		ResetType:        iwfidl.HISTORY_EVENT_TIME,
		Reason:           reason,
		HistoryEventTime: &historyEventTime,
	}
}

func ResetToStateId(stateId, reason string) ResetWorkflowTypeAndOptions {
	return ResetWorkflowTypeAndOptions{
		ResetType: iwfidl.STATE_ID,
		Reason:    reason,
		StateId:   &stateId,
	}
}

func ResetToStateExecutionId(stateExecutionId, reason string) ResetWorkflowTypeAndOptions {
	return ResetWorkflowTypeAndOptions{
		ResetType:        iwfidl.STATE_ID,
		Reason:           reason,
		StateExecutionId: &stateExecutionId,
	}
}
