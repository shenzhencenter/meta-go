package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCreativeProductData struct {
	ProductID     *core.ID `json:"product_id,omitempty"`
	ProductSource *string  `json:"product_source,omitempty"`
}

var AdCreativeProductDataFields = struct {
	ProductID     string
	ProductSource string
}{
	ProductID:     "product_id",
	ProductSource: "product_source",
}
