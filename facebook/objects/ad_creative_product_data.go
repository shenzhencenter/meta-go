package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdCreativeProductData struct {
	ProductID     *core.ID `json:"product_id,omitempty"`
	ProductSource *string  `json:"product_source,omitempty"`
}
