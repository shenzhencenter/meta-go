package objects

type AdAssetFeedSpecImage struct {
	Adlabels   *[]AdAssetFeedSpecAssetLabel `json:"adlabels,omitempty"`
	Hash       *string                      `json:"hash,omitempty"`
	ImageCrops *AdsImageCrops               `json:"image_crops,omitempty"`
	URL        *string                      `json:"url,omitempty"`
	URLTags    *string                      `json:"url_tags,omitempty"`
}
