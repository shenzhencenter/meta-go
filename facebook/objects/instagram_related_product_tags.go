package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type InstagramRelatedProductTags struct {
	CheckoutSetting *string  `json:"checkout_setting,omitempty"`
	ID              *core.ID `json:"id,omitempty"`
	ImageURI        *string  `json:"image_uri,omitempty"`
	Name            *string  `json:"name,omitempty"`
	PriceLabel      *string  `json:"price_label,omitempty"`
	SalePriceLabel  *string  `json:"sale_price_label,omitempty"`
}
