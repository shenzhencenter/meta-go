package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsInsights struct {
	AccountCurrency                                             *string                   `json:"account_currency,omitempty"`
	AccountID                                                   *core.ID                  `json:"account_id,omitempty"`
	AccountName                                                 *string                   `json:"account_name,omitempty"`
	ActionValues                                                *[]AdsActionStats         `json:"action_values,omitempty"`
	Actions                                                     *[]AdsActionStats         `json:"actions,omitempty"`
	AdClickActions                                              *[]AdsActionStats         `json:"ad_click_actions,omitempty"`
	AdID                                                        *core.ID                  `json:"ad_id,omitempty"`
	AdImpressionActions                                         *[]AdsActionStats         `json:"ad_impression_actions,omitempty"`
	AdName                                                      *string                   `json:"ad_name,omitempty"`
	AdsetEnd                                                    *string                   `json:"adset_end,omitempty"`
	AdsetID                                                     *core.ID                  `json:"adset_id,omitempty"`
	AdsetName                                                   *string                   `json:"adset_name,omitempty"`
	AdsetStart                                                  *string                   `json:"adset_start,omitempty"`
	AdvancedActionsX28dView                                     *[]AdsActionStats         `json:"advanced_actions_28d_view,omitempty"`
	AdvancedReachX1dLookback                                    *string                   `json:"advanced_reach_1d_lookback,omitempty"`
	AdvancedReachX28dLookback                                   *string                   `json:"advanced_reach_28d_lookback,omitempty"`
	AdvancedReachX7dLookback                                    *string                   `json:"advanced_reach_7d_lookback,omitempty"`
	AgeTargeting                                                *string                   `json:"age_targeting,omitempty"`
	AnchorEventAttributionSetting                               *string                   `json:"anchor_event_attribution_setting,omitempty"`
	AnchorEventsPerformanceIndicator                            *string                   `json:"anchor_events_performance_indicator,omitempty"`
	AttributionSetting                                          *string                   `json:"attribution_setting,omitempty"`
	AuctionBid                                                  *string                   `json:"auction_bid,omitempty"`
	AuctionCompetitiveness                                      *string                   `json:"auction_competitiveness,omitempty"`
	AuctionMaxCompetitorBid                                     *string                   `json:"auction_max_competitor_bid,omitempty"`
	AveragePurchasesConversionValue                             *[]AdsActionStats         `json:"average_purchases_conversion_value,omitempty"`
	BuyingType                                                  *string                   `json:"buying_type,omitempty"`
	CampaignID                                                  *core.ID                  `json:"campaign_id,omitempty"`
	CampaignName                                                *string                   `json:"campaign_name,omitempty"`
	CanvasAvgViewPercent                                        *string                   `json:"canvas_avg_view_percent,omitempty"`
	CanvasAvgViewTime                                           *string                   `json:"canvas_avg_view_time,omitempty"`
	CatalogSegmentActions                                       *[]AdsActionStats         `json:"catalog_segment_actions,omitempty"`
	CatalogSegmentValue                                         *[]AdsActionStats         `json:"catalog_segment_value,omitempty"`
	CatalogSegmentValueMobilePurchaseRoas                       *[]AdsActionStats         `json:"catalog_segment_value_mobile_purchase_roas,omitempty"`
	CatalogSegmentValueOmniPurchaseRoas                         *[]AdsActionStats         `json:"catalog_segment_value_omni_purchase_roas,omitempty"`
	CatalogSegmentValueWebsitePurchaseRoas                      *[]AdsActionStats         `json:"catalog_segment_value_website_purchase_roas,omitempty"`
	Clicks                                                      *string                   `json:"clicks,omitempty"`
	ConfigurableAttributionAction                               *[]AdsActionStats         `json:"configurable_attribution_action,omitempty"`
	ConfigurableAttributionActionvalue                          *[]AdsActionStats         `json:"configurable_attribution_actionvalue,omitempty"`
	ConfigurableAudienceOverlapReach                            *[]AdsActionStats         `json:"configurable_audience_overlap_reach,omitempty"`
	ConfigurableReachbyfrequencyAction                          *[]AdsActionStats         `json:"configurable_reachbyfrequency_action,omitempty"`
	ConfigurableReachbyfrequencyConvertersCount                 *string                   `json:"configurable_reachbyfrequency_converters_count,omitempty"`
	ConfigurableReachbyfrequencyImpressionsCost                 *string                   `json:"configurable_reachbyfrequency_impressions_cost,omitempty"`
	ConfigurableReachbyfrequencyImpressionsCount                *string                   `json:"configurable_reachbyfrequency_impressions_count,omitempty"`
	ConfigurableReachbyfrequencyReach                           *string                   `json:"configurable_reachbyfrequency_reach,omitempty"`
	ConversionLeadRate                                          *[]AdsActionStats         `json:"conversion_lead_rate,omitempty"`
	ConversionLeads                                             *[]AdsActionStats         `json:"conversion_leads,omitempty"`
	ConversionRateRanking                                       *string                   `json:"conversion_rate_ranking,omitempty"`
	ConversionValues                                            *[]AdsActionStats         `json:"conversion_values,omitempty"`
	Conversions                                                 *[]AdsActionStats         `json:"conversions,omitempty"`
	ConvertedProductAppCustomEventFbMobilePurchase              *[]AdsActionStats         `json:"converted_product_app_custom_event_fb_mobile_purchase,omitempty"`
	ConvertedProductAppCustomEventFbMobilePurchaseValue         *[]AdsActionStats         `json:"converted_product_app_custom_event_fb_mobile_purchase_value,omitempty"`
	ConvertedProductOfflinePurchase                             *[]AdsActionStats         `json:"converted_product_offline_purchase,omitempty"`
	ConvertedProductOfflinePurchaseValue                        *[]AdsActionStats         `json:"converted_product_offline_purchase_value,omitempty"`
	ConvertedProductOmniPurchase                                *[]AdsActionStats         `json:"converted_product_omni_purchase,omitempty"`
	ConvertedProductOmniPurchaseValues                          *[]AdsActionStats         `json:"converted_product_omni_purchase_values,omitempty"`
	ConvertedProductQuantity                                    *[]AdsActionStats         `json:"converted_product_quantity,omitempty"`
	ConvertedProductValue                                       *[]AdsActionStats         `json:"converted_product_value,omitempty"`
	ConvertedProductWebsitePixelPurchase                        *[]AdsActionStats         `json:"converted_product_website_pixel_purchase,omitempty"`
	ConvertedProductWebsitePixelPurchaseValue                   *[]AdsActionStats         `json:"converted_product_website_pixel_purchase_value,omitempty"`
	ConvertedPromotedProductAppCustomEventFbMobilePurchase      *[]AdsActionStats         `json:"converted_promoted_product_app_custom_event_fb_mobile_purchase,omitempty"`
	ConvertedPromotedProductAppCustomEventFbMobilePurchaseValue *[]AdsActionStats         `json:"converted_promoted_product_app_custom_event_fb_mobile_purchase_value,omitempty"`
	ConvertedPromotedProductOfflinePurchase                     *[]AdsActionStats         `json:"converted_promoted_product_offline_purchase,omitempty"`
	ConvertedPromotedProductOfflinePurchaseValue                *[]AdsActionStats         `json:"converted_promoted_product_offline_purchase_value,omitempty"`
	ConvertedPromotedProductOmniPurchase                        *[]AdsActionStats         `json:"converted_promoted_product_omni_purchase,omitempty"`
	ConvertedPromotedProductOmniPurchaseValues                  *[]AdsActionStats         `json:"converted_promoted_product_omni_purchase_values,omitempty"`
	ConvertedPromotedProductQuantity                            *[]AdsActionStats         `json:"converted_promoted_product_quantity,omitempty"`
	ConvertedPromotedProductValue                               *[]AdsActionStats         `json:"converted_promoted_product_value,omitempty"`
	ConvertedPromotedProductWebsitePixelPurchase                *[]AdsActionStats         `json:"converted_promoted_product_website_pixel_purchase,omitempty"`
	ConvertedPromotedProductWebsitePixelPurchaseValue           *[]AdsActionStats         `json:"converted_promoted_product_website_pixel_purchase_value,omitempty"`
	CostPerX15SecVideoView                                      *[]AdsActionStats         `json:"cost_per_15_sec_video_view,omitempty"`
	CostPerX2SecContinuousVideoView                             *[]AdsActionStats         `json:"cost_per_2_sec_continuous_video_view,omitempty"`
	CostPerX6SecVideoView                                       *[]AdsActionStats         `json:"cost_per_6_sec_video_view,omitempty"`
	CostPerActionType                                           *[]AdsActionStats         `json:"cost_per_action_type,omitempty"`
	CostPerAdClick                                              *[]AdsActionStats         `json:"cost_per_ad_click,omitempty"`
	CostPerConversion                                           *[]AdsActionStats         `json:"cost_per_conversion,omitempty"`
	CostPerConversionLead                                       *[]AdsActionStats         `json:"cost_per_conversion_lead,omitempty"`
	CostPerDdaCountbyConvs                                      *string                   `json:"cost_per_dda_countby_convs,omitempty"`
	CostPerEstimatedAdRecallers                                 *string                   `json:"cost_per_estimated_ad_recallers,omitempty"`
	CostPerInlineLinkClick                                      *string                   `json:"cost_per_inline_link_click,omitempty"`
	CostPerInlinePostEngagement                                 *string                   `json:"cost_per_inline_post_engagement,omitempty"`
	CostPerMessageDelivered                                     *string                   `json:"cost_per_message_delivered,omitempty"`
	CostPerObjectiveResult                                      *[]map[string]interface{} `json:"cost_per_objective_result,omitempty"`
	CostPerOneThousandAdImpression                              *[]AdsActionStats         `json:"cost_per_one_thousand_ad_impression,omitempty"`
	CostPerOutboundClick                                        *[]AdsActionStats         `json:"cost_per_outbound_click,omitempty"`
	CostPerResult                                               *[]map[string]interface{} `json:"cost_per_result,omitempty"`
	CostPerThruplay                                             *[]AdsActionStats         `json:"cost_per_thruplay,omitempty"`
	CostPerUniqueActionType                                     *[]AdsActionStats         `json:"cost_per_unique_action_type,omitempty"`
	CostPerUniqueClick                                          *string                   `json:"cost_per_unique_click,omitempty"`
	CostPerUniqueConversion                                     *[]AdsActionStats         `json:"cost_per_unique_conversion,omitempty"`
	CostPerUniqueInlineLinkClick                                *string                   `json:"cost_per_unique_inline_link_click,omitempty"`
	CostPerUniqueOutboundClick                                  *[]AdsActionStats         `json:"cost_per_unique_outbound_click,omitempty"`
	Cpc                                                         *string                   `json:"cpc,omitempty"`
	Cpm                                                         *string                   `json:"cpm,omitempty"`
	Cpp                                                         *string                   `json:"cpp,omitempty"`
	CreatedTime                                                 *string                   `json:"created_time,omitempty"`
	CreativeDiversityData                                       *[]map[string]interface{} `json:"creative_diversity_data,omitempty"`
	CreativeDiversityLabel                                      *string                   `json:"creative_diversity_label,omitempty"`
	CreativeDiversityScore                                      *string                   `json:"creative_diversity_score,omitempty"`
	CreativeMediaType                                           *string                   `json:"creative_media_type,omitempty"`
	Ctr                                                         *string                   `json:"ctr,omitempty"`
	DateStart                                                   *string                   `json:"date_start,omitempty"`
	DateStop                                                    *string                   `json:"date_stop,omitempty"`
	DdaCountbyConvs                                             *string                   `json:"dda_countby_convs,omitempty"`
	DdaResults                                                  *[]map[string]interface{} `json:"dda_results,omitempty"`
	EngagementRateRanking                                       *string                   `json:"engagement_rate_ranking,omitempty"`
	EstimatedAdRecallRate                                       *string                   `json:"estimated_ad_recall_rate,omitempty"`
	EstimatedAdRecallRateLowerBound                             *string                   `json:"estimated_ad_recall_rate_lower_bound,omitempty"`
	EstimatedAdRecallRateUpperBound                             *string                   `json:"estimated_ad_recall_rate_upper_bound,omitempty"`
	EstimatedAdRecallers                                        *string                   `json:"estimated_ad_recallers,omitempty"`
	EstimatedAdRecallersLowerBound                              *string                   `json:"estimated_ad_recallers_lower_bound,omitempty"`
	EstimatedAdRecallersUpperBound                              *string                   `json:"estimated_ad_recallers_upper_bound,omitempty"`
	Frequency                                                   *string                   `json:"frequency,omitempty"`
	FullViewImpressions                                         *string                   `json:"full_view_impressions,omitempty"`
	FullViewReach                                               *string                   `json:"full_view_reach,omitempty"`
	GenderTargeting                                             *string                   `json:"gender_targeting,omitempty"`
	Impressions                                                 *string                   `json:"impressions,omitempty"`
	InlineLinkClickCtr                                          *string                   `json:"inline_link_click_ctr,omitempty"`
	InlineLinkClicks                                            *string                   `json:"inline_link_clicks,omitempty"`
	InlinePostEngagement                                        *string                   `json:"inline_post_engagement,omitempty"`
	InstagramProfileVisits                                      *string                   `json:"instagram_profile_visits,omitempty"`
	InstagramUpcomingEventRemindersSet                          *string                   `json:"instagram_upcoming_event_reminders_set,omitempty"`
	InstantExperienceClicksToOpen                               *string                   `json:"instant_experience_clicks_to_open,omitempty"`
	InstantExperienceClicksToStart                              *string                   `json:"instant_experience_clicks_to_start,omitempty"`
	InstantExperienceOutboundClicks                             *[]AdsActionStats         `json:"instant_experience_outbound_clicks,omitempty"`
	InteractiveComponentTap                                     *[]AdsActionStats         `json:"interactive_component_tap,omitempty"`
	Labels                                                      *string                   `json:"labels,omitempty"`
	LandingPageViewActionsPerLinkClick                          *string                   `json:"landing_page_view_actions_per_link_click,omitempty"`
	LandingPageViewPerLinkClick                                 *string                   `json:"landing_page_view_per_link_click,omitempty"`
	LandingPageViewPerPurchaseRate                              *string                   `json:"landing_page_view_per_purchase_rate,omitempty"`
	LinkClicksPerResults                                        *[]map[string]interface{} `json:"link_clicks_per_results,omitempty"`
	Location                                                    *string                   `json:"location,omitempty"`
	MarketingMessagesClickRateBenchmark                         *string                   `json:"marketing_messages_click_rate_benchmark,omitempty"`
	MarketingMessagesCostPerDelivered                           *string                   `json:"marketing_messages_cost_per_delivered,omitempty"`
	MarketingMessagesCostPerLinkBtnClick                        *string                   `json:"marketing_messages_cost_per_link_btn_click,omitempty"`
	MarketingMessagesDelivered                                  *string                   `json:"marketing_messages_delivered,omitempty"`
	MarketingMessagesDeliveryRate                               *string                   `json:"marketing_messages_delivery_rate,omitempty"`
	MarketingMessagesLinkBtnClick                               *string                   `json:"marketing_messages_link_btn_click,omitempty"`
	MarketingMessagesLinkBtnClickRate                           *string                   `json:"marketing_messages_link_btn_click_rate,omitempty"`
	MarketingMessagesMediaViewRate                              *string                   `json:"marketing_messages_media_view_rate,omitempty"`
	MarketingMessagesPhoneCallBtnClickRate                      *string                   `json:"marketing_messages_phone_call_btn_click_rate,omitempty"`
	MarketingMessagesQuickReplyBtnClick                         *string                   `json:"marketing_messages_quick_reply_btn_click,omitempty"`
	MarketingMessagesQuickReplyBtnClickRate                     *string                   `json:"marketing_messages_quick_reply_btn_click_rate,omitempty"`
	MarketingMessagesRead                                       *string                   `json:"marketing_messages_read,omitempty"`
	MarketingMessagesReadRate                                   *string                   `json:"marketing_messages_read_rate,omitempty"`
	MarketingMessagesReadRateBenchmark                          *string                   `json:"marketing_messages_read_rate_benchmark,omitempty"`
	MarketingMessagesSent                                       *string                   `json:"marketing_messages_sent,omitempty"`
	MarketingMessagesSpend                                      *string                   `json:"marketing_messages_spend,omitempty"`
	MarketingMessagesSpendCurrency                              *string                   `json:"marketing_messages_spend_currency,omitempty"`
	MarketingMessagesWebsiteAddToCart                           *string                   `json:"marketing_messages_website_add_to_cart,omitempty"`
	MarketingMessagesWebsiteInitiateCheckout                    *string                   `json:"marketing_messages_website_initiate_checkout,omitempty"`
	MarketingMessagesWebsitePurchase                            *string                   `json:"marketing_messages_website_purchase,omitempty"`
	MarketingMessagesWebsitePurchaseValues                      *string                   `json:"marketing_messages_website_purchase_values,omitempty"`
	MessagesDelivered                                           *string                   `json:"messages_delivered,omitempty"`
	MessagesDeliveredCtr                                        *string                   `json:"messages_delivered_ctr,omitempty"`
	MobileAppPurchaseRoas                                       *[]AdsActionStats         `json:"mobile_app_purchase_roas,omitempty"`
	MultiEventConversionAttributionSetting                      *string                   `json:"multi_event_conversion_attribution_setting,omitempty"`
	Objective                                                   *string                   `json:"objective,omitempty"`
	ObjectiveResultRate                                         *[]map[string]interface{} `json:"objective_result_rate,omitempty"`
	ObjectiveResults                                            *[]map[string]interface{} `json:"objective_results,omitempty"`
	OnsiteConversionMessagingDetectedPurchaseDeduped            *[]AdsActionStats         `json:"onsite_conversion_messaging_detected_purchase_deduped,omitempty"`
	OpportunityScoreL4                                          *string                   `json:"opportunity_score_l4,omitempty"`
	OptimizationGoal                                            *string                   `json:"optimization_goal,omitempty"`
	OutboundClicks                                              *[]AdsActionStats         `json:"outbound_clicks,omitempty"`
	OutboundClicksCtr                                           *[]AdsActionStats         `json:"outbound_clicks_ctr,omitempty"`
	PlacePageName                                               *string                   `json:"place_page_name,omitempty"`
	ProductGroupRetailerID                                      *core.ID                  `json:"product_group_retailer_id,omitempty"`
	ProductRetailerID                                           *core.ID                  `json:"product_retailer_id,omitempty"`
	ProductViews                                                *string                   `json:"product_views,omitempty"`
	PurchasePerLandingPageView                                  *string                   `json:"purchase_per_landing_page_view,omitempty"`
	PurchaseRoas                                                *[]AdsActionStats         `json:"purchase_roas,omitempty"`
	PurchasesPerLinkClick                                       *string                   `json:"purchases_per_link_click,omitempty"`
	QualifyingQuestionQualifyAnswerRate                         *string                   `json:"qualifying_question_qualify_answer_rate,omitempty"`
	QualityRanking                                              *string                   `json:"quality_ranking,omitempty"`
	Reach                                                       *string                   `json:"reach,omitempty"`
	ReadRate                                                    *string                   `json:"read_rate,omitempty"`
	ResultRate                                                  *[]map[string]interface{} `json:"result_rate,omitempty"`
	ResultValuesPerformanceIndicator                            *string                   `json:"result_values_performance_indicator,omitempty"`
	Results                                                     *[]map[string]interface{} `json:"results,omitempty"`
	ShopClicks                                                  *string                   `json:"shop_clicks,omitempty"`
	ShopsAssistedPurchases                                      *string                   `json:"shops_assisted_purchases,omitempty"`
	SocialSpend                                                 *string                   `json:"social_spend,omitempty"`
	Spend                                                       *string                   `json:"spend,omitempty"`
	TotalCardView                                               *string                   `json:"total_card_view,omitempty"`
	TotalPostbacks                                              *string                   `json:"total_postbacks,omitempty"`
	TotalPostbacksDetailed                                      *[]AdsActionStats         `json:"total_postbacks_detailed,omitempty"`
	TotalPostbacksDetailedV4                                    *[]AdsActionStats         `json:"total_postbacks_detailed_v4,omitempty"`
	UniqueActions                                               *[]AdsActionStats         `json:"unique_actions,omitempty"`
	UniqueClicks                                                *string                   `json:"unique_clicks,omitempty"`
	UniqueConversions                                           *[]AdsActionStats         `json:"unique_conversions,omitempty"`
	UniqueCtr                                                   *string                   `json:"unique_ctr,omitempty"`
	UniqueInlineLinkClickCtr                                    *string                   `json:"unique_inline_link_click_ctr,omitempty"`
	UniqueInlineLinkClicks                                      *string                   `json:"unique_inline_link_clicks,omitempty"`
	UniqueLinkClicksCtr                                         *string                   `json:"unique_link_clicks_ctr,omitempty"`
	UniqueOutboundClicks                                        *[]AdsActionStats         `json:"unique_outbound_clicks,omitempty"`
	UniqueOutboundClicksCtr                                     *[]AdsActionStats         `json:"unique_outbound_clicks_ctr,omitempty"`
	UniqueVideoContinuousX2SecWatchedActions                    *[]AdsActionStats         `json:"unique_video_continuous_2_sec_watched_actions,omitempty"`
	UniqueVideoViewX15Sec                                       *[]AdsActionStats         `json:"unique_video_view_15_sec,omitempty"`
	UpdatedTime                                                 *string                   `json:"updated_time,omitempty"`
	VideoX15SecWatchedActions                                   *[]AdsActionStats         `json:"video_15_sec_watched_actions,omitempty"`
	VideoX30SecWatchedActions                                   *[]AdsActionStats         `json:"video_30_sec_watched_actions,omitempty"`
	VideoX6SecWatchedActions                                    *[]AdsActionStats         `json:"video_6_sec_watched_actions,omitempty"`
	VideoAvgTimeWatchedActions                                  *[]AdsActionStats         `json:"video_avg_time_watched_actions,omitempty"`
	VideoContinuousX2SecWatchedActions                          *[]AdsActionStats         `json:"video_continuous_2_sec_watched_actions,omitempty"`
	VideoP100WatchedActions                                     *[]AdsActionStats         `json:"video_p100_watched_actions,omitempty"`
	VideoP25WatchedActions                                      *[]AdsActionStats         `json:"video_p25_watched_actions,omitempty"`
	VideoP50WatchedActions                                      *[]AdsActionStats         `json:"video_p50_watched_actions,omitempty"`
	VideoP75WatchedActions                                      *[]AdsActionStats         `json:"video_p75_watched_actions,omitempty"`
	VideoP95WatchedActions                                      *[]AdsActionStats         `json:"video_p95_watched_actions,omitempty"`
	VideoPlayActions                                            *[]AdsActionStats         `json:"video_play_actions,omitempty"`
	VideoPlayCurveActions                                       *[]AdsHistogramStats      `json:"video_play_curve_actions,omitempty"`
	VideoPlayRetentionX0ToX15sActions                           *[]AdsHistogramStats      `json:"video_play_retention_0_to_15s_actions,omitempty"`
	VideoPlayRetentionX20ToX60sActions                          *[]AdsHistogramStats      `json:"video_play_retention_20_to_60s_actions,omitempty"`
	VideoPlayRetentionGraphActions                              *[]AdsHistogramStats      `json:"video_play_retention_graph_actions,omitempty"`
	VideoThruplayWatchedActions                                 *[]AdsActionStats         `json:"video_thruplay_watched_actions,omitempty"`
	VideoTimeWatchedActions                                     *[]AdsActionStats         `json:"video_time_watched_actions,omitempty"`
	VideoViewPerImpression                                      *[]AdsActionStats         `json:"video_view_per_impression,omitempty"`
	WebsiteCtr                                                  *[]AdsActionStats         `json:"website_ctr,omitempty"`
	WebsitePurchaseRoas                                         *[]AdsActionStats         `json:"website_purchase_roas,omitempty"`
	WishBid                                                     *string                   `json:"wish_bid,omitempty"`
}

