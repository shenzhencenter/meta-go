package objects

type AdCampaignDeliveryStatsUnsupportedReasons struct {
	ReasonData *[]map[string]string `json:"reason_data,omitempty"`
	ReasonType *string              `json:"reason_type,omitempty"`
}

var AdCampaignDeliveryStatsUnsupportedReasonsFields = struct {
	ReasonData string
	ReasonType string
}{
	ReasonData: "reason_data",
	ReasonType: "reason_type",
}
