package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type IGShoppingProductAppeal struct {
	EligibleForAppeal   *bool     `json:"eligible_for_appeal,omitempty"`
	ProductAppealStatus *string   `json:"product_appeal_status,omitempty"`
	ProductID           *core.ID  `json:"product_id,omitempty"`
	RejectionReasons    *[]string `json:"rejection_reasons,omitempty"`
	ReviewStatus        *string   `json:"review_status,omitempty"`
}
