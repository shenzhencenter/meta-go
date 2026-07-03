package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdAssetBody struct {
	ID      *core.ID `json:"id,omitempty"`
	Text    *string  `json:"text,omitempty"`
	URLTags *string  `json:"url_tags,omitempty"`
}
