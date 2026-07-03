package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ShadowIGUserCatalogProductVariant struct {
	ProductID   *core.ID `json:"product_id,omitempty"`
	VariantName *string  `json:"variant_name,omitempty"`
}

var ShadowIGUserCatalogProductVariantFields = struct {
	ProductID   string
	VariantName string
}{
	ProductID:   "product_id",
	VariantName: "variant_name",
}
