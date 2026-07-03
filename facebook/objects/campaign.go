package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type Campaign struct {
	AccountID                        *core.ID                           `json:"account_id,omitempty"`
	Adlabels                         *[]AdLabel                         `json:"adlabels,omitempty"`
	AdvantageStateInfo               *AdCampaignGroupAdvantageState     `json:"advantage_state_info,omitempty"`
	BidStrategy                      *enums.CampaignBidStrategy         `json:"bid_strategy,omitempty"`
	BoostedObjectID                  *core.ID                           `json:"boosted_object_id,omitempty"`
	BrandLiftStudies                 *[]AdStudy                         `json:"brand_lift_studies,omitempty"`
	BudgetRebalanceFlag              *bool                              `json:"budget_rebalance_flag,omitempty"`
	BudgetRemaining                  *string                            `json:"budget_remaining,omitempty"`
	BuyingType                       *string                            `json:"buying_type,omitempty"`
	CampaignGroupActiveTime          *string                            `json:"campaign_group_active_time,omitempty"`
	CanCreateBrandLiftStudy          *bool                              `json:"can_create_brand_lift_study,omitempty"`
	CanUseSpendCap                   *bool                              `json:"can_use_spend_cap,omitempty"`
	ConfiguredStatus                 *enums.CampaignConfiguredStatus    `json:"configured_status,omitempty"`
	CreatedTime                      *core.Time                         `json:"created_time,omitempty"`
	DailyBudget                      *string                            `json:"daily_budget,omitempty"`
	EffectiveStatus                  *enums.CampaignEffectiveStatus     `json:"effective_status,omitempty"`
	FrequencyControlSpecs            *[]AdCampaignFrequencyControlSpecs `json:"frequency_control_specs,omitempty"`
	HasSecondarySkadnetworkReporting *bool                              `json:"has_secondary_skadnetwork_reporting,omitempty"`
	ID                               *core.ID                           `json:"id,omitempty"`
	IsAdsetBudgetSharingEnabled      *bool                              `json:"is_adset_budget_sharing_enabled,omitempty"`
	IsBudgetScheduleEnabled          *bool                              `json:"is_budget_schedule_enabled,omitempty"`
	IsDirectSendCampaign             *bool                              `json:"is_direct_send_campaign,omitempty"`
	IsMessageCampaign                *bool                              `json:"is_message_campaign,omitempty"`
	IsMetaMomentMakerEnabled         *bool                              `json:"is_meta_moment_maker_enabled,omitempty"`
	IsReelsTrendingAdsEnabled        *bool                              `json:"is_reels_trending_ads_enabled,omitempty"`
	IsSkadnetworkAttribution         *bool                              `json:"is_skadnetwork_attribution,omitempty"`
	IssuesInfo                       *[]AdCampaignIssuesInfo            `json:"issues_info,omitempty"`
	LastBudgetTogglingTime           *core.Time                         `json:"last_budget_toggling_time,omitempty"`
	LifetimeBudget                   *string                            `json:"lifetime_budget,omitempty"`
	Name                             *string                            `json:"name,omitempty"`
	Objective                        *string                            `json:"objective,omitempty"`
	PacingType                       *[]string                          `json:"pacing_type,omitempty"`
	PrimaryAttribution               *string                            `json:"primary_attribution,omitempty"`
	PromotedObject                   *AdPromotedObject                  `json:"promoted_object,omitempty"`
	Recommendations                  *[]AdRecommendation                `json:"recommendations,omitempty"`
	SmartPromotionType               *string                            `json:"smart_promotion_type,omitempty"`
	SourceCampaign                   *Campaign                          `json:"source_campaign,omitempty"`
	SourceCampaignID                 *core.ID                           `json:"source_campaign_id,omitempty"`
	SourceRecommendationType         *string                            `json:"source_recommendation_type,omitempty"`
	SpecialAdCategories              *[]string                          `json:"special_ad_categories,omitempty"`
	SpecialAdCategory                *string                            `json:"special_ad_category,omitempty"`
	SpecialAdCategoryCountry         *[]string                          `json:"special_ad_category_country,omitempty"`
	SpendCap                         *string                            `json:"spend_cap,omitempty"`
	StartTime                        *core.Time                         `json:"start_time,omitempty"`
	Status                           *enums.CampaignStatus              `json:"status,omitempty"`
	StopTime                         *core.Time                         `json:"stop_time,omitempty"`
	ToplineID                        *core.ID                           `json:"topline_id,omitempty"`
	UpdatedTime                      *core.Time                         `json:"updated_time,omitempty"`
}

