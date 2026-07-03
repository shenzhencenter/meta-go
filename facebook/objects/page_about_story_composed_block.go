package objects

type PageAboutStoryComposedBlock struct {
	Depth             *int                                       `json:"depth,omitempty"`
	EntityRanges      *[]PageAboutStoryComposedBlockEntityRanges `json:"entity_ranges,omitempty"`
	InlineStyleRanges *[]PageAboutStoryComposedBlockInlineStyle  `json:"inline_style_ranges,omitempty"`
	Text              *string                                    `json:"text,omitempty"`
	Type              *string                                    `json:"type,omitempty"`
}
