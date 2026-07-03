package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ShadowIGMediaProductTags struct {
	ImageURL                *string  `json:"image_url,omitempty"`
	IsCheckout              *bool    `json:"is_checkout,omitempty"`
	MerchantID              *core.ID `json:"merchant_id,omitempty"`
	Name                    *string  `json:"name,omitempty"`
	PriceString             *string  `json:"price_string,omitempty"`
	ProductID               *core.ID `json:"product_id,omitempty"`
	ReviewStatus            *string  `json:"review_status,omitempty"`
	StrippedPriceString     *string  `json:"stripped_price_string,omitempty"`
	StrippedSalePriceString *string  `json:"stripped_sale_price_string,omitempty"`
	X                       *float64 `json:"x,omitempty"`
	Y                       *float64 `json:"y,omitempty"`
}
