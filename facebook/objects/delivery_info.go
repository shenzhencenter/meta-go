package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type DeliveryInfo struct {
	ActiveAcceleratedCampaignCount            *int               `json:"active_accelerated_campaign_count,omitempty"`
	ActiveDayPartedCampaignCount              *int               `json:"active_day_parted_campaign_count,omitempty"`
	AdPenaltyMap                              *[]map[string]bool `json:"ad_penalty_map,omitempty"`
	AreAllDailyBudgetsSpent                   *bool              `json:"are_all_daily_budgets_spent,omitempty"`
	CreditNeededAdsCount                      *int               `json:"credit_needed_ads_count,omitempty"`
	EligibleForDeliveryInsights               *bool              `json:"eligible_for_delivery_insights,omitempty"`
	EndTime                                   *core.Time         `json:"end_time,omitempty"`
	HasAccountHitSpendLimit                   *bool              `json:"has_account_hit_spend_limit,omitempty"`
	HasCampaignGroupHitSpendLimit             *bool              `json:"has_campaign_group_hit_spend_limit,omitempty"`
	HasNoActiveAds                            *bool              `json:"has_no_active_ads,omitempty"`
	HasNoAds                                  *bool              `json:"has_no_ads,omitempty"`
	InactiveAdsCount                          *int               `json:"inactive_ads_count,omitempty"`
	InactiveCampaignCount                     *int               `json:"inactive_campaign_count,omitempty"`
	IsAccountClosed                           *bool              `json:"is_account_closed,omitempty"`
	IsAccountDisabled                         *bool              `json:"is_account_disabled,omitempty"`
	IsAdUneconomical                          *bool              `json:"is_ad_uneconomical,omitempty"`
	IsAdfarmPenalized                         *bool              `json:"is_adfarm_penalized,omitempty"`
	IsAdgroupPartiallyRejected                *bool              `json:"is_adgroup_partially_rejected,omitempty"`
	IsCampaignAccelerated                     *bool              `json:"is_campaign_accelerated,omitempty"`
	IsCampaignCompleted                       *bool              `json:"is_campaign_completed,omitempty"`
	IsCampaignDayParted                       *bool              `json:"is_campaign_day_parted,omitempty"`
	IsCampaignDisabled                        *bool              `json:"is_campaign_disabled,omitempty"`
	IsCampaignGroupDisabled                   *bool              `json:"is_campaign_group_disabled,omitempty"`
	IsClickbaitPenalized                      *bool              `json:"is_clickbait_penalized,omitempty"`
	IsDailyBudgetSpent                        *bool              `json:"is_daily_budget_spent,omitempty"`
	IsEngagementBaitPenalized                 *bool              `json:"is_engagement_bait_penalized,omitempty"`
	IsLqwePenalized                           *bool              `json:"is_lqwe_penalized,omitempty"`
	IsReachAndFrequencyMisconfigured          *bool              `json:"is_reach_and_frequency_misconfigured,omitempty"`
	IsSensationalismPenalized                 *bool              `json:"is_sensationalism_penalized,omitempty"`
	IsSplitTestActive                         *bool              `json:"is_split_test_active,omitempty"`
	IsSplitTestValid                          *bool              `json:"is_split_test_valid,omitempty"`
	LiftStudyTimePeriod                       *string            `json:"lift_study_time_period,omitempty"`
	NeedsCredit                               *bool              `json:"needs_credit,omitempty"`
	NeedsTaxNumber                            *bool              `json:"needs_tax_number,omitempty"`
	NonDeletedAdsCount                        *int               `json:"non_deleted_ads_count,omitempty"`
	NotDeliveringCampaignCount                *int               `json:"not_delivering_campaign_count,omitempty"`
	PendingAdsCount                           *int               `json:"pending_ads_count,omitempty"`
	ReachFrequencyCampaignUnderdeliveryReason *string            `json:"reach_frequency_campaign_underdelivery_reason,omitempty"`
	RejectedAdsCount                          *int               `json:"rejected_ads_count,omitempty"`
	StartTime                                 *core.Time         `json:"start_time,omitempty"`
	Status                                    *string            `json:"status,omitempty"`
	TextPenaltyLevel                          *string            `json:"text_penalty_level,omitempty"`
}

