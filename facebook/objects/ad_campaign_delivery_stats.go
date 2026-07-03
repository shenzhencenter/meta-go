package objects

type AdCampaignDeliveryStats struct {
	BidRecommendation     *int                                                    `json:"bid_recommendation,omitempty"`
	CurrentAverageCost    *float64                                                `json:"current_average_cost,omitempty"`
	LastSignificantEditTs *int                                                    `json:"last_significant_edit_ts,omitempty"`
	LearningStageExitInfo *map[string]interface{}                                 `json:"learning_stage_exit_info,omitempty"`
	LearningStageInfo     *AdCampaignLearningStageInfo                            `json:"learning_stage_info,omitempty"`
	UnsupportedFeatures   *[]map[string]AdCampaignDeliveryStatsUnsupportedReasons `json:"unsupported_features,omitempty"`
}
