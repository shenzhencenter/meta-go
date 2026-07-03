package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"time"
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
	CreatedTime                                  *time.Time                         `json:"created_time,omitempty"`
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
	EndTime                                      *time.Time                         `json:"end_time,omitempty"`
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
	StartTime                                    *time.Time                         `json:"start_time,omitempty"`
	Status                                       *enums.AdsetStatus                 `json:"status,omitempty"`
	Targeting                                    *Targeting                         `json:"targeting,omitempty"`
	TargetingOptimizationTypes                   *[]map[string]int                  `json:"targeting_optimization_types,omitempty"`
	TimeBasedAdRotationIDBlocks                  *[][]int                           `json:"time_based_ad_rotation_id_blocks,omitempty"`
	TimeBasedAdRotationIntervals                 *[]uint64                          `json:"time_based_ad_rotation_intervals,omitempty"`
	TrendingTopicsSpec                           *TrendingTopicsSpec                `json:"trending_topics_spec,omitempty"`
	UpdatedTime                                  *time.Time                         `json:"updated_time,omitempty"`
	UseNewAppClick                               *bool                              `json:"use_new_app_click,omitempty"`
	ValueRuleSetID                               *core.ID                           `json:"value_rule_set_id,omitempty"`
	ValueRulesApplied                            *bool                              `json:"value_rules_applied,omitempty"`
}
