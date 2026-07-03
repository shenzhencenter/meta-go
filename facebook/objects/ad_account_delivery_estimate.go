package objects

type AdAccountDeliveryEstimate struct {
	DailyOutcomesCurve         *[]OutcomePredictionPoint `json:"daily_outcomes_curve,omitempty"`
	EstimateDau                *int                      `json:"estimate_dau,omitempty"`
	EstimateMauLowerBound      *int                      `json:"estimate_mau_lower_bound,omitempty"`
	EstimateMauUpperBound      *int                      `json:"estimate_mau_upper_bound,omitempty"`
	EstimateReady              *bool                     `json:"estimate_ready,omitempty"`
	TargetingOptimizationTypes *[]map[string]int         `json:"targeting_optimization_types,omitempty"`
}
