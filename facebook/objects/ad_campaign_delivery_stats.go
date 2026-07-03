package objects

type AdCampaignDeliveryStats struct {
	BidRecommendation     *int                                                    `json:"bid_recommendation,omitempty"`
	CurrentAverageCost    *float64                                                `json:"current_average_cost,omitempty"`
	LastSignificantEditTs *int                                                    `json:"last_significant_edit_ts,omitempty"`
	LearningStageExitInfo *map[string]interface{}                                 `json:"learning_stage_exit_info,omitempty"`
	LearningStageInfo     *AdCampaignLearningStageInfo                            `json:"learning_stage_info,omitempty"`
	UnsupportedFeatures   *[]map[string]AdCampaignDeliveryStatsUnsupportedReasons `json:"unsupported_features,omitempty"`
}

var AdCampaignDeliveryStatsFields = struct {
	BidRecommendation     string
	CurrentAverageCost    string
	LastSignificantEditTs string
	LearningStageExitInfo string
	LearningStageInfo     string
	UnsupportedFeatures   string
}{
	BidRecommendation:     "bid_recommendation",
	CurrentAverageCost:    "current_average_cost",
	LastSignificantEditTs: "last_significant_edit_ts",
	LearningStageExitInfo: "learning_stage_exit_info",
	LearningStageInfo:     "learning_stage_info",
	UnsupportedFeatures:   "unsupported_features",
}
