package objects

type CreativeFatiguePredictionPLE struct {
	CprLiftEstimation *float64 `json:"cpr_lift_estimation,omitempty"`
	ShouldDisplay     *bool    `json:"should_display,omitempty"`
}

var CreativeFatiguePredictionPLEFields = struct {
	CprLiftEstimation string
	ShouldDisplay     string
}{
	CprLiftEstimation: "cpr_lift_estimation",
	ShouldDisplay:     "should_display",
}
