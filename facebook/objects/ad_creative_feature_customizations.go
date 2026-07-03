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
