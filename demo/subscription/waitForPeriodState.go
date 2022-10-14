package subscription

import (
	"fmt"
	"time"

	"github.com/longquanzheng/iwf/gosdk/iwf"
)

type WaitForPeriodState struct {

}
const WF_STATE_WAIT_FOR_NEXT_PERIOD = "waitForNextPeriod"

func (w WaitForPeriodState) GetStateId() string {
	return WF_STATE_WAIT_FOR_NEXT_PERIOD
}

func (w WaitForPeriodState) GetInputType() iwf.NewTypePtr {
	return nil
}

func (w WaitForPeriodState) GetStateOptions() iwf.StateOptions {
	return nil
}

func (w WaitForPeriodState) Execute(ctx iwf.WorkflowContext, input interface{}, searchAttributes iwf.SearchAttributesRO, queryAttributes iwf.QueryAttributesRO) (iwf.CommandRequest, error){
	customer, ok := input.(*Customer)
	if !ok{
		return nil, fmt.Errorf("cannot get Customer from input")
	}
	waitTime := time.Duration(customer.Subscription.BillingPeriodInSeconds) * time.Second
	return iwf.RequestAllCommandsCompleted(
		iwf.NewTimerCommand(time.Now().Add(waitTime)),
	), nil
}

func (w WaitForPeriodState) Decide(ctx iwf.WorkflowContext, input interface{}, commandResults iwf.CommandResults, searchAttributes iwf.SearchAttributesRW, queryAttributes iwf.QueryAttributesRW) (iwf.StateDecision, error) {
	customer, ok := input.(*Customer)
	if !ok{
		return nil, fmt.Errorf("cannot get Customer from input")
	}
	currentPeriodNum, ok := queryAttributes.Get(QUERY_ATTRIBUTE_BILLING_PERIOD_NUMBER).(int)
	if !ok{
		return nil, fmt.Errorf("cannot get QUERY_ATTRIBUTE_BILLING_PERIOD_NUMBER from queryAttributes")
	}

	var nextStates []iwf.StateMovement
	if currentPeriodNum < customer.Subscription.PeriodsInSubscription {
		queryAttributes.Upsert(QUERY_ATTRIBUTE_BILLING_PERIOD_NUMBER, currentPeriodNum+1)
		nextStates = append( nextStates, iwf.NewStateMovement(WF_STATE_CHARGE_CURRENT_PERIOD))
	} else {
		nextStates = append( nextStates, iwf.NewStateMovement(WF_STATE_SUBSCRIPTION_OVER))
	}
	return iwf.NewStateDecision(nextStates...), nil
}

var _ iwf.WorkflowState = (*WaitForPeriodState)(nil)