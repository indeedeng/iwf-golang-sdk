package subscription

import "github.com/longquanzheng/iwf/gosdk/iwf"

type WelcomeEmailState struct {
}

const WF_STATE_SEND_WELCOME_EMAIL = "sendWelcomeEmail"

func (w WelcomeEmailState) GetStateId() string {
	return WF_STATE_SEND_WELCOME_EMAIL
}

func (w WelcomeEmailState) GetInputType() iwf.NewTypePtr {
	return func() interface{} {
		return &Customer{}
	}
}

func (w WelcomeEmailState) GetStateOptions() iwf.StateOptions {
	return nil
}

func (w WelcomeEmailState) Execute(ctx iwf.WorkflowContext, customer interface{}, searchAttributes iwf.SearchAttributesRO, queryAttributes iwf.QueryAttributesRO) (iwf.CommandRequest, error) {
	return iwf.RequestAllCommandsCompleted(
		iwf.NewActivityCommand(SEND_WELCOME_EMAIL_ACTIVITY, customer),
	),nil
}

func (w WelcomeEmailState) Decide(ctx iwf.WorkflowContext, customer interface{}, commandResults iwf.CommandResults, searchAttributes iwf.SearchAttributesRW, queryAttributes iwf.QueryAttributesRW) (iwf.StateDecision, error) {
	queryAttributes.Upsert(QUERY_ATTRIBUTE_BILLING_PERIOD_NUMBER, 0) // starting from 0
	queryAttributes.Upsert(QUERY_ATTRIBUTE_CUSTOMER, customer)
	return iwf.NewStateDecision(
		iwf.NewStateMovement(WF_STATE_CANCEL_SUBSCRIPTION),
		iwf.NewStateMovement(WF_STATE_UPDATE_CHARGE_AMOUNT),
		iwf.NewStateMovement(WF_STATE_WAIT_FOR_NEXT_PERIOD),
	),nil
}

var _ iwf.WorkflowState = (*WelcomeEmailState)(nil)