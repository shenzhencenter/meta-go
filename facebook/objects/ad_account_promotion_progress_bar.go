package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAccountPromotionProgressBar struct {
	AdaccountPermission    *bool      `json:"adaccount_permission,omitempty"`
	CouponCurrency         *string    `json:"coupon_currency,omitempty"`
	CouponValue            *int       `json:"coupon_value,omitempty"`
	ExpirationTime         *core.Time `json:"expiration_time,omitempty"`
	ProgressCompleted      *bool      `json:"progress_completed,omitempty"`
	PromotionType          *string    `json:"promotion_type,omitempty"`
	SpendRequirementInCent *int       `json:"spend_requirement_in_cent,omitempty"`
	SpendSinceEnrollment   *int       `json:"spend_since_enrollment,omitempty"`
}

var AdAccountPromotionProgressBarFields = struct {
	AdaccountPermission    string
	CouponCurrency         string
	CouponValue            string
	ExpirationTime         string
	ProgressCompleted      string
	PromotionType          string
	SpendRequirementInCent string
	SpendSinceEnrollment   string
}{
	AdaccountPermission:    "adaccount_permission",
	CouponCurrency:         "coupon_currency",
	CouponValue:            "coupon_value",
	ExpirationTime:         "expiration_time",
	ProgressCompleted:      "progress_completed",
	PromotionType:          "promotion_type",
	SpendRequirementInCent: "spend_requirement_in_cent",
	SpendSinceEnrollment:   "spend_since_enrollment",
}
