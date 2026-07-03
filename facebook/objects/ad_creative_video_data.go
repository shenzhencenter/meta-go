package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCreativeVideoData struct {
	AdditionalImageIndex                *int                                 `json:"additional_image_index,omitempty"`
	BrandedContentSharedToSponsorStatus *string                              `json:"branded_content_shared_to_sponsor_status,omitempty"`
	BrandedContentSponsorPageID         *core.ID                             `json:"branded_content_sponsor_page_id,omitempty"`
	CallToAction                        *AdCreativeLinkDataCallToAction      `json:"call_to_action,omitempty"`
	CollectionThumbnails                *[]AdCreativeCollectionThumbnailInfo `json:"collection_thumbnails,omitempty"`
	CustomizationRulesSpec              *[]AdCustomizationRuleSpec           `json:"customization_rules_spec,omitempty"`
	ImageHash                           *string                              `json:"image_hash,omitempty"`
	ImageURL                            *string                              `json:"image_url,omitempty"`
	LinkDescription                     *string                              `json:"link_description,omitempty"`
	Message                             *string                              `json:"message,omitempty"`
	OfferID                             *core.ID                             `json:"offer_id,omitempty"`
	PageWelcomeMessage                  *string                              `json:"page_welcome_message,omitempty"`
	PostClickConfiguration              *AdCreativePostClickConfiguration    `json:"post_click_configuration,omitempty"`
	RetailerItemIds                     *[]core.ID                           `json:"retailer_item_ids,omitempty"`
	Targeting                           *Targeting                           `json:"targeting,omitempty"`
	Title                               *string                              `json:"title,omitempty"`
	VideoID                             *core.ID                             `json:"video_id,omitempty"`
}
