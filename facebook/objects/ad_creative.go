package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AdCreative struct {
	AccountID                        *core.ID                                  `json:"account_id,omitempty"`
	ActorID                          *core.ID                                  `json:"actor_id,omitempty"`
	AdDisclaimerSpec                 *AdCreativeAdDisclaimer                   `json:"ad_disclaimer_spec,omitempty"`
	Adlabels                         *[]AdLabel                                `json:"adlabels,omitempty"`
	ApplinkTreatment                 *string                                   `json:"applink_treatment,omitempty"`
	AssetFeedSpec                    *AdAssetFeedSpec                          `json:"asset_feed_spec,omitempty"`
	AuthorizationCategory            *string                                   `json:"authorization_category,omitempty"`
	AutoUpdate                       *bool                                     `json:"auto_update,omitempty"`
	Body                             *string                                   `json:"body,omitempty"`
	BrandedContent                   *AdCreativeBrandedContentAds              `json:"branded_content,omitempty"`
	BrandedContentSponsorPageID      *core.ID                                  `json:"branded_content_sponsor_page_id,omitempty"`
	BundleFolderID                   *core.ID                                  `json:"bundle_folder_id,omitempty"`
	CallToAction                     *AdCreativeLinkDataCallToAction           `json:"call_to_action,omitempty"`
	CallToActionType                 *enums.AdcreativeCallToActionType         `json:"call_to_action_type,omitempty"`
	CategorizationCriteria           *string                                   `json:"categorization_criteria,omitempty"`
	CategoryMediaSource              *string                                   `json:"category_media_source,omitempty"`
	CollaborativeAdsLsbImageBankID   *core.ID                                  `json:"collaborative_ads_lsb_image_bank_id,omitempty"`
	ContextualMultiAds               *AdCreativeContextualMultiAds             `json:"contextual_multi_ads,omitempty"`
	CreativeSourcingSpec             *AdCreativeSourcingSpec                   `json:"creative_sourcing_spec,omitempty"`
	DegreesOfFreedomSpec             *AdCreativeDegreesOfFreedomSpec           `json:"degrees_of_freedom_spec,omitempty"`
	DestinationSetID                 *core.ID                                  `json:"destination_set_id,omitempty"`
	DestinationSpec                  *AdCreativeDestinationSpec                `json:"destination_spec,omitempty"`
	DynamicAdVoice                   *string                                   `json:"dynamic_ad_voice,omitempty"`
	EffectiveAuthorizationCategory   *string                                   `json:"effective_authorization_category,omitempty"`
	EffectiveInstagramMediaID        *core.ID                                  `json:"effective_instagram_media_id,omitempty"`
	EffectiveObjectStoryID           *core.ID                                  `json:"effective_object_story_id,omitempty"`
	EnableDirectInstall              *bool                                     `json:"enable_direct_install,omitempty"`
	EnableLaunchInstantApp           *bool                                     `json:"enable_launch_instant_app,omitempty"`
	ExistingPostTitle                *string                                   `json:"existing_post_title,omitempty"`
	FacebookBrandedContent           *AdCreativeFacebookBrandedContent         `json:"facebook_branded_content,omitempty"`
	FormatTransformationSpec         *[]AdCreativeFormatTransformationSpec     `json:"format_transformation_spec,omitempty"`
	GenerativeAssetSpec              *AdCreativeGenerativeAssetSpec            `json:"generative_asset_spec,omitempty"`
	ID                               *core.ID                                  `json:"id,omitempty"`
	ImageCrops                       *AdsImageCrops                            `json:"image_crops,omitempty"`
	ImageHash                        *string                                   `json:"image_hash,omitempty"`
	ImageURL                         *string                                   `json:"image_url,omitempty"`
	InstagramBrandedContent          *AdCreativeInstagramBrandedContent        `json:"instagram_branded_content,omitempty"`
	InstagramPermalinkURL            *string                                   `json:"instagram_permalink_url,omitempty"`
	InstagramUserID                  *core.ID                                  `json:"instagram_user_id,omitempty"`
	InteractiveComponentsSpec        *AdCreativeInteractiveComponentsSpec      `json:"interactive_components_spec,omitempty"`
	LinkDeepLinkURL                  *string                                   `json:"link_deep_link_url,omitempty"`
	LinkDestinationDisplayURL        *string                                   `json:"link_destination_display_url,omitempty"`
	LinkOgID                         *core.ID                                  `json:"link_og_id,omitempty"`
	LinkURL                          *string                                   `json:"link_url,omitempty"`
	MarketingMessageStructuredSpec   *AdCreativeMarketingMessageStructuredSpec `json:"marketing_message_structured_spec,omitempty"`
	MediaSourcingSpec                *AdCreativeMediaSourcingSpec              `json:"media_sourcing_spec,omitempty"`
	MessengerSponsoredMessage        *string                                   `json:"messenger_sponsored_message,omitempty"`
	Name                             *string                                   `json:"name,omitempty"`
	ObjectID                         *core.ID                                  `json:"object_id,omitempty"`
	ObjectStoreURL                   *string                                   `json:"object_store_url,omitempty"`
	ObjectStoryID                    *core.ID                                  `json:"object_story_id,omitempty"`
	ObjectStorySpec                  *AdCreativeObjectStorySpec                `json:"object_story_spec,omitempty"`
	ObjectType                       *enums.AdcreativeObjectType               `json:"object_type,omitempty"`
	ObjectURL                        *string                                   `json:"object_url,omitempty"`
	OmnichannelLinkSpec              *AdCreativeOmnichannelLinkSpec            `json:"omnichannel_link_spec,omitempty"`
	PageWelcomeMessage               *string                                   `json:"page_welcome_message,omitempty"`
	PhotoAlbumSourceObjectStoryID    *core.ID                                  `json:"photo_album_source_object_story_id,omitempty"`
	PlacePageSetID                   *core.ID                                  `json:"place_page_set_id,omitempty"`
	PlatformCustomizations           *AdCreativePlatformCustomization          `json:"platform_customizations,omitempty"`
	PlayableAssetID                  *core.ID                                  `json:"playable_asset_id,omitempty"`
	PortraitCustomizations           *AdCreativePortraitCustomizations         `json:"portrait_customizations,omitempty"`
	ProductData                      *[]AdCreativeProductData                  `json:"product_data,omitempty"`
	ProductSetID                     *core.ID                                  `json:"product_set_id,omitempty"`
	ProductSuggestionSettings        *AdCreativeProductSuggestionSettings      `json:"product_suggestion_settings,omitempty"`
	RecommenderSettings              *AdCreativeRecommenderSettings            `json:"recommender_settings,omitempty"`
	RegionalRegulationDisclaimerSpec *AdCreativeRegionalRegulationDisclaimer   `json:"regional_regulation_disclaimer_spec,omitempty"`
	SourceFacebookPostID             *core.ID                                  `json:"source_facebook_post_id,omitempty"`
	SourceInstagramMediaID           *core.ID                                  `json:"source_instagram_media_id,omitempty"`
	Status                           *enums.AdcreativeStatusEnum               `json:"status,omitempty"`
	TemplateURL                      *string                                   `json:"template_url,omitempty"`
	TemplateURLSpec                  *AdCreativeTemplateURLSpec                `json:"template_url_spec,omitempty"`
	ThumbnailID                      *core.ID                                  `json:"thumbnail_id,omitempty"`
	ThumbnailURL                     *string                                   `json:"thumbnail_url,omitempty"`
	Title                            *string                                   `json:"title,omitempty"`
	URLTags                          *string                                   `json:"url_tags,omitempty"`
	UsePageActorOverride             *bool                                     `json:"use_page_actor_override,omitempty"`
	VideoID                          *core.ID                                  `json:"video_id,omitempty"`
	WamoWhatsappIdentitySpec         *AdCreativeWAMOWhatsAppIdentitySpec       `json:"wamo_whatsapp_identity_spec,omitempty"`
}
