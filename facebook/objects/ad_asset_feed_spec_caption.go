package objects

type AdAssetFeedSpecCaption struct {
	Adlabels *[]AdAssetFeedSpecAssetLabel `json:"adlabels,omitempty"`
	Text     *string                      `json:"text,omitempty"`
	URLTags  *string                      `json:"url_tags,omitempty"`
}
