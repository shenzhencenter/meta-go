package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AudienceSharingAccountValid struct {
	AccountID                         *core.ID `json:"account_id,omitempty"`
	AccountType                       *string  `json:"account_type,omitempty"`
	BusinessID                        *core.ID `json:"business_id,omitempty"`
	BusinessName                      *string  `json:"business_name,omitempty"`
	CanAdAccountUseLookalikeContainer *bool    `json:"can_ad_account_use_lookalike_container,omitempty"`
	SharingAgreementStatus            *int     `json:"sharing_agreement_status,omitempty"`
}
