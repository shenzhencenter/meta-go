package objects

type CreativeFatiguePredictionPLE struct {
	CprLiftEstimation *float64 `json:"cpr_lift_estimation,omitempty"`
	ShouldDisplay     *bool    `json:"should_display,omitempty"`
}
