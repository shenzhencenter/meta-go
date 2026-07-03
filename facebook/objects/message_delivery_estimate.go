package objects

type MessageDeliveryEstimate struct {
	EstimateCost               *float64 `json:"estimate_cost,omitempty"`
	EstimateCostLowerBound     *float64 `json:"estimate_cost_lower_bound,omitempty"`
	EstimateCostUpperBound     *float64 `json:"estimate_cost_upper_bound,omitempty"`
	EstimateCoverageLowerBound *int     `json:"estimate_coverage_lower_bound,omitempty"`
	EstimateCoverageUpperBound *int     `json:"estimate_coverage_upper_bound,omitempty"`
	EstimateDelivery           *int     `json:"estimate_delivery,omitempty"`
	EstimateDeliveryLowerBound *int     `json:"estimate_delivery_lower_bound,omitempty"`
	EstimateDeliveryUpperBound *int     `json:"estimate_delivery_upper_bound,omitempty"`
	EstimateStatus             *string  `json:"estimate_status,omitempty"`
}

var MessageDeliveryEstimateFields = struct {
	EstimateCost               string
	EstimateCostLowerBound     string
	EstimateCostUpperBound     string
	EstimateCoverageLowerBound string
	EstimateCoverageUpperBound string
	EstimateDelivery           string
	EstimateDeliveryLowerBound string
	EstimateDeliveryUpperBound string
	EstimateStatus             string
}{
	EstimateCost:               "estimate_cost",
	EstimateCostLowerBound:     "estimate_cost_lower_bound",
	EstimateCostUpperBound:     "estimate_cost_upper_bound",
	EstimateCoverageLowerBound: "estimate_coverage_lower_bound",
	EstimateCoverageUpperBound: "estimate_coverage_upper_bound",
	EstimateDelivery:           "estimate_delivery",
	EstimateDeliveryLowerBound: "estimate_delivery_lower_bound",
	EstimateDeliveryUpperBound: "estimate_delivery_upper_bound",
	EstimateStatus:             "estimate_status",
}
