package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdCreativePromotionMetadataSpec struct {
	EndDate         *int     `json:"end_date,omitempty"`
	ID              *core.ID `json:"id,omitempty"`
	PromotionSource *string  `json:"promotion_source,omitempty"`
	PromotionType   *string  `json:"promotion_type,omitempty"`
	PromotionValue  *float64 `json:"promotion_value,omitempty"`
	RequiredCode    *string  `json:"required_code,omitempty"`
	StartDate       *int     `json:"start_date,omitempty"`
}
