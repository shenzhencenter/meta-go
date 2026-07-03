package objects

type AdDefaultValues struct {
	CampaignGroup *map[string]interface{} `json:"campaign_group,omitempty"`
}

var AdDefaultValuesFields = struct {
	CampaignGroup string
}{
	CampaignGroup: "campaign_group",
}
