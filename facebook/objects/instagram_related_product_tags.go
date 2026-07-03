package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type InstagramRelatedProductTags struct {
	CheckoutSetting *string  `json:"checkout_setting,omitempty"`
	ID              *core.ID `json:"id,omitempty"`
	ImageURI        *string  `json:"image_uri,omitempty"`
	Name            *string  `json:"name,omitempty"`
	PriceLabel      *string  `json:"price_label,omitempty"`
	SalePriceLabel  *string  `json:"sale_price_label,omitempty"`
}

var InstagramRelatedProductTagsFields = struct {
	CheckoutSetting string
	ID              string
	ImageURI        string
	Name            string
	PriceLabel      string
	SalePriceLabel  string
}{
	CheckoutSetting: "checkout_setting",
	ID:              "id",
	ImageURI:        "image_uri",
	Name:            "name",
	PriceLabel:      "price_label",
	SalePriceLabel:  "sale_price_label",
}
