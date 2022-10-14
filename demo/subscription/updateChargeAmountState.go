package subscription

import (
	"fmt"

	"github.com/longquanzheng/iwf/gosdk/iwf"
)

type UpdateChargeAmountState struct {

}

const WF_STATE_UPDATE_CHARGE_AMOUNT = "updateChargeAmount"

func (w UpdateChargeAmountState) GetStateId() string {
	return WF_STATE_UPDATE_CHARGE_AMOUNT
}

func (w UpdateChargeAmountState) GetInputType() iwf.NewTypePtr {
	return nil
}

func (w UpdateChargeAmountState) GetStateOptions() iwf.StateOptions {
	return nil
}

func (w UpdateChargeAmountState) Execute(ctx iwf.WorkflowContext, input interface{}, searchAttributes iwf.SearchAttributesRO, queryAttributes iwf.QueryAttributesRO) (iwf.CommandRequest, error) {
	return iwf.RequestAllCommandsCompleted(
		iwf.NewSignalCommand(SIGNAL_METHOD_UPDATE_BILLING_PERIOD_CHARGE_AMOUNT),
		), nil
}

func (w UpdateChargeAmountState) Decide(ctx iwf.WorkflowContext, input interface{}, commandResults iwf.CommandResults, searchAttributes iwf.SearchAttributesRW, queryAttributes iwf.QueryAttributesRW) (iwf.StateDecision, error) {
	newAmount, ok := commandResults.GetSignalValueByIndex(0).(int)
	if !ok{
		return nil,fmt.Errorf("cannot get newAmount from signal")
	}
	customer, ok := input.(*Customer)
	if !ok{
		return nil, fmt.Errorf("cannot get Customer from input")
	}
	customer.Subscription.BillingPeriodCharge = newAmount
	queryAttributes.Upsert(QUERY_ATTRIBUTE_CUSTOMER, customer)
	return iwf.NewStateDecision(
		iwf.NewStateMovement(WF_STATE_UPDATE_CHARGE_AMOUNT),
		), nil
}

var _ iwf.WorkflowState = (*UpdateChargeAmountState)(nil)