package objects

type AdAccountRecommendedCamapaignBudget struct {
	Daily     *string `json:"daily,omitempty"`
	Lifetime  *string `json:"lifetime,omitempty"`
	Objective *string `json:"objective,omitempty"`
}

var AdAccountRecommendedCamapaignBudgetFields = struct {
	Daily     string
	Lifetime  string
	Objective string
}{
	Daily:     "daily",
	Lifetime:  "lifetime",
	Objective: "objective",
}
