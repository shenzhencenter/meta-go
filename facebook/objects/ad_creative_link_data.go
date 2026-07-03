package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
)

type AdCreativeLinkData struct {
	AdContext                           *string                               `json:"ad_context,omitempty"`
	AdditionalImageIndex                *int                                  `json:"additional_image_index,omitempty"`
	AppLinkSpec                         *AdCreativeLinkDataAppLinkSpec        `json:"app_link_spec,omitempty"`
	AttachmentStyle                     *string                               `json:"attachment_style,omitempty"`
	AutomatedProductTags                *bool                                 `json:"automated_product_tags,omitempty"`
	BoostedProductSetID                 *core.ID                              `json:"boosted_product_set_id,omitempty"`
	BrandedContentSharedToSponsorStatus *string                               `json:"branded_content_shared_to_sponsor_status,omitempty"`
	BrandedContentSponsorPageID         *core.ID                              `json:"branded_content_sponsor_page_id,omitempty"`
	CallToAction                        *AdCreativeLinkDataCallToAction       `json:"call_to_action,omitempty"`
	Caption                             *string                               `json:"caption,omitempty"`
	ChildAttachments                    *[]AdCreativeLinkDataChildAttachment  `json:"child_attachments,omitempty"`
	CollectionThumbnails                *[]AdCreativeCollectionThumbnailInfo  `json:"collection_thumbnails,omitempty"`
	CustomizationRulesSpec              *[]AdCustomizationRuleSpec            `json:"customization_rules_spec,omitempty"`
	Description                         *string                               `json:"description,omitempty"`
	EventID                             *core.ID                              `json:"event_id,omitempty"`
	ForceSingleLink                     *bool                                 `json:"force_single_link,omitempty"`
	FormatOption                        *enums.AdcreativelinkdataFormatOption `json:"format_option,omitempty"`
	ImageCrops                          *AdsImageCrops                        `json:"image_crops,omitempty"`
	ImageHash                           *string                               `json:"image_hash,omitempty"`
	ImageLayerSpecs                     *[]AdCreativeLinkDataImageLayerSpec   `json:"image_layer_specs,omitempty"`
	ImageOverlaySpec                    *AdCreativeLinkDataImageOverlaySpec   `json:"image_overlay_spec,omitempty"`
	IsLocalExpansion                    *bool                                 `json:"is_local_expansion,omitempty"`
	Link                                *string                               `json:"link,omitempty"`
	Message                             *string                               `json:"message,omitempty"`
	MultiShareEndCard                   *bool                                 `json:"multi_share_end_card,omitempty"`
	MultiShareOptimized                 *bool                                 `json:"multi_share_optimized,omitempty"`
	Name                                *string                               `json:"name,omitempty"`
	OfferID                             *core.ID                              `json:"offer_id,omitempty"`
	PageWelcomeMessage                  *string                               `json:"page_welcome_message,omitempty"`
	Picture                             *string                               `json:"picture,omitempty"`
	PostClickConfiguration              *AdCreativePostClickConfiguration     `json:"post_click_configuration,omitempty"`
	PreferredImageTags                  *[]string                             `json:"preferred_image_tags,omitempty"`
	PreferredVideoTags                  *[]string                             `json:"preferred_video_tags,omitempty"`
	RetailerItemIds                     *[]core.ID                            `json:"retailer_item_ids,omitempty"`
	ShowMultipleImages                  *bool                                 `json:"show_multiple_images,omitempty"`
	SmartPseEnabled                     *bool                                 `json:"smart_pse_enabled,omitempty"`
	StaticFallbackSpec                  *AdCreativeStaticFallbackSpec         `json:"static_fallback_spec,omitempty"`
	UseFlexibleImageAspectRatio         *bool                                 `json:"use_flexible_image_aspect_ratio,omitempty"`
}
