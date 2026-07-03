package objects

type AdAccountIosFourteenCampaignLimits struct {
	CampaignGroupLimit         *int                      `json:"campaign_group_limit,omitempty"`
	CampaignGroupLimitsDetails *[]map[string]interface{} `json:"campaign_group_limits_details,omitempty"`
	CampaignLimit              *int                      `json:"campaign_limit,omitempty"`
}

var AdAccountIosFourteenCampaignLimitsFields = struct {
	CampaignGroupLimit         string
	CampaignGroupLimitsDetails string
	CampaignLimit              string
}{
	CampaignGroupLimit:         "campaign_group_limit",
	CampaignGroupLimitsDetails: "campaign_group_limits_details",
	CampaignLimit:              "campaign_limit",
}
