package objects

type AdCampaignGroupDelete struct {
	Success *bool `json:"success,omitempty"`
}

var AdCampaignGroupDeleteFields = struct {
	Success string
}{
	Success: "success",
}
