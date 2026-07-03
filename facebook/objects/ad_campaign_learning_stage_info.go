package objects

type AdCampaignLearningStageInfo struct {
	AttributionWindows            *[]string               `json:"attribution_windows,omitempty"`
	CasSegment                    *string                 `json:"cas_segment,omitempty"`
	Conversions                   *uint64                 `json:"conversions,omitempty"`
	CurrentBudgetPrediction       *map[string]interface{} `json:"current_budget_prediction,omitempty"`
	DynamicLpConversionsThreshold *uint64                 `json:"dynamic_lp_conversions_threshold,omitempty"`
	DynamicLpDaysThreshold        *uint64                 `json:"dynamic_lp_days_threshold,omitempty"`
	DynamicLpStatus               *string                 `json:"dynamic_lp_status,omitempty"`
	IsOneToOneCboBudget           *bool                   `json:"is_one_to_one_cbo_budget,omitempty"`
	LastSigEditTs                 *int                    `json:"last_sig_edit_ts,omitempty"`
	RecommendedBudgetPrediction   *map[string]interface{} `json:"recommended_budget_prediction,omitempty"`
	SigeditRemoval                *bool                   `json:"sigedit_removal,omitempty"`
	Status                        *string                 `json:"status,omitempty"`
}

var AdCampaignLearningStageInfoFields = struct {
	AttributionWindows            string
	CasSegment                    string
	Conversions                   string
	CurrentBudgetPrediction       string
	DynamicLpConversionsThreshold string
	DynamicLpDaysThreshold        string
	DynamicLpStatus               string
	IsOneToOneCboBudget           string
	LastSigEditTs                 string
	RecommendedBudgetPrediction   string
	SigeditRemoval                string
	Status                        string
}{
	AttributionWindows:            "attribution_windows",
	CasSegment:                    "cas_segment",
	Conversions:                   "conversions",
	CurrentBudgetPrediction:       "current_budget_prediction",
	DynamicLpConversionsThreshold: "dynamic_lp_conversions_threshold",
	DynamicLpDaysThreshold:        "dynamic_lp_days_threshold",
	DynamicLpStatus:               "dynamic_lp_status",
	IsOneToOneCboBudget:           "is_one_to_one_cbo_budget",
	LastSigEditTs:                 "last_sig_edit_ts",
	RecommendedBudgetPrediction:   "recommended_budget_prediction",
	SigeditRemoval:                "sigedit_removal",
	Status:                        "status",
}
