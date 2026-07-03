package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var AdsCustomPivotsPreviewFields = struct {
	AccountID       string
	AccountName     string
	AdID            string
	AdName          string
	AdsetID         string
	AdsetName       string
	CampaignID      string
	CampaignName    string
	CustomBreakdown string
}{
	AccountID:       "account_id",
	AccountName:     "account_name",
	AdID:            "ad_id",
	AdName:          "ad_name",
	AdsetID:         "adset_id",
	AdsetName:       "adset_name",
	CampaignID:      "campaign_id",
	CampaignName:    "campaign_name",
	CustomBreakdown: "custom_breakdown",
}
