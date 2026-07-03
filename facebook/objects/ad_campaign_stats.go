package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var AdCampaignStatsFields = struct {
	AccountID               string
	Actions                 string
	AdgroupID               string
	CampaignID              string
	CampaignIds             string
	Clicks                  string
	EndTime                 string
	ID                      string
	Impressions             string
	InlineActions           string
	IoNumber                string
	IsCompleted             string
	LineNumber              string
	NewsfeedPosition        string
	SocialClicks            string
	SocialImpressions       string
	SocialSpent             string
	SocialUniqueClicks      string
	SocialUniqueImpressions string
	Spent                   string
	StartTime               string
	ToplineID               string
	UniqueClicks            string
	UniqueImpressions       string
}{
	AccountID:               "account_id",
	Actions:                 "actions",
	AdgroupID:               "adgroup_id",
	CampaignID:              "campaign_id",
	CampaignIds:             "campaign_ids",
	Clicks:                  "clicks",
	EndTime:                 "end_time",
	ID:                      "id",
	Impressions:             "impressions",
	InlineActions:           "inline_actions",
	IoNumber:                "io_number",
	IsCompleted:             "is_completed",
	LineNumber:              "line_number",
	NewsfeedPosition:        "newsfeed_position",
	SocialClicks:            "social_clicks",
	SocialImpressions:       "social_impressions",
	SocialSpent:             "social_spent",
	SocialUniqueClicks:      "social_unique_clicks",
	SocialUniqueImpressions: "social_unique_impressions",
	Spent:                   "spent",
	StartTime:               "start_time",
	ToplineID:               "topline_id",
	UniqueClicks:            "unique_clicks",
	UniqueImpressions:       "unique_impressions",
}
