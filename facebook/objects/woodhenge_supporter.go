package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type WoodhengeSupporter struct {
	CreationTime               *time.Time `json:"creation_time,omitempty"`
	ID                         *core.ID   `json:"id,omitempty"`
	IsGiftedSubscription       *bool      `json:"is_gifted_subscription,omitempty"`
	MostRecentSubscriptionTime *time.Time `json:"most_recent_subscription_time,omitempty"`
	NumberOfMonthsSubscribed   *uint64    `json:"number_of_months_subscribed,omitempty"`
	User                       *User      `json:"user,omitempty"`
}
