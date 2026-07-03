package objects

type RecommendedIGMedia struct {
	IntentScore *float64 `json:"intent_score,omitempty"`
	Media       *IGMedia `json:"media,omitempty"`
}

var RecommendedIGMediaFields = struct {
	IntentScore string
	Media       string
}{
	IntentScore: "intent_score",
	Media:       "media",
}
