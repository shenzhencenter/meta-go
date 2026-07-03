package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type WoodhengeSupporter struct {
	CreationTime               *core.Time `json:"creation_time,omitempty"`
	ID                         *core.ID   `json:"id,omitempty"`
	IsGiftedSubscription       *bool      `json:"is_gifted_subscription,omitempty"`
	MostRecentSubscriptionTime *core.Time `json:"most_recent_subscription_time,omitempty"`
	NumberOfMonthsSubscribed   *uint64    `json:"number_of_months_subscribed,omitempty"`
	User                       *User      `json:"user,omitempty"`
}

var WoodhengeSupporterFields = struct {
	CreationTime               string
	ID                         string
	IsGiftedSubscription       string
	MostRecentSubscriptionTime string
	NumberOfMonthsSubscribed   string
	User                       string
}{
	CreationTime:               "creation_time",
	ID:                         "id",
	IsGiftedSubscription:       "is_gifted_subscription",
	MostRecentSubscriptionTime: "most_recent_subscription_time",
	NumberOfMonthsSubscribed:   "number_of_months_subscribed",
	User:                       "user",
}
