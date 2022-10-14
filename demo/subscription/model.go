package subscription

type Customer struct {
	Id           string
	Subscription Subscription
}

type Subscription struct {
	PeriodsInSubscription int
	BillingPeriodInSeconds int
	BillingPeriodCharge    int
}
