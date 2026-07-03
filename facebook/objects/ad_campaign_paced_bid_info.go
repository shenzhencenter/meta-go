package objects

type AdCampaignPacedBidInfo struct {
	BiddingStatus *string `json:"bidding_status,omitempty"`
}

var AdCampaignPacedBidInfoFields = struct {
	BiddingStatus string
}{
	BiddingStatus: "bidding_status",
}
