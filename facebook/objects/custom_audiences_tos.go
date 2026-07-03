package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CustomAudiencesTOS struct {
	Content *string  `json:"content,omitempty"`
	ID      *core.ID `json:"id,omitempty"`
	Type    *string  `json:"type,omitempty"`
}

var CustomAudiencesTOSFields = struct {
	Content string
	ID      string
	Type    string
}{
	Content: "content",
	ID:      "id",
	Type:    "type",
}
