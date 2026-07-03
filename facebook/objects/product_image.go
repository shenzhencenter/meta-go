package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ProductImage struct {
	Height   *int     `json:"height,omitempty"`
	ID       *core.ID `json:"id,omitempty"`
	ImageURL *string  `json:"image_url,omitempty"`
	Width    *int     `json:"width,omitempty"`
}

var ProductImageFields = struct {
	Height   string
	ID       string
	ImageURL string
	Width    string
}{
	Height:   "height",
	ID:       "id",
	ImageURL: "image_url",
	Width:    "width",
}
