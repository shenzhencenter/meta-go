package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ReachFrequencyPrediction struct {
	AccountID                       *core.ID                                                `json:"account_id,omitempty"`
	ActivityStatus                  *ReachFrequencyActivity                                 `json:"activity_status,omitempty"`
	AdFormats                       *[]ReachFrequencyAdFormat                               `json:"ad_formats,omitempty"`
	AuctionEntryOptionIndex         *int                                                    `json:"auction_entry_option_index,omitempty"`
	AudienceSizeLowerBound          *uint64                                                 `json:"audience_size_lower_bound,omitempty"`
	AudienceSizeUpperBound          *uint64                                                 `json:"audience_size_upper_bound,omitempty"`
	BusinessID                      *core.ID                                                `json:"business_id,omitempty"`
	BuyingType                      *string                                                 `json:"buying_type,omitempty"`
	CampaignGroupID                 *core.ID                                                `json:"campaign_group_id,omitempty"`
	CampaignID                      *core.ID                                                `json:"campaign_id,omitempty"`
	CampaignTimeStart               *core.Time                                              `json:"campaign_time_start,omitempty"`
	CampaignTimeStop                *core.Time                                              `json:"campaign_time_stop,omitempty"`
	Currency                        *string                                                 `json:"currency,omitempty"`
	CurveBudgetReach                *ReachFrequencyEstimatesCurve                           `json:"curve_budget_reach,omitempty"`
	CurveReach                      *[]uint64                                               `json:"curve_reach,omitempty"`
	DailyGrpCurve                   *[]float64                                              `json:"daily_grp_curve,omitempty"`
	DailyImpressionCurve            *[]float64                                              `json:"daily_impression_curve,omitempty"`
	DailyImpressionCurveMap         *[]map[string][]float64                                 `json:"daily_impression_curve_map,omitempty"`
	DayPartingSchedule              *[]ReachFrequencyDayPart                                `json:"day_parting_schedule,omitempty"`
	DestinationID                   *core.ID                                                `json:"destination_id,omitempty"`
	EndTime                         *core.Time                                              `json:"end_time,omitempty"`
	ExpirationTime                  *core.Time                                              `json:"expiration_time,omitempty"`
	ExternalBudget                  *int                                                    `json:"external_budget,omitempty"`
	ExternalImpression              *uint64                                                 `json:"external_impression,omitempty"`
	ExternalMaximumBudget           *int                                                    `json:"external_maximum_budget,omitempty"`
	ExternalMaximumImpression       *string                                                 `json:"external_maximum_impression,omitempty"`
	ExternalMaximumReach            *uint64                                                 `json:"external_maximum_reach,omitempty"`
	ExternalMinimumBudget           *int                                                    `json:"external_minimum_budget,omitempty"`
	ExternalMinimumImpression       *uint64                                                 `json:"external_minimum_impression,omitempty"`
	ExternalMinimumReach            *uint64                                                 `json:"external_minimum_reach,omitempty"`
	ExternalReach                   *uint64                                                 `json:"external_reach,omitempty"`
	FeedRatioX0000                  *uint64                                                 `json:"feed_ratio_0000,omitempty"`
	FrequencyCap                    *uint64                                                 `json:"frequency_cap,omitempty"`
	FrequencyDistributionMap        *[]map[string][]float64                                 `json:"frequency_distribution_map,omitempty"`
	FrequencyDistributionMapAgg     *[]map[string][]uint64                                  `json:"frequency_distribution_map_agg,omitempty"`
	GrpAudienceSize                 *float64                                                `json:"grp_audience_size,omitempty"`
	GrpAvgProbabilityMap            *string                                                 `json:"grp_avg_probability_map,omitempty"`
	GrpCountryAudienceSize          *float64                                                `json:"grp_country_audience_size,omitempty"`
	GrpCurve                        *[]float64                                              `json:"grp_curve,omitempty"`
	GrpDmasAudienceSize             *float64                                                `json:"grp_dmas_audience_size,omitempty"`
	GrpFilteringThresholdX00        *uint64                                                 `json:"grp_filtering_threshold_00,omitempty"`
	GrpPoints                       *float64                                                `json:"grp_points,omitempty"`
	GrpRatio                        *float64                                                `json:"grp_ratio,omitempty"`
	GrpReachRatio                   *float64                                                `json:"grp_reach_ratio,omitempty"`
	GrpStatus                       *string                                                 `json:"grp_status,omitempty"`
	HoldoutPercentage               *uint64                                                 `json:"holdout_percentage,omitempty"`
	ID                              *core.ID                                                `json:"id,omitempty"`
	ImpressionCurve                 *[]uint64                                               `json:"impression_curve,omitempty"`
	InstagramDestinationID          *core.ID                                                `json:"instagram_destination_id,omitempty"`
	InstreamPackages                *[]string                                               `json:"instream_packages,omitempty"`
	IntervalFrequencyCap            *uint64                                                 `json:"interval_frequency_cap,omitempty"`
	IntervalFrequencyCapResetPeriod *uint64                                                 `json:"interval_frequency_cap_reset_period,omitempty"`
	IsBalancedFrequency             *bool                                                   `json:"is_balanced_frequency,omitempty"`
	IsBonusMedia                    *uint64                                                 `json:"is_bonus_media,omitempty"`
	IsConversionGoal                *uint64                                                 `json:"is_conversion_goal,omitempty"`
	IsHigherAverageFrequency        *bool                                                   `json:"is_higher_average_frequency,omitempty"`
	IsIo                            *bool                                                   `json:"is_io,omitempty"`
	IsReservedBuying                *uint64                                                 `json:"is_reserved_buying,omitempty"`
	IsTrp                           *bool                                                   `json:"is_trp,omitempty"`
	MetaMomentMakerSpec             *MetaMomentMakerConfig                                  `json:"meta_moment_maker_spec,omitempty"`
	Name                            *string                                                 `json:"name,omitempty"`
	Objective                       *uint64                                                 `json:"objective,omitempty"`
	ObjectiveName                   *string                                                 `json:"objective_name,omitempty"`
	OdaxObjective                   *uint64                                                 `json:"odax_objective,omitempty"`
	OdaxObjectiveName               *string                                                 `json:"odax_objective_name,omitempty"`
	OptimizationGoal                *uint64                                                 `json:"optimization_goal,omitempty"`
	OptimizationGoalName            *string                                                 `json:"optimization_goal_name,omitempty"`
	PausePeriods                    *[]map[string]interface{}                               `json:"pause_periods,omitempty"`
	PercentReachAtTargetFrequency   *int                                                    `json:"percent_reach_at_target_frequency,omitempty"`
	PlacementBreakdown              *ReachFrequencyEstimatesPlacementBreakdown              `json:"placement_breakdown,omitempty"`
	PlacementBreakdownMap           *[]map[string]ReachFrequencyEstimatesPlacementBreakdown `json:"placement_breakdown_map,omitempty"`
	PlanName                        *string                                                 `json:"plan_name,omitempty"`
	PlanType                        *string                                                 `json:"plan_type,omitempty"`
	PredictionMode                  *uint64                                                 `json:"prediction_mode,omitempty"`
	PredictionProgress              *uint64                                                 `json:"prediction_progress,omitempty"`
	ReferenceID                     *core.ID                                                `json:"reference_id,omitempty"`
	ReservationStatus               *uint64                                                 `json:"reservation_status,omitempty"`
	StartTime                       *core.Time                                              `json:"start_time,omitempty"`
	Status                          *uint64                                                 `json:"status,omitempty"`
	StoryEventType                  *uint64                                                 `json:"story_event_type,omitempty"`
	TargetCpm                       *uint64                                                 `json:"target_cpm,omitempty"`
	TargetFrequency                 *uint64                                                 `json:"target_frequency,omitempty"`
	TargetFrequencyResetPeriod      *uint64                                                 `json:"target_frequency_reset_period,omitempty"`
	TargetSpec                      *Targeting                                              `json:"target_spec,omitempty"`
	TimeCreated                     *core.Time                                              `json:"time_created,omitempty"`
	TimeUpdated                     *core.Time                                              `json:"time_updated,omitempty"`
	TimezoneID                      *core.ID                                                `json:"timezone_id,omitempty"`
	TimezoneName                    *string                                                 `json:"timezone_name,omitempty"`
	ToplineID                       *core.ID                                                `json:"topline_id,omitempty"`
	TrendingTopicsSpec              *TrendingTopicsSpec                                     `json:"trending_topics_spec,omitempty"`
	VideoViewLengthConstraint       *uint64                                                 `json:"video_view_length_constraint,omitempty"`
	Viewtag                         *string                                                 `json:"viewtag,omitempty"`
}

