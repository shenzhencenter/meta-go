package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type DynamicPostChildAttachment struct {
	Description *string  `json:"description,omitempty"`
	ImageURL    *string  `json:"image_url,omitempty"`
	Link        *string  `json:"link,omitempty"`
	PlaceID     *core.ID `json:"place_id,omitempty"`
	ProductID   *core.ID `json:"product_id,omitempty"`
	Title       *string  `json:"title,omitempty"`
}
