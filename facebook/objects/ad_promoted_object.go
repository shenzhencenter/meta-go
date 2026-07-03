package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AdPromotedObject struct {
	ApplicationID                 *core.ID                                      `json:"application_id,omitempty"`
	BoostedProductSetID           *core.ID                                      `json:"boosted_product_set_id,omitempty"`
	ConversionGoalID              *core.ID                                      `json:"conversion_goal_id,omitempty"`
	CustomAttributionSourceIds    *[]core.ID                                    `json:"custom_attribution_source_ids,omitempty"`
	CustomConversionID            *core.ID                                      `json:"custom_conversion_id,omitempty"`
	CustomEventStr                *string                                       `json:"custom_event_str,omitempty"`
	CustomEventType               *enums.AdpromotedobjectCustomEventType        `json:"custom_event_type,omitempty"`
	DatasetSplitID                *core.ID                                      `json:"dataset_split_id,omitempty"`
	DatasetSplitIds               *[]core.ID                                    `json:"dataset_split_ids,omitempty"`
	EventID                       *core.ID                                      `json:"event_id,omitempty"`
	FullFunnelObjective           *enums.AdpromotedobjectFullFunnelObjective    `json:"full_funnel_objective,omitempty"`
	FundraiserCampaignID          *core.ID                                      `json:"fundraiser_campaign_id,omitempty"`
	InstagramActorID              *core.ID                                      `json:"instagram_actor_id,omitempty"`
	JobListingID                  *core.ID                                      `json:"job_listing_id,omitempty"`
	LeadAdsCustomEventStr         *string                                       `json:"lead_ads_custom_event_str,omitempty"`
	LeadAdsCustomEventType        *enums.AdpromotedobjectLeadAdsCustomEventType `json:"lead_ads_custom_event_type,omitempty"`
	LeadAdsFollowUpEvent          *string                                       `json:"lead_ads_follow_up_event,omitempty"`
	LeadAdsFormEventSourceType    *string                                       `json:"lead_ads_form_event_source_type,omitempty"`
	LeadAdsOffsiteConversionType  *string                                       `json:"lead_ads_offsite_conversion_type,omitempty"`
	LeadAdsQualityVolumeSetting   *string                                       `json:"lead_ads_quality_volume_setting,omitempty"`
	LeadAdsSelectedPixelID        *core.ID                                      `json:"lead_ads_selected_pixel_id,omitempty"`
	LiveVideoDestination          *string                                       `json:"live_video_destination,omitempty"`
	McmeConversionID              *core.ID                                      `json:"mcme_conversion_id,omitempty"`
	MultiEventProduct             *int                                          `json:"multi_event_product,omitempty"`
	ObjectStoreURL                *string                                       `json:"object_store_url,omitempty"`
	ObjectStoreUrls               *[]string                                     `json:"object_store_urls,omitempty"`
	OfferID                       *core.ID                                      `json:"offer_id,omitempty"`
	OfflineConversionDataSetID    *core.ID                                      `json:"offline_conversion_data_set_id,omitempty"`
	OffsiteConversionEventID      *core.ID                                      `json:"offsite_conversion_event_id,omitempty"`
	OmnichannelObject             *map[string]interface{}                       `json:"omnichannel_object,omitempty"`
	PageID                        *core.ID                                      `json:"page_id,omitempty"`
	PassbackApplicationID         *core.ID                                      `json:"passback_application_id,omitempty"`
	PassbackPixelID               *core.ID                                      `json:"passback_pixel_id,omitempty"`
	PixelAggregationRule          *string                                       `json:"pixel_aggregation_rule,omitempty"`
	PixelID                       *core.ID                                      `json:"pixel_id,omitempty"`
	PixelRule                     *string                                       `json:"pixel_rule,omitempty"`
	PlacePageSet                  *AdPlacePageSet                               `json:"place_page_set,omitempty"`
	PlacePageSetID                *core.ID                                      `json:"place_page_set_id,omitempty"`
	ProductCatalogID              *core.ID                                      `json:"product_catalog_id,omitempty"`
	ProductItemID                 *core.ID                                      `json:"product_item_id,omitempty"`
	ProductSet                    *ProductSet                                   `json:"product_set,omitempty"`
	ProductSetID                  *core.ID                                      `json:"product_set_id,omitempty"`
	ProductSetOptimization        *string                                       `json:"product_set_optimization,omitempty"`
	RetentionDays                 *string                                       `json:"retention_days,omitempty"`
	SmartPseEnabled               *bool                                         `json:"smart_pse_enabled,omitempty"`
	SmartPseSetting               *string                                       `json:"smart_pse_setting,omitempty"`
	ValueSemanticType             *string                                       `json:"value_semantic_type,omitempty"`
	Variation                     *string                                       `json:"variation,omitempty"`
	WhatsAppBusinessPhoneNumberID *core.ID                                      `json:"whats_app_business_phone_number_id,omitempty"`
	WhatsappPhoneNumber           *string                                       `json:"whatsapp_phone_number,omitempty"`
}
