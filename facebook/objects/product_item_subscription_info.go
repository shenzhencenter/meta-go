package objects

type ProductItemSubscriptionInfo struct {
	IsSubscribable            *bool   `json:"is_subscribable,omitempty"`
	SubscriptionBillingPeriod *uint64 `json:"subscription_billing_period,omitempty"`
	SubscriptionBillingType   *string `json:"subscription_billing_type,omitempty"`
}
