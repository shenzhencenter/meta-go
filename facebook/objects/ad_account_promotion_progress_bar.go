package objects

import (
	"time"
)

type AdAccountPromotionProgressBar struct {
	AdaccountPermission    *bool      `json:"adaccount_permission,omitempty"`
	CouponCurrency         *string    `json:"coupon_currency,omitempty"`
	CouponValue            *int       `json:"coupon_value,omitempty"`
	ExpirationTime         *time.Time `json:"expiration_time,omitempty"`
	ProgressCompleted      *bool      `json:"progress_completed,omitempty"`
	PromotionType          *string    `json:"promotion_type,omitempty"`
	SpendRequirementInCent *int       `json:"spend_requirement_in_cent,omitempty"`
	SpendSinceEnrollment   *int       `json:"spend_since_enrollment,omitempty"`
}
