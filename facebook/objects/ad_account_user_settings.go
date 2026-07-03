package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"time"
)

type AdAccountUserSettings struct {
	AcfShouldOptOutVideoAdjustments                     *bool                                                  `json:"acf_should_opt_out_video_adjustments,omitempty"`
	AcoStickySettings                                   *[]map[string]string                                   `json:"aco_sticky_settings,omitempty"`
	ActionsQuickViewCreated                             *bool                                                  `json:"actions_quick_view_created,omitempty"`
	ActiveAdsQuickViewCreated                           *bool                                                  `json:"active_ads_quick_view_created,omitempty"`
	AdAccount                                           *AdAccount                                             `json:"ad_account,omitempty"`
	AdObjectExportFormat                                *string                                                `json:"ad_object_export_format,omitempty"`
	AdsManagerFooterRowToastImpressions                 *int                                                   `json:"ads_manager_footer_row_toast_impressions,omitempty"`
	AutoReviewVideoCaption                              *bool                                                  `json:"auto_review_video_caption,omitempty"`
	BudgetOptimizationQuickViewCreated                  *bool                                                  `json:"budget_optimization_quick_view_created,omitempty"`
	CampaignOverviewColumns                             *[]string                                              `json:"campaign_overview_columns,omitempty"`
	ColumnSuggestionStatus                              *string                                                `json:"column_suggestion_status,omitempty"`
	ConditionalFormattingRules                          *[]string                                              `json:"conditional_formatting_rules,omitempty"`
	DefaultAccountOverviewAgegenderMetrics              *[]string                                              `json:"default_account_overview_agegender_metrics,omitempty"`
	DefaultAccountOverviewLocationMetrics               *[]string                                              `json:"default_account_overview_location_metrics,omitempty"`
	DefaultAccountOverviewMetrics                       *[]string                                              `json:"default_account_overview_metrics,omitempty"`
	DefaultAccountOverviewTimeMetrics                   *[]string                                              `json:"default_account_overview_time_metrics,omitempty"`
	DefaultBuiltinColumnPreset                          *string                                                `json:"default_builtin_column_preset,omitempty"`
	DefaultNamTimeRange                                 *string                                                `json:"default_nam_time_range,omitempty"`
	DraftModeEnabled                                    *bool                                                  `json:"draft_mode_enabled,omitempty"`
	ExportDeletedItemsWithDelivery                      *bool                                                  `json:"export_deleted_items_with_delivery,omitempty"`
	ExportSummaryRow                                    *bool                                                  `json:"export_summary_row,omitempty"`
	HadDeliveryQuickViewCreated                         *bool                                                  `json:"had_delivery_quick_view_created,omitempty"`
	HasSeenGroupsColumnFlexingExperience                *bool                                                  `json:"has_seen_groups_column_flexing_experience,omitempty"`
	HasSeenInstagramColumnFlexingExperience             *bool                                                  `json:"has_seen_instagram_column_flexing_experience,omitempty"`
	HasSeenLeadsColumnFlexingExperience                 *bool                                                  `json:"has_seen_leads_column_flexing_experience,omitempty"`
	HasSeenShopsAdsMetricsOnboardingTour                *bool                                                  `json:"has_seen_shops_ads_metrics_onboarding_tour,omitempty"`
	HasSeenShopsColumnFlexingExperience                 *bool                                                  `json:"has_seen_shops_column_flexing_experience,omitempty"`
	HasUsedQuickViewsPanel                              *bool                                                  `json:"has_used_quick_views_panel,omitempty"`
	HiddenOptimizationTips                              *[]map[string]bool                                     `json:"hidden_optimization_tips,omitempty"`
	HighPerformingQuickViewCreated                      *bool                                                  `json:"high_performing_quick_view_created,omitempty"`
	ID                                                  *core.ID                                               `json:"id,omitempty"`
	IsX3pAuthSettingSet                                 *bool                                                  `json:"is_3p_auth_setting_set,omitempty"`
	IsAdsManagerFooterRowPreferenceSet                  *bool                                                  `json:"is_ads_manager_footer_row_preference_set,omitempty"`
	IsAdsManagerFooterRowShown                          *bool                                                  `json:"is_ads_manager_footer_row_shown,omitempty"`
	LastUsedColumns                                     *map[string]interface{}                                `json:"last_used_columns,omitempty"`
	LastUsedPeFilters                                   *[]map[string]interface{}                              `json:"last_used_pe_filters,omitempty"`
	LastUsedWebsiteUrls                                 *[]string                                              `json:"last_used_website_urls,omitempty"`
	OutlierPreferences                                  *map[string]interface{}                                `json:"outlier_preferences,omitempty"`
	PinnedAdObjectIds                                   *[]core.ID                                             `json:"pinned_ad_object_ids,omitempty"`
	RbExportFormat                                      *string                                                `json:"rb_export_format,omitempty"`
	RbExportRawData                                     *bool                                                  `json:"rb_export_raw_data,omitempty"`
	RbExportSummaryRow                                  *bool                                                  `json:"rb_export_summary_row,omitempty"`
	RecentlyUsedQuickViews                              *[]string                                              `json:"recently_used_quick_views,omitempty"`
	SaipAdvertiserSetupOptimisationGuidanceOverallState *string                                                `json:"saip_advertiser_setup_optimisation_guidance_overall_state,omitempty"`
	SaipAdvertiserSetupOptimisationGuidanceState        *[]map[string]string                                   `json:"saip_advertiser_setup_optimisation_guidance_state,omitempty"`
	ShopsAdsMetricsOnboardingTourCloseCount             *int                                                   `json:"shops_ads_metrics_onboarding_tour_close_count,omitempty"`
	ShopsAdsMetricsOnboardingTourLastActionTime         *time.Time                                             `json:"shops_ads_metrics_onboarding_tour_last_action_time,omitempty"`
	ShouldDefaultImageAutoCrop                          *bool                                                  `json:"should_default_image_auto_crop,omitempty"`
	ShouldDefaultImageAutoCropForTail                   *bool                                                  `json:"should_default_image_auto_crop_for_tail,omitempty"`
	ShouldDefaultImageAutoCropOptimization              *bool                                                  `json:"should_default_image_auto_crop_optimization,omitempty"`
	ShouldDefaultImageDofToggle                         *bool                                                  `json:"should_default_image_dof_toggle,omitempty"`
	ShouldDefaultImageLppAdsToSquare                    *bool                                                  `json:"should_default_image_lpp_ads_to_square,omitempty"`
	ShouldDefaultInstagramProfileCardOptimization       *bool                                                  `json:"should_default_instagram_profile_card_optimization,omitempty"`
	ShouldDefaultTextSwappingOptimization               *bool                                                  `json:"should_default_text_swapping_optimization,omitempty"`
	ShouldLogoutOfX3pSourcing                           *bool                                                  `json:"should_logout_of_3p_sourcing,omitempty"`
	ShouldShowShopsAdsMetricsOnboardingTour             *bool                                                  `json:"should_show_shops_ads_metrics_onboarding_tour,omitempty"`
	ShowArchivedData                                    *bool                                                  `json:"show_archived_data,omitempty"`
	SydCampaignTrendsActivemetric                       *string                                                `json:"syd_campaign_trends_activemetric,omitempty"`
	SydCampaignTrendsAttribution                        *string                                                `json:"syd_campaign_trends_attribution,omitempty"`
	SydCampaignTrendsMetrics                            *[]string                                              `json:"syd_campaign_trends_metrics,omitempty"`
	SydCampaignTrendsObjective                          *enums.AdaccountusersettingsSydCampaignTrendsObjective `json:"syd_campaign_trends_objective,omitempty"`
	SydCampaignTrendsTimeRange                          *string                                                `json:"syd_campaign_trends_time_range,omitempty"`
	SydLandingPageOptInStatus                           *string                                                `json:"syd_landing_page_opt_in_status,omitempty"`
	TextGenPersonaOptInType                             *string                                                `json:"text_gen_persona_opt_in_type,omitempty"`
	TextVariationsHlOptInOutTs                          *time.Time                                             `json:"text_variations_hl_opt_in_out_ts,omitempty"`
	TextVariationsHlOptInType                           *string                                                `json:"text_variations_hl_opt_in_type,omitempty"`
	TextVariationsOptInOutTs                            *time.Time                                             `json:"text_variations_opt_in_out_ts,omitempty"`
	TextVariationsOptInType                             *string                                                `json:"text_variations_opt_in_type,omitempty"`
	User                                                *User                                                  `json:"user,omitempty"`
	ValueOptimizedQvCreated                             *bool                                                  `json:"value_optimized_qv_created,omitempty"`
	ValueQvNuxImpressions                               *int                                                   `json:"value_qv_nux_impressions,omitempty"`
	ValueSuggestedColumnStatus                          *string                                                `json:"value_suggested_column_status,omitempty"`
}
