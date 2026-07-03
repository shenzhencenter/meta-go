package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AdSet struct {
	AccountID                                    *core.ID                           `json:"account_id,omitempty"`
	Adlabels                                     *[]AdLabel                         `json:"adlabels,omitempty"`
	AdsetSchedule                                *[]DayPart                         `json:"adset_schedule,omitempty"`
	AssetFeedID                                  *core.ID                           `json:"asset_feed_id,omitempty"`
	AttributionCountType                         *string                            `json:"attribution_count_type,omitempty"`
	AttributionSpec                              *[]AttributionSpec                 `json:"attribution_spec,omitempty"`
	AutomaticManualState                         *string                            `json:"automatic_manual_state,omitempty"`
	BidAdjustments                               *AdBidAdjustments                  `json:"bid_adjustments,omitempty"`
	BidAmount                                    *uint64                            `json:"bid_amount,omitempty"`
	BidConstraints                               *AdCampaignBidConstraint           `json:"bid_constraints,omitempty"`
	BidInfo                                      *map[string]uint64                 `json:"bid_info,omitempty"`
	BidStrategy                                  *enums.AdsetBidStrategy            `json:"bid_strategy,omitempty"`
	BillingEvent                                 *enums.AdsetBillingEvent           `json:"billing_event,omitempty"`
	BrandSafetyConfig                            *BrandSafetyCampaignConfig         `json:"brand_safety_config,omitempty"`
	BudgetRemaining                              *string                            `json:"budget_remaining,omitempty"`
	Campaign                                     *Campaign                          `json:"campaign,omitempty"`
	CampaignActiveTime                           *string                            `json:"campaign_active_time,omitempty"`
	CampaignAttribution                          *string                            `json:"campaign_attribution,omitempty"`
	CampaignID                                   *core.ID                           `json:"campaign_id,omitempty"`
	ConfiguredStatus                             *enums.AdsetConfiguredStatus       `json:"configured_status,omitempty"`
	CostBiddingMode                              *string                            `json:"cost_bidding_mode,omitempty"`
	CreatedTime                                  *core.Time                         `json:"created_time,omitempty"`
	CreativeDiversityLabel                       *[]string                          `json:"creative_diversity_label,omitempty"`
	CreativeDiversityScore                       *[]string                          `json:"creative_diversity_score,omitempty"`
	CreativeSequence                             *[]string                          `json:"creative_sequence,omitempty"`
	CreativeSequenceRepetitionPattern            *string                            `json:"creative_sequence_repetition_pattern,omitempty"`
	DailyBudget                                  *string                            `json:"daily_budget,omitempty"`
	DailyMinSpendTarget                          *string                            `json:"daily_min_spend_target,omitempty"`
	DailySpendCap                                *string                            `json:"daily_spend_cap,omitempty"`
	DestinationType                              *string                            `json:"destination_type,omitempty"`
	DsaBeneficiary                               *string                            `json:"dsa_beneficiary,omitempty"`
	DsaPayor                                     *string                            `json:"dsa_payor,omitempty"`
	EffectiveStatus                              *enums.AdsetEffectiveStatus        `json:"effective_status,omitempty"`
	EndTime                                      *core.Time                         `json:"end_time,omitempty"`
	ExistingCustomerBudgetPercentage             *uint64                            `json:"existing_customer_budget_percentage,omitempty"`
	FrequencyControlSpecs                        *[]AdCampaignFrequencyControlSpecs `json:"frequency_control_specs,omitempty"`
	FullFunnelExplorationMode                    *string                            `json:"full_funnel_exploration_mode,omitempty"`
	ID                                           *core.ID                           `json:"id,omitempty"`
	InstagramUserID                              *core.ID                           `json:"instagram_user_id,omitempty"`
	IsBaSkipDelayedEligible                      *bool                              `json:"is_ba_skip_delayed_eligible,omitempty"`
	IsBudgetScheduleEnabled                      *bool                              `json:"is_budget_schedule_enabled,omitempty"`
	IsDcFollowOptimized                          *bool                              `json:"is_dc_follow_optimized,omitempty"`
	IsDynamicCreative                            *bool                              `json:"is_dynamic_creative,omitempty"`
	IsIncrementalAttributionEnabled              *bool                              `json:"is_incremental_attribution_enabled,omitempty"`
	IsOrganicAdJointOptimized                    *bool                              `json:"is_organic_ad_joint_optimized,omitempty"`
	IssuesInfo                                   *[]AdCampaignIssuesInfo            `json:"issues_info,omitempty"`
	LearningStageInfo                            *AdCampaignLearningStageInfo       `json:"learning_stage_info,omitempty"`
	LifetimeBudget                               *string                            `json:"lifetime_budget,omitempty"`
	LifetimeImps                                 *int                               `json:"lifetime_imps,omitempty"`
	LifetimeMinSpendTarget                       *string                            `json:"lifetime_min_spend_target,omitempty"`
	LifetimeSpendCap                             *string                            `json:"lifetime_spend_cap,omitempty"`
	LiveVideoAdCampaignConfig                    *LiveVideoAdCampaignConfig         `json:"live_video_ad_campaign_config,omitempty"`
	LowCreativeReach                             *[]string                          `json:"low_creative_reach,omitempty"`
	MaxBudgetSpendPercentage                     *string                            `json:"max_budget_spend_percentage,omitempty"`
	MetaMomentMakerSpec                          *MetaMomentMakerConfig             `json:"meta_moment_maker_spec,omitempty"`
	MinBudgetSpendPercentage                     *string                            `json:"min_budget_spend_percentage,omitempty"`
	MultiEventConversionAttributionWindowSeconds *int                               `json:"multi_event_conversion_attribution_window_seconds,omitempty"`
	MultiOptimizationGoalWeight                  *string                            `json:"multi_optimization_goal_weight,omitempty"`
	Name                                         *string                            `json:"name,omitempty"`
	OptimizationGoal                             *enums.AdsetOptimizationGoal       `json:"optimization_goal,omitempty"`
	OptimizationSubEvent                         *string                            `json:"optimization_sub_event,omitempty"`
	PacingType                                   *[]string                          `json:"pacing_type,omitempty"`
	PlacementSoftOptOut                          *PlacementSoftOptOut               `json:"placement_soft_opt_out,omitempty"`
	PromotedObject                               *AdPromotedObject                  `json:"promoted_object,omitempty"`
	Recommendations                              *[]AdRecommendation                `json:"recommendations,omitempty"`
	RecurringBudgetSemantics                     *bool                              `json:"recurring_budget_semantics,omitempty"`
	RegionalRegulatedCategories                  *[]string                          `json:"regional_regulated_categories,omitempty"`
	RegionalRegulationIdentities                 *RegionalRegulationIdentities      `json:"regional_regulation_identities,omitempty"`
	RelativeValue                                *string                            `json:"relative_value,omitempty"`
	ReviewFeedback                               *string                            `json:"review_feedback,omitempty"`
	RfPredictionID                               *core.ID                           `json:"rf_prediction_id,omitempty"`
	SourceAdset                                  *AdSet                             `json:"source_adset,omitempty"`
	SourceAdsetID                                *core.ID                           `json:"source_adset_id,omitempty"`
	SpecialAdCategories                          *[]string                          `json:"special_ad_categories,omitempty"`
	StartTime                                    *core.Time                         `json:"start_time,omitempty"`
	Status                                       *enums.AdsetStatus                 `json:"status,omitempty"`
	Targeting                                    *Targeting                         `json:"targeting,omitempty"`
	TargetingOptimizationTypes                   *[]map[string]int                  `json:"targeting_optimization_types,omitempty"`
	TimeBasedAdRotationIDBlocks                  *[][]int                           `json:"time_based_ad_rotation_id_blocks,omitempty"`
	TimeBasedAdRotationIntervals                 *[]uint64                          `json:"time_based_ad_rotation_intervals,omitempty"`
	TrendingTopicsSpec                           *TrendingTopicsSpec                `json:"trending_topics_spec,omitempty"`
	UpdatedTime                                  *core.Time                         `json:"updated_time,omitempty"`
	UseNewAppClick                               *bool                              `json:"use_new_app_click,omitempty"`
	ValueRuleSetID                               *core.ID                           `json:"value_rule_set_id,omitempty"`
	ValueRulesApplied                            *bool                              `json:"value_rules_applied,omitempty"`
}

