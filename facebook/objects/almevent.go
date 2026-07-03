package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type ALMEvent struct {
	AdAccountIds        *[]core.ID `json:"ad_account_ids,omitempty"`
	CampaignIds         *[]core.ID `json:"campaign_ids,omitempty"`
	Channel             *string    `json:"channel,omitempty"`
	Event               *string    `json:"event,omitempty"`
	EventTime           *time.Time `json:"event_time,omitempty"`
	Guidance            *string    `json:"guidance,omitempty"`
	GuidanceDetail      *string    `json:"guidance_detail,omitempty"`
	GuidanceType        *string    `json:"guidance_type,omitempty"`
	ID                  *core.ID   `json:"id,omitempty"`
	ParentAdvertiserIds *[]core.ID `json:"parent_advertiser_ids,omitempty"`
	ResellerBusinessID  *core.ID   `json:"reseller_business_id,omitempty"`
	SolutionID          *core.ID   `json:"solution_id,omitempty"`
	SubChannel          *string    `json:"sub_channel,omitempty"`
	UserID              *core.ID   `json:"user_id,omitempty"`
}
