package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type TargetingProductAudienceSpec struct {
	Exclusions   *[]TargetingProductAudienceSubSpec `json:"exclusions,omitempty"`
	Inclusions   *[]TargetingProductAudienceSubSpec `json:"inclusions,omitempty"`
	ProductSetID *core.ID                           `json:"product_set_id,omitempty"`
}

var TargetingProductAudienceSpecFields = struct {
	Exclusions   string
	Inclusions   string
	ProductSetID string
}{
	Exclusions:   "exclusions",
	Inclusions:   "inclusions",
	ProductSetID: "product_set_id",
}
