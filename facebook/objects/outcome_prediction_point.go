package objects

type OutcomePredictionPoint struct {
	Actions     *float64 `json:"actions,omitempty"`
	Impressions *float64 `json:"impressions,omitempty"`
	Reach       *float64 `json:"reach,omitempty"`
	Spend       *int     `json:"spend,omitempty"`
}
