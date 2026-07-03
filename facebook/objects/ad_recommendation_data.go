package objects

type AdRecommendationData struct {
	Link *string `json:"link,omitempty"`
}

var AdRecommendationDataFields = struct {
	Link string
}{
	Link: "link",
}
