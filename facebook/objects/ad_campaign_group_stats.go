package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type AdCampaignGroupStats struct {
	Actions                 *map[string]int `json:"actions,omitempty"`
	CampaignGroupID         *core.ID        `json:"campaign_group_id,omitempty"`
	Clicks                  *uint64         `json:"clicks,omitempty"`
	EndTime                 *time.Time      `json:"end_time,omitempty"`
	Impressions             *uint64         `json:"impressions,omitempty"`
	InlineActions           *map[string]int `json:"inline_actions,omitempty"`
	SocialClicks            *uint64         `json:"social_clicks,omitempty"`
	SocialImpressions       *uint64         `json:"social_impressions,omitempty"`
	SocialSpent             *uint64         `json:"social_spent,omitempty"`
	SocialUniqueClicks      *uint64         `json:"social_unique_clicks,omitempty"`
	SocialUniqueImpressions *uint64         `json:"social_unique_impressions,omitempty"`
	Spent                   *uint64         `json:"spent,omitempty"`
	StartTime               *time.Time      `json:"start_time,omitempty"`
	UniqueClicks            *uint64         `json:"unique_clicks,omitempty"`
	UniqueImpressions       *uint64         `json:"unique_impressions,omitempty"`
}
