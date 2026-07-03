package objects

type AdAssetFeedSpecImage struct {
	Adlabels   *[]AdAssetFeedSpecAssetLabel `json:"adlabels,omitempty"`
	Hash       *string                      `json:"hash,omitempty"`
	ImageCrops *AdsImageCrops               `json:"image_crops,omitempty"`
	URL        *string                      `json:"url,omitempty"`
	URLTags    *string                      `json:"url_tags,omitempty"`
}

var AdAssetFeedSpecImageFields = struct {
	Adlabels   string
	Hash       string
	ImageCrops string
	URL        string
	URLTags    string
}{
	Adlabels:   "adlabels",
	Hash:       "hash",
	ImageCrops: "image_crops",
	URL:        "url",
	URLTags:    "url_tags",
}
