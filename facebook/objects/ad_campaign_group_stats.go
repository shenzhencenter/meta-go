package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCampaignGroupStats struct {
	Actions                 *map[string]int `json:"actions,omitempty"`
	CampaignGroupID         *core.ID        `json:"campaign_group_id,omitempty"`
	Clicks                  *uint64         `json:"clicks,omitempty"`
	EndTime                 *core.Time      `json:"end_time,omitempty"`
	Impressions             *uint64         `json:"impressions,omitempty"`
	InlineActions           *map[string]int `json:"inline_actions,omitempty"`
	SocialClicks            *uint64         `json:"social_clicks,omitempty"`
	SocialImpressions       *uint64         `json:"social_impressions,omitempty"`
	SocialSpent             *uint64         `json:"social_spent,omitempty"`
	SocialUniqueClicks      *uint64         `json:"social_unique_clicks,omitempty"`
	SocialUniqueImpressions *uint64         `json:"social_unique_impressions,omitempty"`
	Spent                   *uint64         `json:"spent,omitempty"`
	StartTime               *core.Time      `json:"start_time,omitempty"`
	UniqueClicks            *uint64         `json:"unique_clicks,omitempty"`
	UniqueImpressions       *uint64         `json:"unique_impressions,omitempty"`
}

var AdCampaignGroupStatsFields = struct {
	Actions                 string
	CampaignGroupID         string
	Clicks                  string
	EndTime                 string
	Impressions             string
	InlineActions           string
	SocialClicks            string
	SocialImpressions       string
	SocialSpent             string
	SocialUniqueClicks      string
	SocialUniqueImpressions string
	Spent                   string
	StartTime               string
	UniqueClicks            string
	UniqueImpressions       string
}{
	Actions:                 "actions",
	CampaignGroupID:         "campaign_group_id",
	Clicks:                  "clicks",
	EndTime:                 "end_time",
	Impressions:             "impressions",
	InlineActions:           "inline_actions",
	SocialClicks:            "social_clicks",
	SocialImpressions:       "social_impressions",
	SocialSpent:             "social_spent",
	SocialUniqueClicks:      "social_unique_clicks",
	SocialUniqueImpressions: "social_unique_impressions",
	Spent:                   "spent",
	StartTime:               "start_time",
	UniqueClicks:            "unique_clicks",
	UniqueImpressions:       "unique_impressions",
}
