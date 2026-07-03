package objects

type AdCreativeFeatureCustomizations struct {
	AspectRatioConfig   *map[string]interface{} `json:"aspect_ratio_config,omitempty"`
	BackgroundColor     *string                 `json:"background_color,omitempty"`
	CatalogFeedTagName  *string                 `json:"catalog_feed_tag_name,omitempty"`
	FbFeedTagName       *string                 `json:"fb_feed_tag_name,omitempty"`
	FbReelsTagName      *string                 `json:"fb_reels_tag_name,omitempty"`
	FbStoryTagName      *string                 `json:"fb_story_tag_name,omitempty"`
	FontName            *string                 `json:"font_name,omitempty"`
	IgFeedTagName       *string                 `json:"ig_feed_tag_name,omitempty"`
	IgReelsTagName      *string                 `json:"ig_reels_tag_name,omitempty"`
	IgStreamTagName     *string                 `json:"ig_stream_tag_name,omitempty"`
	ImageCropStyle      *string                 `json:"image_crop_style,omitempty"`
	PeCarousel          *map[string]interface{} `json:"pe_carousel,omitempty"`
	RecompositionType   *map[string]interface{} `json:"recomposition_type,omitempty"`
	ShowcaseCardDisplay *string                 `json:"showcase_card_display,omitempty"`
	TextExtraction      *map[string]interface{} `json:"text_extraction,omitempty"`
	TextStyle           *string                 `json:"text_style,omitempty"`
}

var AdCreativeFeatureCustomizationsFields = struct {
	AspectRatioConfig   string
	BackgroundColor     string
	CatalogFeedTagName  string
	FbFeedTagName       string
	FbReelsTagName      string
	FbStoryTagName      string
	FontName            string
	IgFeedTagName       string
	IgReelsTagName      string
	IgStreamTagName     string
	ImageCropStyle      string
	PeCarousel          string
	RecompositionType   string
	ShowcaseCardDisplay string
	TextExtraction      string
	TextStyle           string
}{
	AspectRatioConfig:   "aspect_ratio_config",
	BackgroundColor:     "background_color",
	CatalogFeedTagName:  "catalog_feed_tag_name",
	FbFeedTagName:       "fb_feed_tag_name",
	FbReelsTagName:      "fb_reels_tag_name",
	FbStoryTagName:      "fb_story_tag_name",
	FontName:            "font_name",
	IgFeedTagName:       "ig_feed_tag_name",
	IgReelsTagName:      "ig_reels_tag_name",
	IgStreamTagName:     "ig_stream_tag_name",
	ImageCropStyle:      "image_crop_style",
	PeCarousel:          "pe_carousel",
	RecompositionType:   "recomposition_type",
	ShowcaseCardDisplay: "showcase_card_display",
	TextExtraction:      "text_extraction",
	TextStyle:           "text_style",
}
