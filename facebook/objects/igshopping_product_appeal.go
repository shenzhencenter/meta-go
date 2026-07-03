package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type IGShoppingProductAppeal struct {
	EligibleForAppeal   *bool     `json:"eligible_for_appeal,omitempty"`
	ProductAppealStatus *string   `json:"product_appeal_status,omitempty"`
	ProductID           *core.ID  `json:"product_id,omitempty"`
	RejectionReasons    *[]string `json:"rejection_reasons,omitempty"`
	ReviewStatus        *string   `json:"review_status,omitempty"`
}

var IGShoppingProductAppealFields = struct {
	EligibleForAppeal   string
	ProductAppealStatus string
	ProductID           string
	RejectionReasons    string
	ReviewStatus        string
}{
	EligibleForAppeal:   "eligible_for_appeal",
	ProductAppealStatus: "product_appeal_status",
	ProductID:           "product_id",
	RejectionReasons:    "rejection_reasons",
	ReviewStatus:        "review_status",
}
