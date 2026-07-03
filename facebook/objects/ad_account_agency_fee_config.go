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