var AdSetFields = struct {
	AccountID                                    string
	Adlabels                                     string
	AdsetSchedule                                string
	AssetFeedID                                  string
	AttributionCountType                         string
	AttributionSpec                              string
	AutomaticManualState                         string
	BidAdjustments                               string
	BidAmount                                    string
	BidConstraints                               string
	BidInfo                                      string
	BidStrategy                                  string
	BillingEvent                                 string
	BrandSafetyConfig                            string
	BudgetRemaining                              string
	Campaign                                     string
	CampaignActiveTime                           string
	CampaignAttribution                          string
	CampaignID                                   string
	ConfiguredStatus                             string
	CostBiddingMode                              string
	CreatedTime                                  string
	CreativeDiversityLabel                       string
	CreativeDiversityScore                       string
	CreativeSequence                             string
	CreativeSequenceRepetitionPattern            string
	DailyBudget                                  string
	DailyMinSpendTarget                          string
	DailySpendCap                                string
	DestinationType                              string
	DsaBeneficiary                               string
	DsaPayor                                     string
	EffectiveStatus                              string
	EndTime                                      string
	ExistingCustomerBudgetPercentage             string
	FrequencyControlSpecs                        string
	FullFunnelExplorationMode                    string
	ID                                           string
	InstagramUserID                              string
	IsBaSkipDelayedEligible                      string
	IsBudgetScheduleEnabled                      string
	IsDcFollowOptimized                          string
	IsDynamicCreative                            string
	IsIncrementalAttributionEnabled              string
	IsOrganicAdJointOptimized                    string
	IssuesInfo                                   string
	LearningStageInfo                            string
	LifetimeBudget                               string
	LifetimeImps                                 string
	LifetimeMinSpendTarget                       string
	LifetimeSpendCap                             string
	LiveVideoAdCampaignConfig                    string
	LowCreativeReach                             string
	MaxBudgetSpendPercentage                     string
	MetaMomentMakerSpec                          string
	MinBudgetSpendPercentage                     string
	MultiEventConversionAttributionWindowSeconds string
	MultiOptimizationGoalWeight                  string
	Name                                         string
	OptimizationGoal                             string
	OptimizationSubEvent                         string
	PacingType                                   string
	PlacementSoftOptOut                          string
	PromotedObject                               string
	Recommendations                              string
	RecurringBudgetSemantics                     string
	RegionalRegulatedCategories                  string
	RegionalRegulationIdentities                 string
	RelativeValue                                string
	ReviewFeedback                               string
	RfPredictionID                               string
	SourceAdset                                  string
	SourceAdsetID                                string
	SpecialAdCategories                          string
	StartTime                                    string
	Status                                       string
	Targeting                                    string
	TargetingOptimizationTypes                   string
	TimeBasedAdRotationIDBlocks                  string
	TimeBasedAdRotationIntervals                 string
	TrendingTopicsSpec                           string
	UpdatedTime                                  string
	UseNewAppClick                               string
	ValueRuleSetID                               string
	ValueRulesApplied                            string
}{
	AccountID:                         "account_id",
	Adlabels:                          "adlabels",
	AdsetSchedule:                     "adset_schedule",
	AssetFeedID:                       "asset_feed_id",
	AttributionCountType:              "attribution_count_type",
	AttributionSpec:                   "attribution_spec",
	AutomaticManualState:              "automatic_manual_state",
	BidAdjustments:                    "bid_adjustments",
	BidAmount:                         "bid_amount",
	BidConstraints:                    "bid_constraints",
	BidInfo:                           "bid_info",
	BidStrategy:                       "bid_strategy",
	BillingEvent:                      "billing_event",
	BrandSafetyConfig:                 "brand_safety_config",
	BudgetRemaining:                   "budget_remaining",
	Campaign:                          "campaign",
	CampaignActiveTime:                "campaign_active_time",
	CampaignAttribution:               "campaign_attribution",
	CampaignID:                        "campaign_id",
	ConfiguredStatus:                  "configured_status",
	CostBiddingMode:                   "cost_bidding_mode",
	CreatedTime:                       "created_time",
	CreativeDiversityLabel:            "creative_diversity_label",
	CreativeDiversityScore:            "creative_diversity_score",
	CreativeSequence:                  "creative_sequence",
	CreativeSequenceRepetitionPattern: "creative_sequence_repetition_pattern",
	DailyBudget:                       "daily_budget",
	DailyMinSpendTarget:               "daily_min_spend_target",
	DailySpendCap:                     "daily_spend_cap",
	DestinationType:                   "destination_type",
	DsaBeneficiary:                    "dsa_beneficiary",
	DsaPayor:                          "dsa_payor",
	EffectiveStatus:                   "effective_status",
	EndTime:                           "end_time",
	ExistingCustomerBudgetPercentage:  "existing_customer_budget_percentage",
	FrequencyControlSpecs:             "frequency_control_specs",
	FullFunnelExplorationMode:         "full_funnel_exploration_mode",
	ID:                                "id",
	InstagramUserID:                   "instagram_user_id",
	IsBaSkipDelayedEligible:           "is_ba_skip_delayed_eligible",
	IsBudgetScheduleEnabled:           "is_budget_schedule_enabled",
	IsDcFollowOptimized:               "is_dc_follow_optimized",
	IsDynamicCreative:                 "is_dynamic_creative",
	IsIncrementalAttributionEnabled:   "is_incremental_attribution_enabled",
	IsOrganicAdJointOptimized:         "is_organic_ad_joint_optimized",
	IssuesInfo:                        "issues_info",
	LearningStageInfo:                 "learning_stage_info",
	LifetimeBudget:                    "lifetime_budget",
	LifetimeImps:                      "lifetime_imps",
	LifetimeMinSpendTarget:            "lifetime_min_spend_target",
	LifetimeSpendCap:                  "lifetime_spend_cap",
	LiveVideoAdCampaignConfig:         "live_video_ad_campaign_config",
	LowCreativeReach:                  "low_creative_reach",
	MaxBudgetSpendPercentage:          "max_budget_spend_percentage",
	MetaMomentMakerSpec:               "meta_moment_maker_spec",
	MinBudgetSpendPercentage:          "min_budget_spend_percentage",
	MultiEventConversionAttributionWindowSeconds: "multi_event_conversion_attribution_window_seconds",
	MultiOptimizationGoalWeight:                  "multi_optimization_goal_weight",
	Name:                                         "name",
	OptimizationGoal:                             "optimization_goal",
	OptimizationSubEvent:                         "optimization_sub_event",
	PacingType:                                   "pacing_type",
	PlacementSoftOptOut:                          "placement_soft_opt_out",
	PromotedObject:                               "promoted_object",
	Recommendations:                              "recommendations",
	RecurringBudgetSemantics:                     "recurring_budget_semantics",
	RegionalRegulatedCategories:                  "regional_regulated_categories",
	RegionalRegulationIdentities:                 "regional_regulation_identities",
	RelativeValue:                                "relative_value",
	ReviewFeedback:                               "review_feedback",
	RfPredictionID:                               "rf_prediction_id",
	SourceAdset:                                  "source_adset",
	SourceAdsetID:                                "source_adset_id",
	SpecialAdCategories:                          "special_ad_categories",
	StartTime:                                    "start_time",
	Status:                                       "status",
	Targeting:                                    "targeting",
	TargetingOptimizationTypes:                   "targeting_optimization_types",
	TimeBasedAdRotationIDBlocks:                  "time_based_ad_rotation_id_blocks",
	TimeBasedAdRotationIntervals:                 "time_based_ad_rotation_intervals",
	TrendingTopicsSpec:                           "trending_topics_spec",
	UpdatedTime:                                  "updated_time",
	UseNewAppClick:                               "use_new_app_click",
	ValueRuleSetID:                               "value_rule_set_id",
	ValueRulesApplied:                            "value_rules_applied",
}
