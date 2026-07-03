package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
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

var AdAssetFeedSpecFields = struct {
	AdFormats               string
	AdditionalData          string
	AppProductPageID        string
	AssetCustomizationRules string
	Audios                  string
	Autotranslate           string
	Bodies                  string
	CallAdsConfiguration    string
	CallToActionTypes       string
	CallToActions           string
	Captions                string
	Carousels               string
	CtwaConsentData         string
	Descriptions            string
	Events                  string
	Groups                  string
	Images                  string
	LinkUrls                string
	MessageExtensions       string
	OnsiteDestinations      string
	OptimizationType        string
	PromotionalMetadata     string
	ReasonsToShop           string
	ShopsBundle             string
	Titles                  string
	Translations            string
	UpcomingEvents          string
	Videos                  string
	WebDestinationSpec      string
}{
	AdFormats:               "ad_formats",
	AdditionalData:          "additional_data",
	AppProductPageID:        "app_product_page_id",
	AssetCustomizationRules: "asset_customization_rules",
	Audios:                  "audios",
	Autotranslate:           "autotranslate",
	Bodies:                  "bodies",
	CallAdsConfiguration:    "call_ads_configuration",
	CallToActionTypes:       "call_to_action_types",
	CallToActions:           "call_to_actions",
	Captions:                "captions",
	Carousels:               "carousels",
	CtwaConsentData:         "ctwa_consent_data",
	Descriptions:            "descriptions",
	Events:                  "events",
	Groups:                  "groups",
	Images:                  "images",
	LinkUrls:                "link_urls",
	MessageExtensions:       "message_extensions",
	OnsiteDestinations:      "onsite_destinations",
	OptimizationType:        "optimization_type",
	PromotionalMetadata:     "promotional_metadata",
	ReasonsToShop:           "reasons_to_shop",
	ShopsBundle:             "shops_bundle",
	Titles:                  "titles",
	Translations:            "translations",
	UpcomingEvents:          "upcoming_events",
	Videos:                  "videos",
	WebDestinationSpec:      "web_destination_spec",
}
