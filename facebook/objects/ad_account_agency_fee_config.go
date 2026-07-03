package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AdAccountAgencyFeeConfig struct {
	CanAddAgencyFeeToInvoice *bool                                 `json:"can_add_agency_fee_to_invoice,omitempty"`
	DefaultAgencyFeePct      *float64                              `json:"default_agency_fee_pct,omitempty"`
	ID                       *core.ID                              `json:"id,omitempty"`
	IsAgencyFeeDisabled      *bool                                 `json:"is_agency_fee_disabled,omitempty"`
	Status                   *enums.AdaccountagencyfeeconfigStatus `json:"status,omitempty"`
}

var AdAccountAgencyFeeConfigFields = struct {
	CanAddAgencyFeeToInvoice string
	DefaultAgencyFeePct      string
	ID                       string
	IsAgencyFeeDisabled      string
	Status                   string
}{
	CanAddAgencyFeeToInvoice: "can_add_agency_fee_to_invoice",
	DefaultAgencyFeePct:      "default_agency_fee_pct",
	ID:                       "id",
	IsAgencyFeeDisabled:      "is_agency_fee_disabled",
	Status:                   "status",
}
