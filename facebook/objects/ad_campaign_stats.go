package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdCampaignStats struct {
	AccountID               *core.ID                `json:"account_id,omitempty"`
	Actions                 *map[string]interface{} `json:"actions,omitempty"`
	AdgroupID               *core.ID                `json:"adgroup_id,omitempty"`
	CampaignID              *core.ID                `json:"campaign_id,omitempty"`
	CampaignIds             *[]core.ID              `json:"campaign_ids,omitempty"`
	Clicks                  *uint64                 `json:"clicks,omitempty"`
	EndTime                 *map[string]interface{} `json:"end_time,omitempty"`
	ID                      *core.ID                `json:"id,omitempty"`
	Impressions             *string                 `json:"impressions,omitempty"`
	InlineActions           *map[string]interface{} `json:"inline_actions,omitempty"`
	IoNumber                *uint64                 `json:"io_number,omitempty"`
	IsCompleted             *bool                   `json:"is_completed,omitempty"`
	LineNumber              *uint64                 `json:"line_number,omitempty"`
	NewsfeedPosition        *map[string]interface{} `json:"newsfeed_position,omitempty"`
	SocialClicks            *uint64                 `json:"social_clicks,omitempty"`
	SocialImpressions       *string                 `json:"social_impressions,omitempty"`
	SocialSpent             *uint64                 `json:"social_spent,omitempty"`
	SocialUniqueClicks      *uint64                 `json:"social_unique_clicks,omitempty"`
	SocialUniqueImpressions *string                 `json:"social_unique_impressions,omitempty"`
	Spent                   *int                    `json:"spent,omitempty"`
	StartTime               *map[string]interface{} `json:"start_time,omitempty"`
	ToplineID               *core.ID                `json:"topline_id,omitempty"`
	UniqueClicks            *uint64                 `json:"unique_clicks,omitempty"`
	UniqueImpressions       *string                 `json:"unique_impressions,omitempty"`
}
