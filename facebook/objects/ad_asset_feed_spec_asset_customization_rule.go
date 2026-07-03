package objects

type AdAssetFeedSpecAssetCustomizationRule struct {
	BodyLabel             *AdAssetFeedSpecAssetLabel                 `json:"body_label,omitempty"`
	CallToActionLabel     *AdAssetFeedSpecAssetLabel                 `json:"call_to_action_label,omitempty"`
	CallToActionTypeLabel *AdAssetFeedSpecAssetLabel                 `json:"call_to_action_type_label,omitempty"`
	CaptionLabel          *AdAssetFeedSpecAssetLabel                 `json:"caption_label,omitempty"`
	CarouselLabel         *AdAssetFeedSpecAssetLabel                 `json:"carousel_label,omitempty"`
	CustomizationSpec     *AdAssetCustomizationRuleCustomizationSpec `json:"customization_spec,omitempty"`
	DescriptionLabel      *AdAssetFeedSpecAssetLabel                 `json:"description_label,omitempty"`
	ImageLabel            *AdAssetFeedSpecAssetLabel                 `json:"image_label,omitempty"`
	IsDefault             *bool                                      `json:"is_default,omitempty"`
	LinkURLLabel          *AdAssetFeedSpecAssetLabel                 `json:"link_url_label,omitempty"`
	Priority              *int                                       `json:"priority,omitempty"`
	TitleLabel            *AdAssetFeedSpecAssetLabel                 `json:"title_label,omitempty"`
	VideoLabel            *AdAssetFeedSpecAssetLabel                 `json:"video_label,omitempty"`
}