var DeliveryInfoFields = struct {
	ActiveAcceleratedCampaignCount            string
	ActiveDayPartedCampaignCount              string
	AdPenaltyMap                              string
	AreAllDailyBudgetsSpent                   string
	CreditNeededAdsCount                      string
	EligibleForDeliveryInsights               string
	EndTime                                   string
	HasAccountHitSpendLimit                   string
	HasCampaignGroupHitSpendLimit             string
	HasNoActiveAds                            string
	HasNoAds                                  string
	InactiveAdsCount                          string
	InactiveCampaignCount                     string
	IsAccountClosed                           string
	IsAccountDisabled                         string
	IsAdUneconomical                          string
	IsAdfarmPenalized                         string
	IsAdgroupPartiallyRejected                string
	IsCampaignAccelerated                     string
	IsCampaignCompleted                       string
	IsCampaignDayParted                       string
	IsCampaignDisabled                        string
	IsCampaignGroupDisabled                   string
	IsClickbaitPenalized                      string
	IsDailyBudgetSpent                        string
	IsEngagementBaitPenalized                 string
	IsLqwePenalized                           string
	IsReachAndFrequencyMisconfigured          string
	IsSensationalismPenalized                 string
	IsSplitTestActive                         string
	IsSplitTestValid                          string
	LiftStudyTimePeriod                       string
	NeedsCredit                               string
	NeedsTaxNumber                            string
	NonDeletedAdsCount                        string
	NotDeliveringCampaignCount                string
	PendingAdsCount                           string
	ReachFrequencyCampaignUnderdeliveryReason string
	RejectedAdsCount                          string
	StartTime                                 string
	Status                                    string
	TextPenaltyLevel                          string
}{
	ActiveAcceleratedCampaignCount:            "active_accelerated_campaign_count",
	ActiveDayPartedCampaignCount:              "active_day_parted_campaign_count",
	AdPenaltyMap:                              "ad_penalty_map",
	AreAllDailyBudgetsSpent:                   "are_all_daily_budgets_spent",
	CreditNeededAdsCount:                      "credit_needed_ads_count",
	EligibleForDeliveryInsights:               "eligible_for_delivery_insights",
	EndTime:                                   "end_time",
	HasAccountHitSpendLimit:                   "has_account_hit_spend_limit",
	HasCampaignGroupHitSpendLimit:             "has_campaign_group_hit_spend_limit",
	HasNoActiveAds:                            "has_no_active_ads",
	HasNoAds:                                  "has_no_ads",
	InactiveAdsCount:                          "inactive_ads_count",
	InactiveCampaignCount:                     "inactive_campaign_count",
	IsAccountClosed:                           "is_account_closed",
	IsAccountDisabled:                         "is_account_disabled",
	IsAdUneconomical:                          "is_ad_uneconomical",
	IsAdfarmPenalized:                         "is_adfarm_penalized",
	IsAdgroupPartiallyRejected:                "is_adgroup_partially_rejected",
	IsCampaignAccelerated:                     "is_campaign_accelerated",
	IsCampaignCompleted:                       "is_campaign_completed",
	IsCampaignDayParted:                       "is_campaign_day_parted",
	IsCampaignDisabled:                        "is_campaign_disabled",
	IsCampaignGroupDisabled:                   "is_campaign_group_disabled",
	IsClickbaitPenalized:                      "is_clickbait_penalized",
	IsDailyBudgetSpent:                        "is_daily_budget_spent",
	IsEngagementBaitPenalized:                 "is_engagement_bait_penalized",
	IsLqwePenalized:                           "is_lqwe_penalized",
	IsReachAndFrequencyMisconfigured:          "is_reach_and_frequency_misconfigured",
	IsSensationalismPenalized:                 "is_sensationalism_penalized",
	IsSplitTestActive:                         "is_split_test_active",
	IsSplitTestValid:                          "is_split_test_valid",
	LiftStudyTimePeriod:                       "lift_study_time_period",
	NeedsCredit:                               "needs_credit",
	NeedsTaxNumber:                            "needs_tax_number",
	NonDeletedAdsCount:                        "non_deleted_ads_count",
	NotDeliveringCampaignCount:                "not_delivering_campaign_count",
	PendingAdsCount:                           "pending_ads_count",
	ReachFrequencyCampaignUnderdeliveryReason: "reach_frequency_campaign_underdelivery_reason",
	RejectedAdsCount:                          "rejected_ads_count",
	StartTime:                                 "start_time",
	Status:                                    "status",
	TextPenaltyLevel:                          "text_penalty_level",
}
