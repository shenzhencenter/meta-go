package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
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
