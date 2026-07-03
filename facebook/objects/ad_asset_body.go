package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAssetBody struct {
	ID      *core.ID `json:"id,omitempty"`
	Text    *string  `json:"text,omitempty"`
	URLTags *string  `json:"url_tags,omitempty"`
}

var AdAssetBodyFields = struct {
	ID      string
	Text    string
	URLTags string
}{
	ID:      "id",
	Text:    "text",
	URLTags: "url_tags",
}
