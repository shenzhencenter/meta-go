package objects

type AdAssetFeedSpecGroupRule struct {
	BodyLabel        *AdAssetFeedSpecAssetLabel `json:"body_label,omitempty"`
	CaptionLabel     *AdAssetFeedSpecAssetLabel `json:"caption_label,omitempty"`
	DescriptionLabel *AdAssetFeedSpecAssetLabel `json:"description_label,omitempty"`
	ImageLabel       *AdAssetFeedSpecAssetLabel `json:"image_label,omitempty"`
	LinkURLLabel     *AdAssetFeedSpecAssetLabel `json:"link_url_label,omitempty"`
	TitleLabel       *AdAssetFeedSpecAssetLabel `json:"title_label,omitempty"`
	VideoLabel       *AdAssetFeedSpecAssetLabel `json:"video_label,omitempty"`
}
