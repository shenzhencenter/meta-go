package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type ProductSetTaxonCategory struct {
	CategoryID   *core.ID `json:"category_id,omitempty"`
	CategoryName *string  `json:"category_name,omitempty"`
	ImageURL     *string  `json:"image_url,omitempty"`
}
