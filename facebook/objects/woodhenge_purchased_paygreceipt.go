package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type WoodhengePurchasedPAYGReceipt struct {
	ID                             *core.ID   `json:"id,omitempty"`
	NumberOfSubscriptionsPurchased *int       `json:"number_of_subscriptions_purchased,omitempty"`
	PurchaseTime                   *time.Time `json:"purchase_time,omitempty"`
	User                           *User      `json:"user,omitempty"`
}
