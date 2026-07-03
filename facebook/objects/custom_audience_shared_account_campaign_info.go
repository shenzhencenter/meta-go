package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type CustomAudienceSharedAccountCampaignInfo struct {
	AccountID              *core.ID                  `json:"account_id,omitempty"`
	AccountName            *string                   `json:"account_name,omitempty"`
	AdsetExcludingCount    *uint64                   `json:"adset_excluding_count,omitempty"`
	AdsetIncludingCount    *uint64                   `json:"adset_including_count,omitempty"`
	CampaignDeliveryStatus *string                   `json:"campaign_delivery_status,omitempty"`
	CampaignObjective      *string                   `json:"campaign_objective,omitempty"`
	CampaignPages          *[]map[string]interface{} `json:"campaign_pages,omitempty"`
	CampaignSchedule       *string                   `json:"campaign_schedule,omitempty"`
}
