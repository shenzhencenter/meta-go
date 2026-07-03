package objects

type ProductItemSubscriptionInfo struct {
	IsSubscribable            *bool   `json:"is_subscribable,omitempty"`
	SubscriptionBillingPeriod *uint64 `json:"subscription_billing_period,omitempty"`
	SubscriptionBillingType   *string `json:"subscription_billing_type,omitempty"`
}

var ProductItemSubscriptionInfoFields = struct {
	IsSubscribable            string
	SubscriptionBillingPeriod string
	SubscriptionBillingType   string
}{
	IsSubscribable:            "is_subscribable",
	SubscriptionBillingPeriod: "subscription_billing_period",
	SubscriptionBillingType:   "subscription_billing_type",
}
