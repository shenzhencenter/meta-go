package objects

type AdCampaignDeliveryStatsGet struct {
	Data *[]map[string]interface{} `json:"data,omitempty"`
}

var AdCampaignDeliveryStatsGetFields = struct {
	Data string
}{
	Data: "data",
}
