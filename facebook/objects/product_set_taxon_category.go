package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ProductSetTaxonCategory struct {
	CategoryID   *core.ID `json:"category_id,omitempty"`
	CategoryName *string  `json:"category_name,omitempty"`
	ImageURL     *string  `json:"image_url,omitempty"`
}

var ProductSetTaxonCategoryFields = struct {
	CategoryID   string
	CategoryName string
	ImageURL     string
}{
	CategoryID:   "category_id",
	CategoryName: "category_name",
	ImageURL:     "image_url",
}
