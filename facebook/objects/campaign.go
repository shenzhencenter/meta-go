package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"time"
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
	CreatedTime                      *time.Time                         `json:"created_time,omitempty"`
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
	LastBudgetTogglingTime           *time.Time                         `json:"last_budget_toggling_time,omitempty"`
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
	StartTime                        *time.Time                         `json:"start_time,omitempty"`
	Status                           *enums.CampaignStatus              `json:"status,omitempty"`
	StopTime                         *time.Time                         `json:"stop_time,omitempty"`
	ToplineID                        *core.ID                           `json:"topline_id,omitempty"`
	UpdatedTime                      *time.Time                         `json:"updated_time,omitempty"`
}
