package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsUserSettings struct {
	APlusCSurveySeen                                     *bool                     `json:"a_plus_c_survey_seen,omitempty"`
	AdaptiveGeoExclBannerSeenTs                          *[]map[string]int         `json:"adaptive_geo_excl_banner_seen_ts,omitempty"`
	AddOverlaysOptInStatus                               *string                   `json:"add_overlays_opt_in_status,omitempty"`
	AdgroupNameTemplate                                  *map[string]interface{}   `json:"adgroup_name_template,omitempty"`
	AdsCsCatalogOptOutTimestamp                          *[]map[string]int         `json:"ads_cs_catalog_opt_out_timestamp,omitempty"`
	AdsCsDynamicSeOptInStatus                            *string                   `json:"ads_cs_dynamic_se_opt_in_status,omitempty"`
	AdsCsDynamicSeOptOutTimestamp                        *[]map[string]int         `json:"ads_cs_dynamic_se_opt_out_timestamp,omitempty"`
	AdsCsSitelinksOptInStatus                            *string                   `json:"ads_cs_sitelinks_opt_in_status,omitempty"`
	AdsCsSitelinksOptOutTimestamp                        *[]map[string]int         `json:"ads_cs_sitelinks_opt_out_timestamp,omitempty"`
	AdsDestinationOptimizationOptOutTimestamp            *[]map[string]int         `json:"ads_destination_optimization_opt_out_timestamp,omitempty"`
	AdsToolVisits                                        *[]map[string]interface{} `json:"ads_tool_visits,omitempty"`
	ApluscAiAgentOptInStatus                             *string                   `json:"aplusc_ai_agent_opt_in_status,omitempty"`
	ApluscCarouselCdaOptInStatus                         *string                   `json:"aplusc_carousel_cda_opt_in_status,omitempty"`
	ApluscCarouselInlineCommentOptInStatus               *string                   `json:"aplusc_carousel_inline_comment_opt_in_status,omitempty"`
	ApluscDaOptInStatus                                  *string                   `json:"aplusc_da_opt_in_status,omitempty"`
	ApluscEnhanceCtaOptInStatus                          *string                   `json:"aplusc_enhance_cta_opt_in_status,omitempty"`
	ApluscEpaOptInStatus                                 *string                   `json:"aplusc_epa_opt_in_status,omitempty"`
	ApluscLocalStoreExtensionOptInStatus                 *string                   `json:"aplusc_local_store_extension_opt_in_status,omitempty"`
	ApluscOptOutFriction                                 *[]string                 `json:"aplusc_opt_out_friction,omitempty"`
	ApluscVideofilterOptInStatus                         *string                   `json:"aplusc_videofilter_opt_in_status,omitempty"`
	ApluscVideouncropOptInStatus                         *string                   `json:"aplusc_videouncrop_opt_in_status,omitempty"`
	AppDetailsDataOptInStatus                            *string                   `json:"app_details_data_opt_in_status,omitempty"`
	AutoflowLiteOptInStatus                              *string                   `json:"autoflow_lite_opt_in_status,omitempty"`
	AutoflowLiteShouldOptIn                              *bool                     `json:"autoflow_lite_should_opt_in,omitempty"`
	BlendedAdsCreationDefaultingOptInStatus              *string                   `json:"blended_ads_creation_defaulting_opt_in_status,omitempty"`
	BlendedAdsCreationDefaultingOptOutCampaignGroupIds   *[]core.ID                `json:"blended_ads_creation_defaulting_opt_out_campaign_group_ids,omitempty"`
	BookmarkedPages                                      *[]Page                   `json:"bookmarked_pages,omitempty"`
	CampaignGroupNameTemplate                            *map[string]interface{}   `json:"campaign_group_name_template,omitempty"`
	CampaignNameTemplate                                 *map[string]interface{}   `json:"campaign_name_template,omitempty"`
	CarouselToVideoOptInStatus                           *string                   `json:"carousel_to_video_opt_in_status,omitempty"`
	ConnectedSourcesCatalogOptInStatus                   *string                   `json:"connected_sources_catalog_opt_in_status,omitempty"`
	CreateCtaStickerOptInStatus                          *string                   `json:"create_cta_sticker_opt_in_status,omitempty"`
	CreativeFlexOptInStatus                              *string                   `json:"creative_flex_opt_in_status,omitempty"`
	DaAdaptImagesOptInStatus                             *string                   `json:"da_adapt_images_opt_in_status,omitempty"`
	DaAddOverlaysOptInStatus                             *string                   `json:"da_add_overlays_opt_in_status,omitempty"`
	DaCreativeFlexOptInStatus                            *string                   `json:"da_creative_flex_opt_in_status,omitempty"`
	DaHidePriceOptInStatus                               *string                   `json:"da_hide_price_opt_in_status,omitempty"`
	DaManualMediaNuxImpressions                          *int                      `json:"da_manual_media_nux_impressions,omitempty"`
	DcoToMmuOptOutStatus                                 *string                   `json:"dco_to_mmu_opt_out_status,omitempty"`
	DefaultCreationMode                                  *string                   `json:"default_creation_mode,omitempty"`
	DynamicPartnershipAdsOptInStatus                     *string                   `json:"dynamic_partnership_ads_opt_in_status,omitempty"`
	EnhanceCtaTextExtractionOptInStatus                  *string                   `json:"enhance_cta_text_extraction_opt_in_status,omitempty"`
	ExportFormatDefault                                  *string                   `json:"export_format_default,omitempty"`
	FeedbackSurveys                                      *[]string                 `json:"feedback_surveys,omitempty"`
	FfToMmuOptOutStatus                                  *string                   `json:"ff_to_mmu_opt_out_status,omitempty"`
	FocusModeDefault                                     *string                   `json:"focus_mode_default,omitempty"`
	GenAiAlphaTestStatus                                 *int                      `json:"gen_ai_alpha_test_status,omitempty"`
	GenAiAutoSelectOptInStatus                           *string                   `json:"gen_ai_auto_select_opt_in_status,omitempty"`
	ID                                                   *core.ID                  `json:"id,omitempty"`
	ImageBackgroundGenerationOptInStatus                 *string                   `json:"image_background_generation_opt_in_status,omitempty"`
	ImageBrightnessAndContrastOptInStatus                *string                   `json:"image_brightness_and_contrast_opt_in_status,omitempty"`
	ImageExpansionOptInStatus                            *string                   `json:"image_expansion_opt_in_status,omitempty"`
	ImageTemplatesTextExtractionOptInStatus              *string                   `json:"image_templates_text_extraction_opt_in_status,omitempty"`
	ImageTextTranslationOptInStatus                      *string                   `json:"image_text_translation_opt_in_status,omitempty"`
	IsAdsAiConsented                                     *bool                     `json:"is_ads_ai_consented,omitempty"`
	IsCboDefaultOn                                       *bool                     `json:"is_cbo_default_on,omitempty"`
	IsSeRemovalGuidanceDismissed                         *bool                     `json:"is_se_removal_guidance_dismissed,omitempty"`
	LastUsedPostFormat                                   *string                   `json:"last_used_post_format,omitempty"`
	LastVisitedTime                                      *core.Time                `json:"last_visited_time,omitempty"`
	MetadataBrandKitLastOptOutTimestamp                  *int                      `json:"metadata_brand_kit_last_opt_out_timestamp,omitempty"`
	MetadataBrandKitOptInStatus                          *string                   `json:"metadata_brand_kit_opt_in_status,omitempty"`
	MultiMediaOptOutStatus                               *string                   `json:"multi_media_opt_out_status,omitempty"`
	MusicOnReelsOptIn                                    *[]map[string]string      `json:"music_on_reels_opt_in,omitempty"`
	MutedCboMidflightEducationMessages                   *[]string                 `json:"muted_cbo_midflight_education_messages,omitempty"`
	OnsiteDestinationOptimizationOptIn                   *string                   `json:"onsite_destination_optimization_opt_in,omitempty"`
	OpenTabs                                             *[]string                 `json:"open_tabs,omitempty"`
	PacRelaxationOptInStatus                             *string                   `json:"pac_relaxation_opt_in_status,omitempty"`
	PcauCatOptoutSurveyImpr                              *int                      `json:"pcau_cat_optout_survey_impr,omitempty"`
	PcauCatOptoutSurveyRespTs                            *core.Time                `json:"pcau_cat_optout_survey_resp_ts,omitempty"`
	PeAiRelevancyOptOutTs                                *core.Time                `json:"pe_ai_relevancy_opt_out_ts,omitempty"`
	PeShowProductsSurveyImpr                             *int                      `json:"pe_show_products_survey_impr,omitempty"`
	PeShowProductsSurveyRespTs                           *core.Time                `json:"pe_show_products_survey_resp_ts,omitempty"`
	PlacementGroupSquareOptInStatus                      *string                   `json:"placement_group_square_opt_in_status,omitempty"`
	PlacementGroupVerticalOptInStatus                    *string                   `json:"placement_group_vertical_opt_in_status,omitempty"`
	PreviouslySeenRecommendations                        *[]string                 `json:"previously_seen_recommendations,omitempty"`
	ProductExtensionsOptIn                               *string                   `json:"product_extensions_opt_in,omitempty"`
	ReactiveControlSettings                              *[]map[string]interface{} `json:"reactive_control_settings,omitempty"`
	ReplaceMediaTextOptInStatus                          *string                   `json:"replace_media_text_opt_in_status,omitempty"`
	SaOffConvLocSeen                                     *string                   `json:"sa_off_conv_loc_seen,omitempty"`
	SaoffPublishedL2ConvLocSeen                          *string                   `json:"saoff_published_l2_conv_loc_seen,omitempty"`
	SaonMigrL1SeenStatus                                 *string                   `json:"saon_migr_l1_seen_status,omitempty"`
	SelectedAdAccount                                    *AdAccount                `json:"selected_ad_account,omitempty"`
	SelectedComparisonTimerange                          *map[string]interface{}   `json:"selected_comparison_timerange,omitempty"`
	SelectedMetricCic                                    *string                   `json:"selected_metric_cic,omitempty"`
	SelectedMetricsCic                                   *[]string                 `json:"selected_metrics_cic,omitempty"`
	SelectedPage                                         *Page                     `json:"selected_page,omitempty"`
	SelectedPageSection                                  *string                   `json:"selected_page_section,omitempty"`
	SelectedPowerEditorPane                              *string                   `json:"selected_power_editor_pane,omitempty"`
	SelectedStatRange                                    *map[string]interface{}   `json:"selected_stat_range,omitempty"`
	ShouldExportFilterEmptyCols                          *string                   `json:"should_export_filter_empty_cols,omitempty"`
	ShouldExportRowsWithoutUnsupportedFeature            *string                   `json:"should_export_rows_without_unsupported_feature,omitempty"`
	ShouldNotAutoExpandTreeTable                         *bool                     `json:"should_not_auto_expand_tree_table,omitempty"`
	ShouldNotShowCboCampaignToggleOffConfirmationMessage *bool                     `json:"should_not_show_cbo_campaign_toggle_off_confirmation_message,omitempty"`
	ShouldNotShowPublishMessageOnEditorClose             *bool                     `json:"should_not_show_publish_message_on_editor_close,omitempty"`
	ShowOriginalVideosOptIn                              *string                   `json:"show_original_videos_opt_in,omitempty"`
	ShowSummaryOptInStatus                               *string                   `json:"show_summary_opt_in_status,omitempty"`
	StaticAdProductExtensionsOptIn                       *string                   `json:"static_ad_product_extensions_opt_in,omitempty"`
	StickySettingAfterDefaultOn                          *string                   `json:"sticky_setting_after_default_on,omitempty"`
	SydCampaignTrendsMetric                              *string                   `json:"syd_campaign_trends_metric,omitempty"`
	TextOptimizationsTextExtractionOptInStatus           *string                   `json:"text_optimizations_text_extraction_opt_in_status,omitempty"`
	TextTranslationOptInStatus                           *string                   `json:"text_translation_opt_in_status,omitempty"`
	TextUnificationOptInStatus                           *string                   `json:"text_unification_opt_in_status,omitempty"`
	TextUnificationOptInStatusV2                         *string                   `json:"text_unification_opt_in_status_v2,omitempty"`
	TextVariationsStickyOptInStatus                      *string                   `json:"text_variations_sticky_opt_in_status,omitempty"`
	TotalCouponSydDismissals                             *int                      `json:"total_coupon_syd_dismissals,omitempty"`
	TotalCouponUpsellDismissals                          *int                      `json:"total_coupon_upsell_dismissals,omitempty"`
	URLPrefillRemovalTimestamp                           *int                      `json:"url_prefill_removal_timestamp,omitempty"`
	UsePeCreateFlow                                      *bool                     `json:"use_pe_create_flow,omitempty"`
	UseStepperPrimaryEntry                               *bool                     `json:"use_stepper_primary_entry,omitempty"`
	User                                                 *User                     `json:"user,omitempty"`
	VideoToImageOptInStatus                              *string                   `json:"video_to_image_opt_in_status,omitempty"`
	VoiceoverTransOptInStatus                            *string                   `json:"voiceover_trans_opt_in_status,omitempty"`
	WebsiteMediaOptInStatus                              *string                   `json:"website_media_opt_in_status,omitempty"`
	WebsiteReviewsDataOptInStatus                        *string                   `json:"website_reviews_data_opt_in_status,omitempty"`
	WebsiteSellingPointsDataOptInStatus                  *string                   `json:"website_selling_points_data_opt_in_status,omitempty"`
}

