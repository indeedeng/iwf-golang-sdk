package subscription

import "github.com/longquanzheng/iwf/gosdk/iwf"

type CancelSubscriptionState struct {}

const WF_STATE_CANCEL_SUBSCRIPTION = "cancelSubscription"

func (w CancelSubscriptionState) GetStateId() string {
	return WF_STATE_CANCEL_SUBSCRIPTION
}

func (w CancelSubscriptionState) GetInputType() iwf.NewTypePtr {
	return nil
}

func (w CancelSubscriptionState) GetStateOptions() iwf.StateOptions {
	return nil
}

func (w CancelSubscriptionState) Execute(ctx iwf.WorkflowContext, input interface{}, searchAttributes iwf.SearchAttributesRO, queryAttributes iwf.QueryAttributesRO) (iwf.CommandRequest, error) {
	return iwf.RequestAllCommandsCompleted(
		iwf.NewSignalCommand(SIGNAL_METHOD_CANCEL_SUBSCRIPTION),
	),nil
}

func (w CancelSubscriptionState) Decide(ctx iwf.WorkflowContext, input interface{}, commandResults iwf.CommandResults, searchAttributes iwf.SearchAttributesRW, queryAttributes iwf.QueryAttributesRW) (iwf.StateDecision, error) {
	return iwf.CompletingWorkflow(), nil
}

var _ iwf.WorkflowState = (*CancelSubscriptionState)(nil)