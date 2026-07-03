package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
)

type AdAssetFeedSpec struct {
	AdFormats               *[]string                                 `json:"ad_formats,omitempty"`
	AdditionalData          *AdAssetFeedAdditionalData                `json:"additional_data,omitempty"`
	AppProductPageID        *core.ID                                  `json:"app_product_page_id,omitempty"`
	AssetCustomizationRules *[]AdAssetFeedSpecAssetCustomizationRule  `json:"asset_customization_rules,omitempty"`
	Audios                  *[]map[string]interface{}                 `json:"audios,omitempty"`
	Autotranslate           *[]string                                 `json:"autotranslate,omitempty"`
	Bodies                  *[]AdAssetFeedSpecBody                    `json:"bodies,omitempty"`
	CallAdsConfiguration    *map[string]interface{}                   `json:"call_ads_configuration,omitempty"`
	CallToActionTypes       *[]enums.AdassetfeedspecCallToActionTypes `json:"call_to_action_types,omitempty"`
	CallToActions           *[]AdAssetFeedSpecCallToAction            `json:"call_to_actions,omitempty"`
	Captions                *[]AdAssetFeedSpecCaption                 `json:"captions,omitempty"`
	Carousels               *[]AdAssetFeedSpecCarousel                `json:"carousels,omitempty"`
	CtwaConsentData         *[]map[string]interface{}                 `json:"ctwa_consent_data,omitempty"`
	Descriptions            *[]AdAssetFeedSpecDescription             `json:"descriptions,omitempty"`
	Events                  *[]AdAssetFeedSpecEvents                  `json:"events,omitempty"`
	Groups                  *[]AdAssetFeedSpecGroupRule               `json:"groups,omitempty"`
	Images                  *[]AdAssetFeedSpecImage                   `json:"images,omitempty"`
	LinkUrls                *[]AdAssetFeedSpecLinkURL                 `json:"link_urls,omitempty"`
	MessageExtensions       *[]AdAssetMessageExtensions               `json:"message_extensions,omitempty"`
	OnsiteDestinations      *[]AdAssetOnsiteDestinations              `json:"onsite_destinations,omitempty"`
	OptimizationType        *string                                   `json:"optimization_type,omitempty"`
	PromotionalMetadata     *map[string]interface{}                   `json:"promotional_metadata,omitempty"`
	ReasonsToShop           *bool                                     `json:"reasons_to_shop,omitempty"`
	ShopsBundle             *bool                                     `json:"shops_bundle,omitempty"`
	Titles                  *[]AdAssetFeedSpecTitle                   `json:"titles,omitempty"`
	Translations            *[]map[string]interface{}                 `json:"translations,omitempty"`
	UpcomingEvents          *[]map[string]interface{}                 `json:"upcoming_events,omitempty"`
	Videos                  *[]AdAssetFeedSpecVideo                   `json:"videos,omitempty"`
	WebDestinationSpec      *map[string]interface{}                   `json:"web_destination_spec,omitempty"`
}
