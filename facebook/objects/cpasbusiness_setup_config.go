package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CPASBusinessSetupConfig struct {
	AcceptedCollabAdsTos         *bool                                `json:"accepted_collab_ads_tos,omitempty"`
	Business                     *Business                            `json:"business,omitempty"`
	BusinessCapabilitiesStatus   *[]map[string]string                 `json:"business_capabilities_status,omitempty"`
	CapabilitiesComplianceStatus *[]map[string]map[string]interface{} `json:"capabilities_compliance_status,omitempty"`
	ID                           *core.ID                             `json:"id,omitempty"`
}
