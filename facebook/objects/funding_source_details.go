package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type FundingSourceDetails struct {
	Coupon        *FundingSourceDetailsCoupon   `json:"coupon,omitempty"`
	Coupons       *[]FundingSourceDetailsCoupon `json:"coupons,omitempty"`
	DisplayString *string                       `json:"display_string,omitempty"`
	ID            *core.ID                      `json:"id,omitempty"`
	Type          *int                          `json:"type,omitempty"`
}

var FundingSourceDetailsFields = struct {
	Coupon        string
	Coupons       string
	DisplayString string
	ID            string
	Type          string
}{
	Coupon:        "coupon",
	Coupons:       "coupons",
	DisplayString: "display_string",
	ID:            "id",
	Type:          "type",
}
