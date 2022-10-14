package subscription

import (
	"fmt"

	"github.com/longquanzheng/iwf/gosdk/iwf"
)

type SubscriptioinOverState struct {

}

const WF_STATE_SUBSCRIPTION_OVER = "subscriptionOver"

func (w SubscriptioinOverState) GetStateId() string {
	return WF_STATE_SUBSCRIPTION_OVER
}

func (w SubscriptioinOverState) GetInputType() iwf.NewTypePtr {
	return nil
}

func (w SubscriptioinOverState) GetStateOptions() iwf.StateOptions {
	return nil
}

func (w SubscriptioinOverState) Execute(ctx iwf.WorkflowContext, input interface{}, searchAttributes iwf.SearchAttributesRO, queryAttributes iwf.QueryAttributesRO) (iwf.CommandRequest, error) {
	customer, ok := input.(*Customer)
	if !ok{
		return nil, fmt.Errorf("cannot get Customer from input")
	}
	return iwf.RequestAllCommandsCompleted(
		iwf.NewActivityCommand(SEND_SUBSCRIPTION_OVER_EMAIL_ACTIVITY, customer),
	), nil
}

func (w SubscriptioinOverState) Decide(ctx iwf.WorkflowContext, input interface{}, commandResults iwf.CommandResults, searchAttributes iwf.SearchAttributesRW, queryAttributes iwf.QueryAttributesRW) (iwf.StateDecision, error) {
	return iwf.CompletingWorkflow(), nil
}

var _ iwf.WorkflowState = (*SubscriptioinOverState)(nil)