var ReachFrequencyPredictionFields = struct {
	AccountID                       string
	ActivityStatus                  string
	AdFormats                       string
	AuctionEntryOptionIndex         string
	AudienceSizeLowerBound          string
	AudienceSizeUpperBound          string
	BusinessID                      string
	BuyingType                      string
	CampaignGroupID                 string
	CampaignID                      string
	CampaignTimeStart               string
	CampaignTimeStop                string
	Currency                        string
	CurveBudgetReach                string
	CurveReach                      string
	DailyGrpCurve                   string
	DailyImpressionCurve            string
	DailyImpressionCurveMap         string
	DayPartingSchedule              string
	DestinationID                   string
	EndTime                         string
	ExpirationTime                  string
	ExternalBudget                  string
	ExternalImpression              string
	ExternalMaximumBudget           string
	ExternalMaximumImpression       string
	ExternalMaximumReach            string
	ExternalMinimumBudget           string
	ExternalMinimumImpression       string
	ExternalMinimumReach            string
	ExternalReach                   string
	FeedRatioX0000                  string
	FrequencyCap                    string
	FrequencyDistributionMap        string
	FrequencyDistributionMapAgg     string
	GrpAudienceSize                 string
	GrpAvgProbabilityMap            string
	GrpCountryAudienceSize          string
	GrpCurve                        string
	GrpDmasAudienceSize             string
	GrpFilteringThresholdX00        string
	GrpPoints                       string
	GrpRatio                        string
	GrpReachRatio                   string
	GrpStatus                       string
	HoldoutPercentage               string
	ID                              string
	ImpressionCurve                 string
	InstagramDestinationID          string
	InstreamPackages                string
	IntervalFrequencyCap            string
	IntervalFrequencyCapResetPeriod string
	IsBalancedFrequency             string
	IsBonusMedia                    string
	IsConversionGoal                string
	IsHigherAverageFrequency        string
	IsIo                            string
	IsReservedBuying                string
	IsTrp                           string
	MetaMomentMakerSpec             string
	Name                            string
	Objective                       string
	ObjectiveName                   string
	OdaxObjective                   string
	OdaxObjectiveName               string
	OptimizationGoal                string
	OptimizationGoalName            string
	PausePeriods                    string
	PercentReachAtTargetFrequency   string
	PlacementBreakdown              string
	PlacementBreakdownMap           string
	PlanName                        string
	PlanType                        string
	PredictionMode                  string
	PredictionProgress              string
	ReferenceID                     string
	ReservationStatus               string
	StartTime                       string
	Status                          string
	StoryEventType                  string
	TargetCpm                       string
	TargetFrequency                 string
	TargetFrequencyResetPeriod      string
	TargetSpec                      string
	TimeCreated                     string
	TimeUpdated                     string
	TimezoneID                      string
	TimezoneName                    string
	ToplineID                       string
	TrendingTopicsSpec              string
	VideoViewLengthConstraint       string
	Viewtag                         string
}{
	AccountID:                       "account_id",
	ActivityStatus:                  "activity_status",
	AdFormats:                       "ad_formats",
	AuctionEntryOptionIndex:         "auction_entry_option_index",
	AudienceSizeLowerBound:          "audience_size_lower_bound",
	AudienceSizeUpperBound:          "audience_size_upper_bound",
	BusinessID:                      "business_id",
	BuyingType:                      "buying_type",
	CampaignGroupID:                 "campaign_group_id",
	CampaignID:                      "campaign_id",
	CampaignTimeStart:               "campaign_time_start",
	CampaignTimeStop:                "campaign_time_stop",
	Currency:                        "currency",
	CurveBudgetReach:                "curve_budget_reach",
	CurveReach:                      "curve_reach",
	DailyGrpCurve:                   "daily_grp_curve",
	DailyImpressionCurve:            "daily_impression_curve",
	DailyImpressionCurveMap:         "daily_impression_curve_map",
	DayPartingSchedule:              "day_parting_schedule",
	DestinationID:                   "destination_id",
	EndTime:                         "end_time",
	ExpirationTime:                  "expiration_time",
	ExternalBudget:                  "external_budget",
	ExternalImpression:              "external_impression",
	ExternalMaximumBudget:           "external_maximum_budget",
	ExternalMaximumImpression:       "external_maximum_impression",
	ExternalMaximumReach:            "external_maximum_reach",
	ExternalMinimumBudget:           "external_minimum_budget",
	ExternalMinimumImpression:       "external_minimum_impression",
	ExternalMinimumReach:            "external_minimum_reach",
	ExternalReach:                   "external_reach",
	FeedRatioX0000:                  "feed_ratio_0000",
	FrequencyCap:                    "frequency_cap",
	FrequencyDistributionMap:        "frequency_distribution_map",
	FrequencyDistributionMapAgg:     "frequency_distribution_map_agg",
	GrpAudienceSize:                 "grp_audience_size",
	GrpAvgProbabilityMap:            "grp_avg_probability_map",
	GrpCountryAudienceSize:          "grp_country_audience_size",
	GrpCurve:                        "grp_curve",
	GrpDmasAudienceSize:             "grp_dmas_audience_size",
	GrpFilteringThresholdX00:        "grp_filtering_threshold_00",
	GrpPoints:                       "grp_points",
	GrpRatio:                        "grp_ratio",
	GrpReachRatio:                   "grp_reach_ratio",
	GrpStatus:                       "grp_status",
	HoldoutPercentage:               "holdout_percentage",
	ID:                              "id",
	ImpressionCurve:                 "impression_curve",
	InstagramDestinationID:          "instagram_destination_id",
	InstreamPackages:                "instream_packages",
	IntervalFrequencyCap:            "interval_frequency_cap",
	IntervalFrequencyCapResetPeriod: "interval_frequency_cap_reset_period",
	IsBalancedFrequency:             "is_balanced_frequency",
	IsBonusMedia:                    "is_bonus_media",
	IsConversionGoal:                "is_conversion_goal",
	IsHigherAverageFrequency:        "is_higher_average_frequency",
	IsIo:                            "is_io",
	IsReservedBuying:                "is_reserved_buying",
	IsTrp:                           "is_trp",
	MetaMomentMakerSpec:             "meta_moment_maker_spec",
	Name:                            "name",
	Objective:                       "objective",
	ObjectiveName:                   "objective_name",
	OdaxObjective:                   "odax_objective",
	OdaxObjectiveName:               "odax_objective_name",
	OptimizationGoal:                "optimization_goal",
	OptimizationGoalName:            "optimization_goal_name",
	PausePeriods:                    "pause_periods",
	PercentReachAtTargetFrequency:   "percent_reach_at_target_frequency",
	PlacementBreakdown:              "placement_breakdown",
	PlacementBreakdownMap:           "placement_breakdown_map",
	PlanName:                        "plan_name",
	PlanType:                        "plan_type",
	PredictionMode:                  "prediction_mode",
	PredictionProgress:              "prediction_progress",
	ReferenceID:                     "reference_id",
	ReservationStatus:               "reservation_status",
	StartTime:                       "start_time",
	Status:                          "status",
	StoryEventType:                  "story_event_type",
	TargetCpm:                       "target_cpm",
	TargetFrequency:                 "target_frequency",
	TargetFrequencyResetPeriod:      "target_frequency_reset_period",
	TargetSpec:                      "target_spec",
	TimeCreated:                     "time_created",
	TimeUpdated:                     "time_updated",
	TimezoneID:                      "timezone_id",
	TimezoneName:                    "timezone_name",
	ToplineID:                       "topline_id",
	TrendingTopicsSpec:              "trending_topics_spec",
	VideoViewLengthConstraint:       "video_view_length_constraint",
	Viewtag:                         "viewtag",
}
