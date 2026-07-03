package objects

type AdAccountRecommendations struct {
	Recommendations *[]map[string]interface{} `json:"recommendations,omitempty"`
}

var AdAccountRecommendationsFields = struct {
	Recommendations string
}{
	Recommendations: "recommendations",
}
