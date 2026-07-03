package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AudienceSharingAccountValid struct {
	AccountID                         *core.ID `json:"account_id,omitempty"`
	AccountType                       *string  `json:"account_type,omitempty"`
	BusinessID                        *core.ID `json:"business_id,omitempty"`
	BusinessName                      *string  `json:"business_name,omitempty"`
	CanAdAccountUseLookalikeContainer *bool    `json:"can_ad_account_use_lookalike_container,omitempty"`
	SharingAgreementStatus            *int     `json:"sharing_agreement_status,omitempty"`
}

var AudienceSharingAccountValidFields = struct {
	AccountID                         string
	AccountType                       string
	BusinessID                        string
	BusinessName                      string
	CanAdAccountUseLookalikeContainer string
	SharingAgreementStatus            string
}{
	AccountID:                         "account_id",
	AccountType:                       "account_type",
	BusinessID:                        "business_id",
	BusinessName:                      "business_name",
	CanAdAccountUseLookalikeContainer: "can_ad_account_use_lookalike_container",
	SharingAgreementStatus:            "sharing_agreement_status",
}
