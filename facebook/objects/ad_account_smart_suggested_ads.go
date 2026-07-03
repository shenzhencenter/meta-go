package objects

type AdAccountSmartSuggestedAds struct {
	AdCreativeSpec *string   `json:"ad_creative_spec,omitempty"`
	Description    *string   `json:"description,omitempty"`
	GuidanceSpec   *[]string `json:"guidance_spec,omitempty"`
	ThumbnailURL   *string   `json:"thumbnail_url,omitempty"`
}

var AdAccountSmartSuggestedAdsFields = struct {
	AdCreativeSpec string
	Description    string
	GuidanceSpec   string
	ThumbnailURL   string
}{
	AdCreativeSpec: "ad_creative_spec",
	Description:    "description",
	GuidanceSpec:   "guidance_spec",
	ThumbnailURL:   "thumbnail_url",
}
