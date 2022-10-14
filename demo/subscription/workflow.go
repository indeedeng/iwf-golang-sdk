package subscription

import "github.com/longquanzheng/iwf/gosdk/iwf"

const SIGNAL_METHOD_CANCEL_SUBSCRIPTION = "CancelSubscription"
const SIGNAL_METHOD_UPDATE_BILLING_PERIOD_CHARGE_AMOUNT = "UpdateBillingPeriodChargeAmount"

const QUERY_ATTRIBUTE_BILLING_PERIOD_NUMBER = "BillingPeriodNumber"
const QUERY_ATTRIBUTE_CUSTOMER = "BillingSubscription"

const SEND_WELCOME_EMAIL_ACTIVITY = "SubscriptionActivities::sendWelcomeEmail"
const SEND_SUBSCRIPTION_OVER_EMAIL_ACTIVITY = "SubscriptionActivities::sendSubscriptionOverEmail"
const CHARGE_CUSTOMER_ACTIVITY = "SubscriptionActivities::chargeCustomerForBillingPeriod"

type SubscriptionWorkflow struct {
}

func (s SubscriptionWorkflow) GetStates() []iwf.StateDef {
	return []iwf.StateDef{
		iwf.NewStateDef( &WelcomeEmailState{}, true),
		iwf.NewStateDef( &CancelSubscriptionState{}, false),
		iwf.NewStateDef( &WaitForPeriodState{}, false),
		iwf.NewStateDef( &ChargeCurrentPeriodState{}, false),
		iwf.NewStateDef( &SubscriptioinOverState{}, false),
		iwf.NewStateDef( &UpdateChargeAmountState{}, false),
	}
}

func (s SubscriptionWorkflow) GetActivityTypes() []iwf.ActivityTypeDef {
	return []iwf.ActivityTypeDef{
		iwf.NewActivityDef(SEND_WELCOME_EMAIL_ACTIVITY, nil),
		iwf.NewActivityDef(SEND_SUBSCRIPTION_OVER_EMAIL_ACTIVITY, nil),
		iwf.NewActivityDef(CHARGE_CUSTOMER_ACTIVITY, nil),
	}
}

func (s SubscriptionWorkflow) GetSignalMethods() []iwf.SignalMethodDef {
	return []iwf.SignalMethodDef{
		iwf.NewSignalMethodDef(SIGNAL_METHOD_CANCEL_SUBSCRIPTION, nil),
		iwf.NewSignalMethodDef(SIGNAL_METHOD_UPDATE_BILLING_PERIOD_CHARGE_AMOUNT, func() interface{} {
			var i int
			return &i
		}),
	}
}

func (s SubscriptionWorkflow) GetSearchAttributes() []iwf.SearchAttributeDef {
	return nil
}

func (s SubscriptionWorkflow) GetQueryAttributes() []iwf.QueryAttributeDef {
	return []iwf.QueryAttributeDef{
		iwf.NewQueryAttributeDef(QUERY_ATTRIBUTE_BILLING_PERIOD_NUMBER, func() interface{} {
			var i int
			return &i
		}),
		iwf.NewQueryAttributeDef(QUERY_ATTRIBUTE_CUSTOMER, func() interface{} {
			return &Customer{}
		}),
	}
}

var _ iwf.Workflow = (*SubscriptionWorkflow)(nil)
