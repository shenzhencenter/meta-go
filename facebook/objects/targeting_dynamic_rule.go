package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type TargetingDynamicRule struct {
	ActionType       *string  `json:"action.type,omitempty"`
	AdGroupID        *core.ID `json:"ad_group_id,omitempty"`
	CampaignGroupID  *core.ID `json:"campaign_group_id,omitempty"`
	CampaignID       *core.ID `json:"campaign_id,omitempty"`
	ImpressionCount  *string  `json:"impression_count,omitempty"`
	PageID           *core.ID `json:"page_id,omitempty"`
	Post             *string  `json:"post,omitempty"`
	RetentionSeconds *string  `json:"retention_seconds,omitempty"`
}

var TargetingDynamicRuleFields = struct {
	ActionType       string
	AdGroupID        string
	CampaignGroupID  string
	CampaignID       string
	ImpressionCount  string
	PageID           string
	Post             string
	RetentionSeconds string
}{
	ActionType:       "action.type",
	AdGroupID:        "ad_group_id",
	CampaignGroupID:  "campaign_group_id",
	CampaignID:       "campaign_id",
	ImpressionCount:  "impression_count",
	PageID:           "page_id",
	Post:             "post",
	RetentionSeconds: "retention_seconds",
}
