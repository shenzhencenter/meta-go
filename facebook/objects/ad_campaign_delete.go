package objects

type AdCampaignDelete struct {
	Success *bool `json:"success,omitempty"`
}

var AdCampaignDeleteFields = struct {
	Success string
}{
	Success: "success",
}
