package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type WoodhengePurchasedPAYGReceipt struct {
	ID                             *core.ID   `json:"id,omitempty"`
	NumberOfSubscriptionsPurchased *int       `json:"number_of_subscriptions_purchased,omitempty"`
	PurchaseTime                   *core.Time `json:"purchase_time,omitempty"`
	User                           *User      `json:"user,omitempty"`
}

var WoodhengePurchasedPAYGReceiptFields = struct {
	ID                             string
	NumberOfSubscriptionsPurchased string
	PurchaseTime                   string
	User                           string
}{
	ID:                             "id",
	NumberOfSubscriptionsPurchased: "number_of_subscriptions_purchased",
	PurchaseTime:                   "purchase_time",
	User:                           "user",
}
