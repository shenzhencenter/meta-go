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
