package objects

type AdAccountAAASimilarCampaigns struct {
	SimilarCampaignLimit *uint64                              `json:"similar_campaign_limit,omitempty"`
	SimilarCampaignsInfo *[]map[string]map[string]interface{} `json:"similar_campaigns_info,omitempty"`
	UsedCampaignSlots    *uint64                              `json:"used_campaign_slots,omitempty"`
}

var AdAccountAAASimilarCampaignsFields = struct {
	SimilarCampaignLimit string
	SimilarCampaignsInfo string
	UsedCampaignSlots    string
}{
	SimilarCampaignLimit: "similar_campaign_limit",
	SimilarCampaignsInfo: "similar_campaigns_info",
	UsedCampaignSlots:    "used_campaign_slots",
}
