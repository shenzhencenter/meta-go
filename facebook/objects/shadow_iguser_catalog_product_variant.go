package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type ShadowIGUserCatalogProductVariant struct {
	ProductID   *core.ID `json:"product_id,omitempty"`
	VariantName *string  `json:"variant_name,omitempty"`
}
