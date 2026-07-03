package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCreativeSourcingSpec struct {
	AdExtensionsRelevancySpec        *map[string]interface{}            `json:"ad_extensions_relevancy_spec,omitempty"`
	AssociatedProductSetID           *core.ID                           `json:"associated_product_set_id,omitempty"`
	Brand                            *map[string]interface{}            `json:"brand,omitempty"`
	DestinationScreenshotSpec        *map[string]interface{}            `json:"destination_screenshot_spec,omitempty"`
	DuplicationSource                *string                            `json:"duplication_source,omitempty"`
	DynamicSiteLinksSpec             *map[string]interface{}            `json:"dynamic_site_links_spec,omitempty"`
	EnableSocialFeedbackPreservation *bool                              `json:"enable_social_feedback_preservation,omitempty"`
	FeaturedOfferingSpec             *map[string]interface{}            `json:"featured_offering_spec,omitempty"`
	Intent                           *map[string]interface{}            `json:"intent,omitempty"`
	PcaSpec                          *map[string]interface{}            `json:"pca_spec,omitempty"`
	ProductMediaMetadataSpec         *map[string]interface{}            `json:"product_media_metadata_spec,omitempty"`
	PromotionMetadataSpec            *[]AdCreativePromotionMetadataSpec `json:"promotion_metadata_spec,omitempty"`
	SiteLinksDataConsented           *map[string]interface{}            `json:"site_links_data_consented,omitempty"`
	SiteLinksSpec                    *[]AdCreativeSiteLinksSpec         `json:"site_links_spec,omitempty"`
	SourceURL                        *string                            `json:"source_url,omitempty"`
	WebsiteMediaSpec                 *map[string]interface{}            `json:"website_media_spec,omitempty"`
	WebsiteSummarySpec               *map[string]interface{}            `json:"website_summary_spec,omitempty"`
}

var AdCreativeSourcingSpecFields = struct {
	AdExtensionsRelevancySpec        string
	AssociatedProductSetID           string
	Brand                            string
	DestinationScreenshotSpec        string
	DuplicationSource                string
	DynamicSiteLinksSpec             string
	EnableSocialFeedbackPreservation string
	FeaturedOfferingSpec             string
	Intent                           string
	PcaSpec                          string
	ProductMediaMetadataSpec         string
	PromotionMetadataSpec            string
	SiteLinksDataConsented           string
	SiteLinksSpec                    string
	SourceURL                        string
	WebsiteMediaSpec                 string
	WebsiteSummarySpec               string
}{
	AdExtensionsRelevancySpec:        "ad_extensions_relevancy_spec",
	AssociatedProductSetID:           "associated_product_set_id",
	Brand:                            "brand",
	DestinationScreenshotSpec:        "destination_screenshot_spec",
	DuplicationSource:                "duplication_source",
	DynamicSiteLinksSpec:             "dynamic_site_links_spec",
	EnableSocialFeedbackPreservation: "enable_social_feedback_preservation",
	FeaturedOfferingSpec:             "featured_offering_spec",
	Intent:                           "intent",
	PcaSpec:                          "pca_spec",
	ProductMediaMetadataSpec:         "product_media_metadata_spec",
	PromotionMetadataSpec:            "promotion_metadata_spec",
	SiteLinksDataConsented:           "site_links_data_consented",
	SiteLinksSpec:                    "site_links_spec",
	SourceURL:                        "source_url",
	WebsiteMediaSpec:                 "website_media_spec",
	WebsiteSummarySpec:               "website_summary_spec",
}
