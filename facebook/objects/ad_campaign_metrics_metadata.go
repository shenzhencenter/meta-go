package objects

type AdCampaignMetricsMetadata struct {
	BoostedComponentOptimization *[]string                 `json:"boosted_component_optimization,omitempty"`
	CreationFlowTips             *[]string                 `json:"creation_flow_tips,omitempty"`
	DefaultOptedInPlacements     *[]map[string]interface{} `json:"default_opted_in_placements,omitempty"`
	DeliveryGrowthOptimizations  *[]map[string]interface{} `json:"delivery_growth_optimizations,omitempty"`
	DuplicationFlowTips          *[]string                 `json:"duplication_flow_tips,omitempty"`
	EditFlowTips                 *[]string                 `json:"edit_flow_tips,omitempty"`
}

var AdCampaignMetricsMetadataFields = struct {
	BoostedComponentOptimization string
	CreationFlowTips             string
	DefaultOptedInPlacements     string
	DeliveryGrowthOptimizations  string
	DuplicationFlowTips          string
	EditFlowTips                 string
}{
	BoostedComponentOptimization: "boosted_component_optimization",
	CreationFlowTips:             "creation_flow_tips",
	DefaultOptedInPlacements:     "default_opted_in_placements",
	DeliveryGrowthOptimizations:  "delivery_growth_optimizations",
	DuplicationFlowTips:          "duplication_flow_tips",
	EditFlowTips:                 "edit_flow_tips",
}
