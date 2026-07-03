package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type CPASMerchantConfig struct {
	AcceptedTos              *bool                                `json:"accepted_tos,omitempty"`
	BetaFeatures             *[]string                            `json:"beta_features,omitempty"`
	BusinessOutcomesStatus   *[]map[string]string                 `json:"business_outcomes_status,omitempty"`
	ID                       *core.ID                             `json:"id,omitempty"`
	IsTestMerchant           *bool                                `json:"is_test_merchant,omitempty"`
	OutcomesComplianceStatus *[]map[string]map[string]interface{} `json:"outcomes_compliance_status,omitempty"`
	QualifiedToOnboard       *bool                                `json:"qualified_to_onboard,omitempty"`
}
