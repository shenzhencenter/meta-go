package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCreativeBrandedContentAds struct {
	ActingBusinessID                 *core.ID                               `json:"acting_business_id,omitempty"`
	AdFormat                         *int                                   `json:"ad_format,omitempty"`
	AutomatedAdgroupCreation         *bool                                  `json:"automated_adgroup_creation,omitempty"`
	ContentSearchInput               *string                                `json:"content_search_input,omitempty"`
	CreatorAdPermissionType          *string                                `json:"creator_ad_permission_type,omitempty"`
	DeliverDynamicPartnerContent     *bool                                  `json:"deliver_dynamic_partner_content,omitempty"`
	FacebookBoostPostAccessToken     *string                                `json:"facebook_boost_post_access_token,omitempty"`
	InstagramBoostPostAccessToken    *string                                `json:"instagram_boost_post_access_token,omitempty"`
	IsMcaInternal                    *bool                                  `json:"is_mca_internal,omitempty"`
	ParentSourceFacebookPostID       *core.ID                               `json:"parent_source_facebook_post_id,omitempty"`
	ParentSourceInstagramMediaID     *core.ID                               `json:"parent_source_instagram_media_id,omitempty"`
	Partners                         *[]AdCreativeBrandedContentAdsPartners `json:"partners,omitempty"`
	ProductSetPartnerSelectionStatus *string                                `json:"product_set_partner_selection_status,omitempty"`
	PromotedPageID                   *core.ID                               `json:"promoted_page_id,omitempty"`
	Testimonial                      *string                                `json:"testimonial,omitempty"`
	TestimonialLocale                *string                                `json:"testimonial_locale,omitempty"`
	UiVersion                        *int                                   `json:"ui_version,omitempty"`
}

var AdCreativeBrandedContentAdsFields = struct {
	ActingBusinessID                 string
	AdFormat                         string
	AutomatedAdgroupCreation         string
	ContentSearchInput               string
	CreatorAdPermissionType          string
	DeliverDynamicPartnerContent     string
	FacebookBoostPostAccessToken     string
	InstagramBoostPostAccessToken    string
	IsMcaInternal                    string
	ParentSourceFacebookPostID       string
	ParentSourceInstagramMediaID     string
	Partners                         string
	ProductSetPartnerSelectionStatus string
	PromotedPageID                   string
	Testimonial                      string
	TestimonialLocale                string
	UiVersion                        string
}{
	ActingBusinessID:                 "acting_business_id",
	AdFormat:                         "ad_format",
	AutomatedAdgroupCreation:         "automated_adgroup_creation",
	ContentSearchInput:               "content_search_input",
	CreatorAdPermissionType:          "creator_ad_permission_type",
	DeliverDynamicPartnerContent:     "deliver_dynamic_partner_content",
	FacebookBoostPostAccessToken:     "facebook_boost_post_access_token",
	InstagramBoostPostAccessToken:    "instagram_boost_post_access_token",
	IsMcaInternal:                    "is_mca_internal",
	ParentSourceFacebookPostID:       "parent_source_facebook_post_id",
	ParentSourceInstagramMediaID:     "parent_source_instagram_media_id",
	Partners:                         "partners",
	ProductSetPartnerSelectionStatus: "product_set_partner_selection_status",
	PromotedPageID:                   "promoted_page_id",
	Testimonial:                      "testimonial",
	TestimonialLocale:                "testimonial_locale",
	UiVersion:                        "ui_version",
}
