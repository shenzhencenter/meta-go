package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type FundingSourceDetails struct {
	Coupon        *FundingSourceDetailsCoupon   `json:"coupon,omitempty"`
	Coupons       *[]FundingSourceDetailsCoupon `json:"coupons,omitempty"`
	DisplayString *string                       `json:"display_string,omitempty"`
	ID            *core.ID                      `json:"id,omitempty"`
	Type          *int                          `json:"type,omitempty"`
}
