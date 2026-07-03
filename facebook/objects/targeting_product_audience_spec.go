package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type TargetingProductAudienceSpec struct {
	Exclusions   *[]TargetingProductAudienceSubSpec `json:"exclusions,omitempty"`
	Inclusions   *[]TargetingProductAudienceSubSpec `json:"inclusions,omitempty"`
	ProductSetID *core.ID                           `json:"product_set_id,omitempty"`
}
