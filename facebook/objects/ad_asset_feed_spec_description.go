package objects

type AdAssetFeedSpecDescription struct {
	Adlabels *[]AdAssetFeedSpecAssetLabel `json:"adlabels,omitempty"`
	Text     *string                      `json:"text,omitempty"`
	URLTags  *string                      `json:"url_tags,omitempty"`
}

var AdAssetFeedSpecDescriptionFields = struct {
	Adlabels string
	Text     string
	URLTags  string
}{
	Adlabels: "adlabels",
	Text:     "text",
	URLTags:  "url_tags",
}
