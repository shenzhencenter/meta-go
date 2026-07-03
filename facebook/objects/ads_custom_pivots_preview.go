package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdsCustomPivotsPreview struct {
	AccountID       *core.ID  `json:"account_id,omitempty"`
	AccountName     *string   `json:"account_name,omitempty"`
	AdID            *core.ID  `json:"ad_id,omitempty"`
	AdName          *string   `json:"ad_name,omitempty"`
	AdsetID         *core.ID  `json:"adset_id,omitempty"`
	AdsetName       *string   `json:"adset_name,omitempty"`
	CampaignID      *core.ID  `json:"campaign_id,omitempty"`
	CampaignName    *string   `json:"campaign_name,omitempty"`
	CustomBreakdown *[]string `json:"custom_breakdown,omitempty"`
}
