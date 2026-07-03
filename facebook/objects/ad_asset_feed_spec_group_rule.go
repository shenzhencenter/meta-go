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

var AdAssetFeedSpecGroupRuleFields = struct {
	BodyLabel        string
	CaptionLabel     string
	DescriptionLabel string
	ImageLabel       string
	LinkURLLabel     string
	TitleLabel       string
	VideoLabel       string
}{
	BodyLabel:        "body_label",
	CaptionLabel:     "caption_label",
	DescriptionLabel: "description_label",
	ImageLabel:       "image_label",
	LinkURLLabel:     "link_url_label",
	TitleLabel:       "title_label",
	VideoLabel:       "video_label",
}