var CampaignFields = struct {
	AccountID                        string
	Adlabels                         string
	AdvantageStateInfo               string
	BidStrategy                      string
	BoostedObjectID                  string
	BrandLiftStudies                 string
	BudgetRebalanceFlag              string
	BudgetRemaining                  string
	BuyingType                       string
	CampaignGroupActiveTime          string
	CanCreateBrandLiftStudy          string
	CanUseSpendCap                   string
	ConfiguredStatus                 string
	CreatedTime                      string
	DailyBudget                      string
	EffectiveStatus                  string
	FrequencyControlSpecs            string
	HasSecondarySkadnetworkReporting string
	ID                               string
	IsAdsetBudgetSharingEnabled      string
	IsBudgetScheduleEnabled          string
	IsDirectSendCampaign             string
	IsMessageCampaign                string
	IsMetaMomentMakerEnabled         string
	IsReelsTrendingAdsEnabled        string
	IsSkadnetworkAttribution         string
	IssuesInfo                       string
	LastBudgetTogglingTime           string
	LifetimeBudget                   string
	Name                             string
	Objective                        string
	PacingType                       string
	PrimaryAttribution               string
	PromotedObject                   string
	Recommendations                  string
	SmartPromotionType               string
	SourceCampaign                   string
	SourceCampaignID                 string
	SourceRecommendationType         string
	SpecialAdCategories              string
	SpecialAdCategory                string
	SpecialAdCategoryCountry         string
	SpendCap                         string
	StartTime                        string
	Status                           string
	StopTime                         string
	ToplineID                        string
	UpdatedTime                      string
}{
	AccountID:                        "account_id",
	Adlabels:                         "adlabels",
	AdvantageStateInfo:               "advantage_state_info",
	BidStrategy:                      "bid_strategy",
	BoostedObjectID:                  "boosted_object_id",
	BrandLiftStudies:                 "brand_lift_studies",
	BudgetRebalanceFlag:              "budget_rebalance_flag",
	BudgetRemaining:                  "budget_remaining",
	BuyingType:                       "buying_type",
	CampaignGroupActiveTime:          "campaign_group_active_time",
	CanCreateBrandLiftStudy:          "can_create_brand_lift_study",
	CanUseSpendCap:                   "can_use_spend_cap",
	ConfiguredStatus:                 "configured_status",
	CreatedTime:                      "created_time",
	DailyBudget:                      "daily_budget",
	EffectiveStatus:                  "effective_status",
	FrequencyControlSpecs:            "frequency_control_specs",
	HasSecondarySkadnetworkReporting: "has_secondary_skadnetwork_reporting",
	ID:                               "id",
	IsAdsetBudgetSharingEnabled:      "is_adset_budget_sharing_enabled",
	IsBudgetScheduleEnabled:          "is_budget_schedule_enabled",
	IsDirectSendCampaign:             "is_direct_send_campaign",
	IsMessageCampaign:                "is_message_campaign",
	IsMetaMomentMakerEnabled:         "is_meta_moment_maker_enabled",
	IsReelsTrendingAdsEnabled:        "is_reels_trending_ads_enabled",
	IsSkadnetworkAttribution:         "is_skadnetwork_attribution",
	IssuesInfo:                       "issues_info",
	LastBudgetTogglingTime:           "last_budget_toggling_time",
	LifetimeBudget:                   "lifetime_budget",
	Name:                             "name",
	Objective:                        "objective",
	PacingType:                       "pacing_type",
	PrimaryAttribution:               "primary_attribution",
	PromotedObject:                   "promoted_object",
	Recommendations:                  "recommendations",
	SmartPromotionType:               "smart_promotion_type",
	SourceCampaign:                   "source_campaign",
	SourceCampaignID:                 "source_campaign_id",
	SourceRecommendationType:         "source_recommendation_type",
	SpecialAdCategories:              "special_ad_categories",
	SpecialAdCategory:                "special_ad_category",
	SpecialAdCategoryCountry:         "special_ad_category_country",
	SpendCap:                         "spend_cap",
	StartTime:                        "start_time",
	Status:                           "status",
	StopTime:                         "stop_time",
	ToplineID:                        "topline_id",
	UpdatedTime:                      "updated_time",
}
