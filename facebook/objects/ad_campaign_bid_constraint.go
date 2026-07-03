package objects

type AdCampaignBidConstraint struct {
	RoasAverageFloor *int `json:"roas_average_floor,omitempty"`
}

var AdCampaignBidConstraintFields = struct {
	RoasAverageFloor string
}{
	RoasAverageFloor: "roas_average_floor",
}
