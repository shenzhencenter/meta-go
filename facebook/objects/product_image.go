package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type ProductImage struct {
	Height   *int     `json:"height,omitempty"`
	ID       *core.ID `json:"id,omitempty"`
	ImageURL *string  `json:"image_url,omitempty"`
	Width    *int     `json:"width,omitempty"`
}