var AdsInsightsFields = struct {
	AccountCurrency                                             string
	AccountID                                                   string
	AccountName                                                 string
	ActionValues                                                string
	Actions                                                     string
	AdClickActions                                              string
	AdID                                                        string
	AdImpressionActions                                         string
	AdName                                                      string
	AdsetEnd                                                    string
	AdsetID                                                     string
	AdsetName                                                   string
	AdsetStart                                                  string
	AdvancedActionsX28dView                                     string
	AdvancedReachX1dLookback                                    string
	AdvancedReachX28dLookback                                   string
	AdvancedReachX7dLookback                                    string
	AgeTargeting                                                string
	AnchorEventAttributionSetting                               string
	AnchorEventsPerformanceIndicator                            string
	AttributionSetting                                          string
	AuctionBid                                                  string
	AuctionCompetitiveness                                      string
	AuctionMaxCompetitorBid                                     string
	AveragePurchasesConversionValue                             string
	BuyingType                                                  string
	CampaignID                                                  string
	CampaignName                                                string
	CanvasAvgViewPercent                                        string
	CanvasAvgViewTime                                           string
	CatalogSegmentActions                                       string
	CatalogSegmentValue                                         string
	CatalogSegmentValueMobilePurchaseRoas                       string
	CatalogSegmentValueOmniPurchaseRoas                         string
	CatalogSegmentValueWebsitePurchaseRoas                      string
	Clicks                                                      string
	ConfigurableAttributionAction                               string
	ConfigurableAttributionActionvalue                          string
	ConfigurableAudienceOverlapReach                            string
	ConfigurableReachbyfrequencyAction                          string
	ConfigurableReachbyfrequencyConvertersCount                 string
	ConfigurableReachbyfrequencyImpressionsCost                 string
	ConfigurableReachbyfrequencyImpressionsCount                string
	ConfigurableReachbyfrequencyReach                           string
	ConversionLeadRate                                          string
	ConversionLeads                                             string
	ConversionRateRanking                                       string
	ConversionValues                                            string
	Conversions                                                 string
	ConvertedProductAppCustomEventFbMobilePurchase              string
	ConvertedProductAppCustomEventFbMobilePurchaseValue         string
	ConvertedProductOfflinePurchase                             string
	ConvertedProductOfflinePurchaseValue                        string
	ConvertedProductOmniPurchase                                string
	ConvertedProductOmniPurchaseValues                          string
	ConvertedProductQuantity                                    string
	ConvertedProductValue                                       string
	ConvertedProductWebsitePixelPurchase                        string
	ConvertedProductWebsitePixelPurchaseValue                   string
	ConvertedPromotedProductAppCustomEventFbMobilePurchase      string
	ConvertedPromotedProductAppCustomEventFbMobilePurchaseValue string
	ConvertedPromotedProductOfflinePurchase                     string
	ConvertedPromotedProductOfflinePurchaseValue                string
	ConvertedPromotedProductOmniPurchase                        string
	ConvertedPromotedProductOmniPurchaseValues                  string
	ConvertedPromotedProductQuantity                            string
	ConvertedPromotedProductValue                               string
	ConvertedPromotedProductWebsitePixelPurchase                string
	ConvertedPromotedProductWebsitePixelPurchaseValue           string
	CostPerX15SecVideoView                                      string
	CostPerX2SecContinuousVideoView                             string
	CostPerX6SecVideoView                                       string
	CostPerActionType                                           string
	CostPerAdClick                                              string
	CostPerConversion                                           string
	CostPerConversionLead                                       string
	CostPerDdaCountbyConvs                                      string
	CostPerEstimatedAdRecallers                                 string
	CostPerInlineLinkClick                                      string
	CostPerInlinePostEngagement                                 string
	CostPerMessageDelivered                                     string
	CostPerObjectiveResult                                      string
	CostPerOneThousandAdImpression                              string
	CostPerOutboundClick                                        string
	CostPerResult                                               string
	CostPerThruplay                                             string
	CostPerUniqueActionType                                     string
	CostPerUniqueClick                                          string
	CostPerUniqueConversion                                     string
	CostPerUniqueInlineLinkClick                                string
	CostPerUniqueOutboundClick                                  string
	Cpc                                                         string
	Cpm                                                         string
	Cpp                                                         string
	CreatedTime                                                 string
	CreativeDiversityData                                       string
	CreativeDiversityLabel                                      string
	CreativeDiversityScore                                      string
	CreativeMediaType                                           string
	Ctr                                                         string
	DateStart                                                   string
	DateStop                                                    string
	DdaCountbyConvs                                             string
	DdaResults                                                  string
	EngagementRateRanking                                       string
	EstimatedAdRecallRate                                       string
	EstimatedAdRecallRateLowerBound                             string
	EstimatedAdRecallRateUpperBound                             string
	EstimatedAdRecallers                                        string
	EstimatedAdRecallersLowerBound                              string
	EstimatedAdRecallersUpperBound                              string
	Frequency                                                   string
	FullViewImpressions                                         string
	FullViewReach                                               string
	GenderTargeting                                             string
	Impressions                                                 string
	InlineLinkClickCtr                                          string
	InlineLinkClicks                                            string
	InlinePostEngagement                                        string
	InstagramProfileVisits                                      string
	InstagramUpcomingEventRemindersSet                          string
	InstantExperienceClicksToOpen                               string
	InstantExperienceClicksToStart                              string
	InstantExperienceOutboundClicks                             string
	InteractiveComponentTap                                     string
	Labels                                                      string
	LandingPageViewActionsPerLinkClick                          string
	LandingPageViewPerLinkClick                                 string
	LandingPageViewPerPurchaseRate                              string
	LinkClicksPerResults                                        string
	Location                                                    string
	MarketingMessagesClickRateBenchmark                         string
	MarketingMessagesCostPerDelivered                           string
	MarketingMessagesCostPerLinkBtnClick                        string
	MarketingMessagesDelivered                                  string
	MarketingMessagesDeliveryRate                               string
	MarketingMessagesLinkBtnClick                               string
	MarketingMessagesLinkBtnClickRate                           string
	MarketingMessagesMediaViewRate                              string
	MarketingMessagesPhoneCallBtnClickRate                      string
	MarketingMessagesQuickReplyBtnClick                         string
	MarketingMessagesQuickReplyBtnClickRate                     string
	MarketingMessagesRead                                       string
	MarketingMessagesReadRate                                   string
	MarketingMessagesReadRateBenchmark                          string
	MarketingMessagesSent                                       string
	MarketingMessagesSpend                                      string
	MarketingMessagesSpendCurrency                              string
	MarketingMessagesWebsiteAddToCart                           string
	MarketingMessagesWebsiteInitiateCheckout                    string
	MarketingMessagesWebsitePurchase                            string
	MarketingMessagesWebsitePurchaseValues                      string
	MessagesDelivered                                           string
	MessagesDeliveredCtr                                        string
	MobileAppPurchaseRoas                                       string
	MultiEventConversionAttributionSetting                      string
	Objective                                                   string
	ObjectiveResultRate                                         string
	ObjectiveResults                                            string
	OnsiteConversionMessagingDetectedPurchaseDeduped            string
	OpportunityScoreL4                                          string
	OptimizationGoal                                            string
	OutboundClicks                                              string
	OutboundClicksCtr                                           string
	PlacePageName                                               string
	ProductGroupRetailerID                                      string
	ProductRetailerID                                           string
	ProductViews                                                string
	PurchasePerLandingPageView                                  string
	PurchaseRoas                                                string
	PurchasesPerLinkClick                                       string
	QualifyingQuestionQualifyAnswerRate                         string
	QualityRanking                                              string
	Reach                                                       string
	ReadRate                                                    string
	ResultRate                                                  string
	ResultValuesPerformanceIndicator                            string
	Results                                                     string
	ShopClicks                                                  string
	ShopsAssistedPurchases                                      string
	SocialSpend                                                 string
	Spend                                                       string
	TotalCardView                                               string
	TotalPostbacks                                              string
	TotalPostbacksDetailed                                      string
	TotalPostbacksDetailedV4                                    string
	UniqueActions                                               string
	UniqueClicks                                                string
	UniqueConversions                                           string
	UniqueCtr                                                   string
	UniqueInlineLinkClickCtr                                    string
	UniqueInlineLinkClicks                                      string
	UniqueLinkClicksCtr                                         string
	UniqueOutboundClicks                                        string
	UniqueOutboundClicksCtr                                     string
	UniqueVideoContinuousX2SecWatchedActions                    string
	UniqueVideoViewX15Sec                                       string
	UpdatedTime                                                 string
	VideoX15SecWatchedActions                                   string
	VideoX30SecWatchedActions                                   string
	VideoX6SecWatchedActions                                    string
	VideoAvgTimeWatchedActions                                  string
	VideoContinuousX2SecWatchedActions                          string
	VideoP100WatchedActions                                     string
	VideoP25WatchedActions                                      string
	VideoP50WatchedActions                                      string
	VideoP75WatchedActions                                      string
	VideoP95WatchedActions                                      string
	VideoPlayActions                                            string
	VideoPlayCurveActions                                       string
	VideoPlayRetentionX0ToX15sActions                           string
	VideoPlayRetentionX20ToX60sActions                          string
	VideoPlayRetentionGraphActions                              string
	VideoThruplayWatchedActions                                 string
	VideoTimeWatchedActions                                     string
	VideoViewPerImpression                                      string
	WebsiteCtr                                                  string
	WebsitePurchaseRoas                                         string
	WishBid                                                     string
}{
	AccountCurrency:                        "account_currency",
	AccountID:                              "account_id",
	AccountName:                            "account_name",
	ActionValues:                           "action_values",
	Actions:                                "actions",
	AdClickActions:                         "ad_click_actions",
	AdID:                                   "ad_id",
	AdImpressionActions:                    "ad_impression_actions",
	AdName:                                 "ad_name",
	AdsetEnd:                               "adset_end",
	AdsetID:                                "adset_id",
	AdsetName:                              "adset_name",
	AdsetStart:                             "adset_start",
	AdvancedActionsX28dView:                "advanced_actions_28d_view",
	AdvancedReachX1dLookback:               "advanced_reach_1d_lookback",
	AdvancedReachX28dLookback:              "advanced_reach_28d_lookback",
	AdvancedReachX7dLookback:               "advanced_reach_7d_lookback",
	AgeTargeting:                           "age_targeting",
	AnchorEventAttributionSetting:          "anchor_event_attribution_setting",
	AnchorEventsPerformanceIndicator:       "anchor_events_performance_indicator",
	AttributionSetting:                     "attribution_setting",
	AuctionBid:                             "auction_bid",
	AuctionCompetitiveness:                 "auction_competitiveness",
	AuctionMaxCompetitorBid:                "auction_max_competitor_bid",
	AveragePurchasesConversionValue:        "average_purchases_conversion_value",
	BuyingType:                             "buying_type",
	CampaignID:                             "campaign_id",
	CampaignName:                           "campaign_name",
	CanvasAvgViewPercent:                   "canvas_avg_view_percent",
	CanvasAvgViewTime:                      "canvas_avg_view_time",
	CatalogSegmentActions:                  "catalog_segment_actions",
	CatalogSegmentValue:                    "catalog_segment_value",
	CatalogSegmentValueMobilePurchaseRoas:  "catalog_segment_value_mobile_purchase_roas",
	CatalogSegmentValueOmniPurchaseRoas:    "catalog_segment_value_omni_purchase_roas",
	CatalogSegmentValueWebsitePurchaseRoas: "catalog_segment_value_website_purchase_roas",
	Clicks:                                 "clicks",
	ConfigurableAttributionAction:          "configurable_attribution_action",
	ConfigurableAttributionActionvalue:     "configurable_attribution_actionvalue",
	ConfigurableAudienceOverlapReach:       "configurable_audience_overlap_reach",
	ConfigurableReachbyfrequencyAction:     "configurable_reachbyfrequency_action",
	ConfigurableReachbyfrequencyConvertersCount:                 "configurable_reachbyfrequency_converters_count",
	ConfigurableReachbyfrequencyImpressionsCost:                 "configurable_reachbyfrequency_impressions_cost",
	ConfigurableReachbyfrequencyImpressionsCount:                "configurable_reachbyfrequency_impressions_count",
	ConfigurableReachbyfrequencyReach:                           "configurable_reachbyfrequency_reach",
	ConversionLeadRate:                                          "conversion_lead_rate",
	ConversionLeads:                                             "conversion_leads",
	ConversionRateRanking:                                       "conversion_rate_ranking",
	ConversionValues:                                            "conversion_values",
	Conversions:                                                 "conversions",
	ConvertedProductAppCustomEventFbMobilePurchase:              "converted_product_app_custom_event_fb_mobile_purchase",
	ConvertedProductAppCustomEventFbMobilePurchaseValue:         "converted_product_app_custom_event_fb_mobile_purchase_value",
	ConvertedProductOfflinePurchase:                             "converted_product_offline_purchase",
	ConvertedProductOfflinePurchaseValue:                        "converted_product_offline_purchase_value",
	ConvertedProductOmniPurchase:                                "converted_product_omni_purchase",
	ConvertedProductOmniPurchaseValues:                          "converted_product_omni_purchase_values",
	ConvertedProductQuantity:                                    "converted_product_quantity",
	ConvertedProductValue:                                       "converted_product_value",
	ConvertedProductWebsitePixelPurchase:                        "converted_product_website_pixel_purchase",
	ConvertedProductWebsitePixelPurchaseValue:                   "converted_product_website_pixel_purchase_value",
	ConvertedPromotedProductAppCustomEventFbMobilePurchase:      "converted_promoted_product_app_custom_event_fb_mobile_purchase",
	ConvertedPromotedProductAppCustomEventFbMobilePurchaseValue: "converted_promoted_product_app_custom_event_fb_mobile_purchase_value",
	ConvertedPromotedProductOfflinePurchase:                     "converted_promoted_product_offline_purchase",
	ConvertedPromotedProductOfflinePurchaseValue:                "converted_promoted_product_offline_purchase_value",
	ConvertedPromotedProductOmniPurchase:                        "converted_promoted_product_omni_purchase",
	ConvertedPromotedProductOmniPurchaseValues:                  "converted_promoted_product_omni_purchase_values",
	ConvertedPromotedProductQuantity:                            "converted_promoted_product_quantity",
	ConvertedPromotedProductValue:                               "converted_promoted_product_value",
	ConvertedPromotedProductWebsitePixelPurchase:                "converted_promoted_product_website_pixel_purchase",
	ConvertedPromotedProductWebsitePixelPurchaseValue:           "converted_promoted_product_website_pixel_purchase_value",
	CostPerX15SecVideoView:                                      "cost_per_15_sec_video_view",
	CostPerX2SecContinuousVideoView:                             "cost_per_2_sec_continuous_video_view",
	CostPerX6SecVideoView:                                       "cost_per_6_sec_video_view",
	CostPerActionType:                                           "cost_per_action_type",
	CostPerAdClick:                                              "cost_per_ad_click",
	CostPerConversion:                                           "cost_per_conversion",
	CostPerConversionLead:                                       "cost_per_conversion_lead",
	CostPerDdaCountbyConvs:                                      "cost_per_dda_countby_convs",
	CostPerEstimatedAdRecallers:                                 "cost_per_estimated_ad_recallers",
	CostPerInlineLinkClick:                                      "cost_per_inline_link_click",
	CostPerInlinePostEngagement:                                 "cost_per_inline_post_engagement",
	CostPerMessageDelivered:                                     "cost_per_message_delivered",
	CostPerObjectiveResult:                                      "cost_per_objective_result",
	CostPerOneThousandAdImpression:                              "cost_per_one_thousand_ad_impression",
	CostPerOutboundClick:                                        "cost_per_outbound_click",
	CostPerResult:                                               "cost_per_result",
	CostPerThruplay:                                             "cost_per_thruplay",
	CostPerUniqueActionType:                                     "cost_per_unique_action_type",
	CostPerUniqueClick:                                          "cost_per_unique_click",
	CostPerUniqueConversion:                                     "cost_per_unique_conversion",
	CostPerUniqueInlineLinkClick:                                "cost_per_unique_inline_link_click",
	CostPerUniqueOutboundClick:                                  "cost_per_unique_outbound_click",
	Cpc:                                                         "cpc",
	Cpm:                                                         "cpm",
	Cpp:                                                         "cpp",
	CreatedTime:                                                 "created_time",
	CreativeDiversityData:                                       "creative_diversity_data",
	CreativeDiversityLabel:                                      "creative_diversity_label",
	CreativeDiversityScore:                                      "creative_diversity_score",
	CreativeMediaType:                                           "creative_media_type",
	Ctr:                                                         "ctr",
	DateStart:                                                   "date_start",
	DateStop:                                                    "date_stop",
	DdaCountbyConvs:                                             "dda_countby_convs",
	DdaResults:                                                  "dda_results",
	EngagementRateRanking:                                       "engagement_rate_ranking",
	EstimatedAdRecallRate:                                       "estimated_ad_recall_rate",
	EstimatedAdRecallRateLowerBound:                             "estimated_ad_recall_rate_lower_bound",
	EstimatedAdRecallRateUpperBound:                             "estimated_ad_recall_rate_upper_bound",
	EstimatedAdRecallers:                                        "estimated_ad_recallers",
	EstimatedAdRecallersLowerBound:                              "estimated_ad_recallers_lower_bound",
	EstimatedAdRecallersUpperBound:                              "estimated_ad_recallers_upper_bound",
	Frequency:                                                   "frequency",
	FullViewImpressions:                                         "full_view_impressions",
	FullViewReach:                                               "full_view_reach",
	GenderTargeting:                                             "gender_targeting",
	Impressions:                                                 "impressions",
	InlineLinkClickCtr:                                          "inline_link_click_ctr",
	InlineLinkClicks:                                            "inline_link_clicks",
	InlinePostEngagement:                                        "inline_post_engagement",
	InstagramProfileVisits:                                      "instagram_profile_visits",
	InstagramUpcomingEventRemindersSet:                          "instagram_upcoming_event_reminders_set",
	InstantExperienceClicksToOpen:                               "instant_experience_clicks_to_open",
	InstantExperienceClicksToStart:                              "instant_experience_clicks_to_start",
	InstantExperienceOutboundClicks:                             "instant_experience_outbound_clicks",
	InteractiveComponentTap:                                     "interactive_component_tap",
	Labels:                                                      "labels",
	LandingPageViewActionsPerLinkClick:                          "landing_page_view_actions_per_link_click",
	LandingPageViewPerLinkClick:                                 "landing_page_view_per_link_click",
	LandingPageViewPerPurchaseRate:                              "landing_page_view_per_purchase_rate",
	LinkClicksPerResults:                                        "link_clicks_per_results",
	Location:                                                    "location",
	MarketingMessagesClickRateBenchmark:                         "marketing_messages_click_rate_benchmark",
	MarketingMessagesCostPerDelivered:                           "marketing_messages_cost_per_delivered",
	MarketingMessagesCostPerLinkBtnClick:                        "marketing_messages_cost_per_link_btn_click",
	MarketingMessagesDelivered:                                  "marketing_messages_delivered",
	MarketingMessagesDeliveryRate:                               "marketing_messages_delivery_rate",
	MarketingMessagesLinkBtnClick:                               "marketing_messages_link_btn_click",
	MarketingMessagesLinkBtnClickRate:                           "marketing_messages_link_btn_click_rate",
	MarketingMessagesMediaViewRate:                              "marketing_messages_media_view_rate",
	MarketingMessagesPhoneCallBtnClickRate:                      "marketing_messages_phone_call_btn_click_rate",
	MarketingMessagesQuickReplyBtnClick:                         "marketing_messages_quick_reply_btn_click",
	MarketingMessagesQuickReplyBtnClickRate:                     "marketing_messages_quick_reply_btn_click_rate",
	MarketingMessagesRead:                                       "marketing_messages_read",
	MarketingMessagesReadRate:                                   "marketing_messages_read_rate",
	MarketingMessagesReadRateBenchmark:                          "marketing_messages_read_rate_benchmark",
	MarketingMessagesSent:                                       "marketing_messages_sent",
	MarketingMessagesSpend:                                      "marketing_messages_spend",
	MarketingMessagesSpendCurrency:                              "marketing_messages_spend_currency",
	MarketingMessagesWebsiteAddToCart:                           "marketing_messages_website_add_to_cart",
	MarketingMessagesWebsiteInitiateCheckout:                    "marketing_messages_website_initiate_checkout",
	MarketingMessagesWebsitePurchase:                            "marketing_messages_website_purchase",
	MarketingMessagesWebsitePurchaseValues:                      "marketing_messages_website_purchase_values",
	MessagesDelivered:                                           "messages_delivered",
	MessagesDeliveredCtr:                                        "messages_delivered_ctr",
	MobileAppPurchaseRoas:                                       "mobile_app_purchase_roas",
	MultiEventConversionAttributionSetting:                      "multi_event_conversion_attribution_setting",
	Objective:                                                   "objective",
	ObjectiveResultRate:                                         "objective_result_rate",
	ObjectiveResults:                                            "objective_results",
	OnsiteConversionMessagingDetectedPurchaseDeduped:            "onsite_conversion_messaging_detected_purchase_deduped",
	OpportunityScoreL4:                                          "opportunity_score_l4",
	OptimizationGoal:                                            "optimization_goal",
	OutboundClicks:                                              "outbound_clicks",
	OutboundClicksCtr:                                           "outbound_clicks_ctr",
	PlacePageName:                                               "place_page_name",
	ProductGroupRetailerID:                                      "product_group_retailer_id",
	ProductRetailerID:                                           "product_retailer_id",
	ProductViews:                                                "product_views",
	PurchasePerLandingPageView:                                  "purchase_per_landing_page_view",
	PurchaseRoas:                                                "purchase_roas",
	PurchasesPerLinkClick:                                       "purchases_per_link_click",
	QualifyingQuestionQualifyAnswerRate:                         "qualifying_question_qualify_answer_rate",
	QualityRanking:                                              "quality_ranking",
	Reach:                                                       "reach",
	ReadRate:                                                    "read_rate",
	ResultRate:                                                  "result_rate",
	ResultValuesPerformanceIndicator:                            "result_values_performance_indicator",
	Results:                                                     "results",
	ShopClicks:                                                  "shop_clicks",
	ShopsAssistedPurchases:                                      "shops_assisted_purchases",
	SocialSpend:                                                 "social_spend",
	Spend:                                                       "spend",
	TotalCardView:                                               "total_card_view",
	TotalPostbacks:                                              "total_postbacks",
	TotalPostbacksDetailed:                                      "total_postbacks_detailed",
	TotalPostbacksDetailedV4:                                    "total_postbacks_detailed_v4",
	UniqueActions:                                               "unique_actions",
	UniqueClicks:                                                "unique_clicks",
	UniqueConversions:                                           "unique_conversions",
	UniqueCtr:                                                   "unique_ctr",
	UniqueInlineLinkClickCtr:                                    "unique_inline_link_click_ctr",
	UniqueInlineLinkClicks:                                      "unique_inline_link_clicks",
	UniqueLinkClicksCtr:                                         "unique_link_clicks_ctr",
	UniqueOutboundClicks:                                        "unique_outbound_clicks",
	UniqueOutboundClicksCtr:                                     "unique_outbound_clicks_ctr",
	UniqueVideoContinuousX2SecWatchedActions:                    "unique_video_continuous_2_sec_watched_actions",
	UniqueVideoViewX15Sec:                                       "unique_video_view_15_sec",
	UpdatedTime:                                                 "updated_time",
	VideoX15SecWatchedActions:                                   "video_15_sec_watched_actions",
	VideoX30SecWatchedActions:                                   "video_30_sec_watched_actions",
	VideoX6SecWatchedActions:                                    "video_6_sec_watched_actions",
	VideoAvgTimeWatchedActions:                                  "video_avg_time_watched_actions",
	VideoContinuousX2SecWatchedActions:                          "video_continuous_2_sec_watched_actions",
	VideoP100WatchedActions:                                     "video_p100_watched_actions",
	VideoP25WatchedActions:                                      "video_p25_watched_actions",
	VideoP50WatchedActions:                                      "video_p50_watched_actions",
	VideoP75WatchedActions:                                      "video_p75_watched_actions",
	VideoP95WatchedActions:                                      "video_p95_watched_actions",
	VideoPlayActions:                                            "video_play_actions",
	VideoPlayCurveActions:                                       "video_play_curve_actions",
	VideoPlayRetentionX0ToX15sActions:                           "video_play_retention_0_to_15s_actions",
	VideoPlayRetentionX20ToX60sActions:                          "video_play_retention_20_to_60s_actions",
	VideoPlayRetentionGraphActions:                              "video_play_retention_graph_actions",
	VideoThruplayWatchedActions:                                 "video_thruplay_watched_actions",
	VideoTimeWatchedActions:                                     "video_time_watched_actions",
	VideoViewPerImpression:                                      "video_view_per_impression",
	WebsiteCtr:                                                  "website_ctr",
	WebsitePurchaseRoas:                                         "website_purchase_roas",
	WishBid:                                                     "wish_bid",
}
