package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ALMEvent struct {
	AdAccountIds        *[]core.ID `json:"ad_account_ids,omitempty"`
	CampaignIds         *[]core.ID `json:"campaign_ids,omitempty"`
	Channel             *string    `json:"channel,omitempty"`
	Event               *string    `json:"event,omitempty"`
	EventTime           *core.Time `json:"event_time,omitempty"`
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

var ALMEventFields = struct {
	AdAccountIds        string
	CampaignIds         string
	Channel             string
	Event               string
	EventTime           string
	Guidance            string
	GuidanceDetail      string
	GuidanceType        string
	ID                  string
	ParentAdvertiserIds string
	ResellerBusinessID  string
	SolutionID          string
	SubChannel          string
	UserID              string
}{
	AdAccountIds:        "ad_account_ids",
	CampaignIds:         "campaign_ids",
	Channel:             "channel",
	Event:               "event",
	EventTime:           "event_time",
	Guidance:            "guidance",
	GuidanceDetail:      "guidance_detail",
	GuidanceType:        "guidance_type",
	ID:                  "id",
	ParentAdvertiserIds: "parent_advertiser_ids",
	ResellerBusinessID:  "reseller_business_id",
	SolutionID:          "solution_id",
	SubChannel:          "sub_channel",
	UserID:              "user_id",
}
