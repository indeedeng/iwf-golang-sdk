package subscription

import (
	"fmt"

	"github.com/longquanzheng/iwf/gosdk/iwf"
)

type ChargeCurrentPeriodState struct {

}

const WF_STATE_CHARGE_CURRENT_PERIOD = "chargeCurrentPeriod"

func (w ChargeCurrentPeriodState) GetStateId() string {
	return WF_STATE_CHARGE_CURRENT_PERIOD
}

func (w ChargeCurrentPeriodState) GetInputType() iwf.NewTypePtr {
	return nil
}

func (w ChargeCurrentPeriodState) GetStateOptions() iwf.StateOptions {
	return nil
}

func (w ChargeCurrentPeriodState) Execute(ctx iwf.WorkflowContext, input interface{}, searchAttributes iwf.SearchAttributesRO, queryAttributes iwf.QueryAttributesRO) (iwf.CommandRequest, error) {
	customer, ok := input.(*Customer)
	if !ok{
		return nil, fmt.Errorf("cannot get Customer from input")
	}
	currentPeriodNum, ok := queryAttributes.Get(QUERY_ATTRIBUTE_BILLING_PERIOD_NUMBER).(int)
	if !ok{
		return nil, fmt.Errorf("cannot get QUERY_ATTRIBUTE_BILLING_PERIOD_NUMBER from queryAttributes")
	}
	return iwf.RequestAllCommandsCompleted(
		iwf.NewActivityCommand(CHARGE_CUSTOMER_ACTIVITY, customer, currentPeriodNum),
		), nil
}

func (w ChargeCurrentPeriodState) Decide(ctx iwf.WorkflowContext, input interface{}, commandResults iwf.CommandResults, searchAttributes iwf.SearchAttributesRW, queryAttributes iwf.QueryAttributesRW) (iwf.StateDecision, error) {
	return iwf.NewStateDecision(
		iwf.NewStateMovement(WF_STATE_WAIT_FOR_NEXT_PERIOD),
		), nil
}

var _ iwf.WorkflowState = (*ChargeCurrentPeriodState)(nil)