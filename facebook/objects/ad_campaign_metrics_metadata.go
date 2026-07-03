package objects

type AdCampaignMetricsMetadata struct {
	BoostedComponentOptimization *[]string                 `json:"boosted_component_optimization,omitempty"`
	CreationFlowTips             *[]string                 `json:"creation_flow_tips,omitempty"`
	DefaultOptedInPlacements     *[]map[string]interface{} `json:"default_opted_in_placements,omitempty"`
	DeliveryGrowthOptimizations  *[]map[string]interface{} `json:"delivery_growth_optimizations,omitempty"`
	DuplicationFlowTips          *[]string                 `json:"duplication_flow_tips,omitempty"`
	EditFlowTips                 *[]string                 `json:"edit_flow_tips,omitempty"`
}
