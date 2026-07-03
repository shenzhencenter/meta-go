package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var AdCreativePromotionMetadataSpecFields = struct {
	EndDate         string
	ID              string
	PromotionSource string
	PromotionType   string
	PromotionValue  string
	RequiredCode    string
	StartDate       string
}{
	EndDate:         "end_date",
	ID:              "id",
	PromotionSource: "promotion_source",
	PromotionType:   "promotion_type",
	PromotionValue:  "promotion_value",
	RequiredCode:    "required_code",
	StartDate:       "start_date",
}
