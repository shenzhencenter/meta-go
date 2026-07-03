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

var LeadGenContextCardFields = struct {
	ButtonText string
	Content    string
	CoverPhoto string
	ID         string
	Style      string
	Title      string
}{
	ButtonText: "button_text",
	Content:    "content",
	CoverPhoto: "cover_photo",
	ID:         "id",
	Style:      "style",
	Title:      "title",
}