var AdsUserSettingsFields = struct {
	APlusCSurveySeen                                     string
	AdaptiveGeoExclBannerSeenTs                          string
	AddOverlaysOptInStatus                               string
	AdgroupNameTemplate                                  string
	AdsCsCatalogOptOutTimestamp                          string
	AdsCsDynamicSeOptInStatus                            string
	AdsCsDynamicSeOptOutTimestamp                        string
	AdsCsSitelinksOptInStatus                            string
	AdsCsSitelinksOptOutTimestamp                        string
	AdsDestinationOptimizationOptOutTimestamp            string
	AdsToolVisits                                        string
	ApluscAiAgentOptInStatus                             string
	ApluscCarouselCdaOptInStatus                         string
	ApluscCarouselInlineCommentOptInStatus               string
	ApluscDaOptInStatus                                  string
	ApluscEnhanceCtaOptInStatus                          string
	ApluscEpaOptInStatus                                 string
	ApluscLocalStoreExtensionOptInStatus                 string
	ApluscOptOutFriction                                 string
	ApluscVideofilterOptInStatus                         string
	ApluscVideouncropOptInStatus                         string
	AppDetailsDataOptInStatus                            string
	AutoflowLiteOptInStatus                              string
	AutoflowLiteShouldOptIn                              string
	BlendedAdsCreationDefaultingOptInStatus              string
	BlendedAdsCreationDefaultingOptOutCampaignGroupIds   string
	BookmarkedPages                                      string
	CampaignGroupNameTemplate                            string
	CampaignNameTemplate                                 string
	CarouselToVideoOptInStatus                           string
	ConnectedSourcesCatalogOptInStatus                   string
	CreateCtaStickerOptInStatus                          string
	CreativeFlexOptInStatus                              string
	DaAdaptImagesOptInStatus                             string
	DaAddOverlaysOptInStatus                             string
	DaCreativeFlexOptInStatus                            string
	DaHidePriceOptInStatus                               string
	DaManualMediaNuxImpressions                          string
	DcoToMmuOptOutStatus                                 string
	DefaultCreationMode                                  string
	DynamicPartnershipAdsOptInStatus                     string
	EnhanceCtaTextExtractionOptInStatus                  string
	ExportFormatDefault                                  string
	FeedbackSurveys                                      string
	FfToMmuOptOutStatus                                  string
	FocusModeDefault                                     string
	GenAiAlphaTestStatus                                 string
	GenAiAutoSelectOptInStatus                           string
	ID                                                   string
	ImageBackgroundGenerationOptInStatus                 string
	ImageBrightnessAndContrastOptInStatus                string
	ImageExpansionOptInStatus                            string
	ImageTemplatesTextExtractionOptInStatus              string
	ImageTextTranslationOptInStatus                      string
	IsAdsAiConsented                                     string
	IsCboDefaultOn                                       string
	IsSeRemovalGuidanceDismissed                         string
	LastUsedPostFormat                                   string
	LastVisitedTime                                      string
	MetadataBrandKitLastOptOutTimestamp                  string
	MetadataBrandKitOptInStatus                          string
	MultiMediaOptOutStatus                               string
	MusicOnReelsOptIn                                    string
	MutedCboMidflightEducationMessages                   string
	OnsiteDestinationOptimizationOptIn                   string
	OpenTabs                                             string
	PacRelaxationOptInStatus                             string
	PcauCatOptoutSurveyImpr                              string
	PcauCatOptoutSurveyRespTs                            string
	PeAiRelevancyOptOutTs                                string
	PeShowProductsSurveyImpr                             string
	PeShowProductsSurveyRespTs                           string
	PlacementGroupSquareOptInStatus                      string
	PlacementGroupVerticalOptInStatus                    string
	PreviouslySeenRecommendations                        string
	ProductExtensionsOptIn                               string
	ReactiveControlSettings                              string
	ReplaceMediaTextOptInStatus                          string
	SaOffConvLocSeen                                     string
	SaoffPublishedL2ConvLocSeen                          string
	SaonMigrL1SeenStatus                                 string
	SelectedAdAccount                                    string
	SelectedComparisonTimerange                          string
	SelectedMetricCic                                    string
	SelectedMetricsCic                                   string
	SelectedPage                                         string
	SelectedPageSection                                  string
	SelectedPowerEditorPane                              string
	SelectedStatRange                                    string
	ShouldExportFilterEmptyCols                          string
	ShouldExportRowsWithoutUnsupportedFeature            string
	ShouldNotAutoExpandTreeTable                         string
	ShouldNotShowCboCampaignToggleOffConfirmationMessage string
	ShouldNotShowPublishMessageOnEditorClose             string
	ShowOriginalVideosOptIn                              string
	ShowSummaryOptInStatus                               string
	StaticAdProductExtensionsOptIn                       string
	StickySettingAfterDefaultOn                          string
	SydCampaignTrendsMetric                              string
	TextOptimizationsTextExtractionOptInStatus           string
	TextTranslationOptInStatus                           string
	TextUnificationOptInStatus                           string
	TextUnificationOptInStatusV2                         string
	TextVariationsStickyOptInStatus                      string
	TotalCouponSydDismissals                             string
	TotalCouponUpsellDismissals                          string
	URLPrefillRemovalTimestamp                           string
	UsePeCreateFlow                                      string
	UseStepperPrimaryEntry                               string
	User                                                 string
	VideoToImageOptInStatus                              string
	VoiceoverTransOptInStatus                            string
	WebsiteMediaOptInStatus                              string
	WebsiteReviewsDataOptInStatus                        string
	WebsiteSellingPointsDataOptInStatus                  string
}{
	APlusCSurveySeen:                                     "a_plus_c_survey_seen",
	AdaptiveGeoExclBannerSeenTs:                          "adaptive_geo_excl_banner_seen_ts",
	AddOverlaysOptInStatus:                               "add_overlays_opt_in_status",
	AdgroupNameTemplate:                                  "adgroup_name_template",
	AdsCsCatalogOptOutTimestamp:                          "ads_cs_catalog_opt_out_timestamp",
	AdsCsDynamicSeOptInStatus:                            "ads_cs_dynamic_se_opt_in_status",
	AdsCsDynamicSeOptOutTimestamp:                        "ads_cs_dynamic_se_opt_out_timestamp",
	AdsCsSitelinksOptInStatus:                            "ads_cs_sitelinks_opt_in_status",
	AdsCsSitelinksOptOutTimestamp:                        "ads_cs_sitelinks_opt_out_timestamp",
	AdsDestinationOptimizationOptOutTimestamp:            "ads_destination_optimization_opt_out_timestamp",
	AdsToolVisits:                                        "ads_tool_visits",
	ApluscAiAgentOptInStatus:                             "aplusc_ai_agent_opt_in_status",
	ApluscCarouselCdaOptInStatus:                         "aplusc_carousel_cda_opt_in_status",
	ApluscCarouselInlineCommentOptInStatus:               "aplusc_carousel_inline_comment_opt_in_status",
	ApluscDaOptInStatus:                                  "aplusc_da_opt_in_status",
	ApluscEnhanceCtaOptInStatus:                          "aplusc_enhance_cta_opt_in_status",
	ApluscEpaOptInStatus:                                 "aplusc_epa_opt_in_status",
	ApluscLocalStoreExtensionOptInStatus:                 "aplusc_local_store_extension_opt_in_status",
	ApluscOptOutFriction:                                 "aplusc_opt_out_friction",
	ApluscVideofilterOptInStatus:                         "aplusc_videofilter_opt_in_status",
	ApluscVideouncropOptInStatus:                         "aplusc_videouncrop_opt_in_status",
	AppDetailsDataOptInStatus:                            "app_details_data_opt_in_status",
	AutoflowLiteOptInStatus:                              "autoflow_lite_opt_in_status",
	AutoflowLiteShouldOptIn:                              "autoflow_lite_should_opt_in",
	BlendedAdsCreationDefaultingOptInStatus:              "blended_ads_creation_defaulting_opt_in_status",
	BlendedAdsCreationDefaultingOptOutCampaignGroupIds:   "blended_ads_creation_defaulting_opt_out_campaign_group_ids",
	BookmarkedPages:                                      "bookmarked_pages",
	CampaignGroupNameTemplate:                            "campaign_group_name_template",
	CampaignNameTemplate:                                 "campaign_name_template",
	CarouselToVideoOptInStatus:                           "carousel_to_video_opt_in_status",
	ConnectedSourcesCatalogOptInStatus:                   "connected_sources_catalog_opt_in_status",
	CreateCtaStickerOptInStatus:                          "create_cta_sticker_opt_in_status",
	CreativeFlexOptInStatus:                              "creative_flex_opt_in_status",
	DaAdaptImagesOptInStatus:                             "da_adapt_images_opt_in_status",
	DaAddOverlaysOptInStatus:                             "da_add_overlays_opt_in_status",
	DaCreativeFlexOptInStatus:                            "da_creative_flex_opt_in_status",
	DaHidePriceOptInStatus:                               "da_hide_price_opt_in_status",
	DaManualMediaNuxImpressions:                          "da_manual_media_nux_impressions",
	DcoToMmuOptOutStatus:                                 "dco_to_mmu_opt_out_status",
	DefaultCreationMode:                                  "default_creation_mode",
	DynamicPartnershipAdsOptInStatus:                     "dynamic_partnership_ads_opt_in_status",
	EnhanceCtaTextExtractionOptInStatus:                  "enhance_cta_text_extraction_opt_in_status",
	ExportFormatDefault:                                  "export_format_default",
	FeedbackSurveys:                                      "feedback_surveys",
	FfToMmuOptOutStatus:                                  "ff_to_mmu_opt_out_status",
	FocusModeDefault:                                     "focus_mode_default",
	GenAiAlphaTestStatus:                                 "gen_ai_alpha_test_status",
	GenAiAutoSelectOptInStatus:                           "gen_ai_auto_select_opt_in_status",
	ID:                                                   "id",
	ImageBackgroundGenerationOptInStatus:                 "image_background_generation_opt_in_status",
	ImageBrightnessAndContrastOptInStatus:                "image_brightness_and_contrast_opt_in_status",
	ImageExpansionOptInStatus:                            "image_expansion_opt_in_status",
	ImageTemplatesTextExtractionOptInStatus:              "image_templates_text_extraction_opt_in_status",
	ImageTextTranslationOptInStatus:                      "image_text_translation_opt_in_status",
	IsAdsAiConsented:                                     "is_ads_ai_consented",
	IsCboDefaultOn:                                       "is_cbo_default_on",
	IsSeRemovalGuidanceDismissed:                         "is_se_removal_guidance_dismissed",
	LastUsedPostFormat:                                   "last_used_post_format",
	LastVisitedTime:                                      "last_visited_time",
	MetadataBrandKitLastOptOutTimestamp:                  "metadata_brand_kit_last_opt_out_timestamp",
	MetadataBrandKitOptInStatus:                          "metadata_brand_kit_opt_in_status",
	MultiMediaOptOutStatus:                               "multi_media_opt_out_status",
	MusicOnReelsOptIn:                                    "music_on_reels_opt_in",
	MutedCboMidflightEducationMessages:                   "muted_cbo_midflight_education_messages",
	OnsiteDestinationOptimizationOptIn:                   "onsite_destination_optimization_opt_in",
	OpenTabs:                                             "open_tabs",
	PacRelaxationOptInStatus:                             "pac_relaxation_opt_in_status",
	PcauCatOptoutSurveyImpr:                              "pcau_cat_optout_survey_impr",
	PcauCatOptoutSurveyRespTs:                            "pcau_cat_optout_survey_resp_ts",
	PeAiRelevancyOptOutTs:                                "pe_ai_relevancy_opt_out_ts",
	PeShowProductsSurveyImpr:                             "pe_show_products_survey_impr",
	PeShowProductsSurveyRespTs:                           "pe_show_products_survey_resp_ts",
	PlacementGroupSquareOptInStatus:                      "placement_group_square_opt_in_status",
	PlacementGroupVerticalOptInStatus:                    "placement_group_vertical_opt_in_status",
	PreviouslySeenRecommendations:                        "previously_seen_recommendations",
	ProductExtensionsOptIn:                               "product_extensions_opt_in",
	ReactiveControlSettings:                              "reactive_control_settings",
	ReplaceMediaTextOptInStatus:                          "replace_media_text_opt_in_status",
	SaOffConvLocSeen:                                     "sa_off_conv_loc_seen",
	SaoffPublishedL2ConvLocSeen:                          "saoff_published_l2_conv_loc_seen",
	SaonMigrL1SeenStatus:                                 "saon_migr_l1_seen_status",
	SelectedAdAccount:                                    "selected_ad_account",
	SelectedComparisonTimerange:                          "selected_comparison_timerange",
	SelectedMetricCic:                                    "selected_metric_cic",
	SelectedMetricsCic:                                   "selected_metrics_cic",
	SelectedPage:                                         "selected_page",
	SelectedPageSection:                                  "selected_page_section",
	SelectedPowerEditorPane:                              "selected_power_editor_pane",
	SelectedStatRange:                                    "selected_stat_range",
	ShouldExportFilterEmptyCols:                          "should_export_filter_empty_cols",
	ShouldExportRowsWithoutUnsupportedFeature:            "should_export_rows_without_unsupported_feature",
	ShouldNotAutoExpandTreeTable:                         "should_not_auto_expand_tree_table",
	ShouldNotShowCboCampaignToggleOffConfirmationMessage: "should_not_show_cbo_campaign_toggle_off_confirmation_message",
	ShouldNotShowPublishMessageOnEditorClose:             "should_not_show_publish_message_on_editor_close",
	ShowOriginalVideosOptIn:                              "show_original_videos_opt_in",
	ShowSummaryOptInStatus:                               "show_summary_opt_in_status",
	StaticAdProductExtensionsOptIn:                       "static_ad_product_extensions_opt_in",
	StickySettingAfterDefaultOn:                          "sticky_setting_after_default_on",
	SydCampaignTrendsMetric:                              "syd_campaign_trends_metric",
	TextOptimizationsTextExtractionOptInStatus:           "text_optimizations_text_extraction_opt_in_status",
	TextTranslationOptInStatus:                           "text_translation_opt_in_status",
	TextUnificationOptInStatus:                           "text_unification_opt_in_status",
	TextUnificationOptInStatusV2:                         "text_unification_opt_in_status_v2",
	TextVariationsStickyOptInStatus:                      "text_variations_sticky_opt_in_status",
	TotalCouponSydDismissals:                             "total_coupon_syd_dismissals",
	TotalCouponUpsellDismissals:                          "total_coupon_upsell_dismissals",
	URLPrefillRemovalTimestamp:                           "url_prefill_removal_timestamp",
	UsePeCreateFlow:                                      "use_pe_create_flow",
	UseStepperPrimaryEntry:                               "use_stepper_primary_entry",
	User:                                                 "user",
	VideoToImageOptInStatus:                              "video_to_image_opt_in_status",
	VoiceoverTransOptInStatus:                            "voiceover_trans_opt_in_status",
	WebsiteMediaOptInStatus:                              "website_media_opt_in_status",
	WebsiteReviewsDataOptInStatus:                        "website_reviews_data_opt_in_status",
	WebsiteSellingPointsDataOptInStatus:                  "website_selling_points_data_opt_in_status",
}
