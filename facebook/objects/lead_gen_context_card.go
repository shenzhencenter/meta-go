package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type LeadGenContextCard struct {
	ButtonText *string   `json:"button_text,omitempty"`
	Content    *[]string `json:"content,omitempty"`
	CoverPhoto *Photo    `json:"cover_photo,omitempty"`
	ID         *core.ID  `json:"id,omitempty"`
	Style      *string   `json:"style,omitempty"`
	Title      *string   `json:"title,omitempty"`
}
