package objects

type AdsOptimalDeliveryGrowthOpportunity struct {
	ChildMetadata    *[]map[string]map[string]interface{} `json:"child_metadata,omitempty"`
	Metadata         *map[string]interface{}              `json:"metadata,omitempty"`
	OptimizationType *string                              `json:"optimization_type,omitempty"`
}

var AdsOptimalDeliveryGrowthOpportunityFields = struct {
	ChildMetadata    string
	Metadata         string
	OptimizationType string
}{
	ChildMetadata:    "child_metadata",
	Metadata:         "metadata",
	OptimizationType: "optimization_type",
}
