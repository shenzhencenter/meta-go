package objects

type PageAboutStoryComposedBlockInlineStyle struct {
	Length *int    `json:"length,omitempty"`
	Offset *int    `json:"offset,omitempty"`
	Style  *string `json:"style,omitempty"`
}

var PageAboutStoryComposedBlockInlineStyleFields = struct {
	Length string
	Offset string
	Style  string
}{
	Length: "length",
	Offset: "offset",
	Style:  "style",
}
