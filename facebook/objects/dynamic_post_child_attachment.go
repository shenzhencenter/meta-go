package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type DynamicPostChildAttachment struct {
	Description *string  `json:"description,omitempty"`
	ImageURL    *string  `json:"image_url,omitempty"`
	Link        *string  `json:"link,omitempty"`
	PlaceID     *core.ID `json:"place_id,omitempty"`
	ProductID   *core.ID `json:"product_id,omitempty"`
	Title       *string  `json:"title,omitempty"`
}

var DynamicPostChildAttachmentFields = struct {
	Description string
	ImageURL    string
	Link        string
	PlaceID     string
	ProductID   string
	Title       string
}{
	Description: "description",
	ImageURL:    "image_url",
	Link:        "link",
	PlaceID:     "place_id",
	ProductID:   "product_id",
	Title:       "title",
}
