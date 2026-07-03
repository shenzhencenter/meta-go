package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var CPASMerchantConfigFields = struct {
	AcceptedTos              string
	BetaFeatures             string
	BusinessOutcomesStatus   string
	ID                       string
	IsTestMerchant           string
	OutcomesComplianceStatus string
	QualifiedToOnboard       string
}{
	AcceptedTos:              "accepted_tos",
	BetaFeatures:             "beta_features",
	BusinessOutcomesStatus:   "business_outcomes_status",
	ID:                       "id",
	IsTestMerchant:           "is_test_merchant",
	OutcomesComplianceStatus: "outcomes_compliance_status",
	QualifiedToOnboard:       "qualified_to_onboard",
}
