package objects

type AdAssetFeedSpecCarouselChildAttachment struct {
	BodyLabel             *AdAssetFeedSpecAssetLabel `json:"body_label,omitempty"`
	CallToActionTypeLabel *AdAssetFeedSpecAssetLabel `json:"call_to_action_type_label,omitempty"`
	CaptionLabel          *AdAssetFeedSpecAssetLabel `json:"caption_label,omitempty"`
	DescriptionLabel      *AdAssetFeedSpecAssetLabel `json:"description_label,omitempty"`
	ImageLabel            *AdAssetFeedSpecAssetLabel `json:"image_label,omitempty"`
	LinkURLLabel          *AdAssetFeedSpecAssetLabel `json:"link_url_label,omitempty"`
	PhoneDataIdsLabel     *AdAssetFeedSpecAssetLabel `json:"phone_data_ids_label,omitempty"`
	StaticCard            *bool                      `json:"static_card,omitempty"`
	TitleLabel            *AdAssetFeedSpecAssetLabel `json:"title_label,omitempty"`
	VideoLabel            *AdAssetFeedSpecAssetLabel `json:"video_label,omitempty"`
}
