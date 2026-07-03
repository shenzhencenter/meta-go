package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type LeadGenDirectCRMIntegrationConfig struct {
	AuthID              *core.ID             `json:"auth_id,omitempty"`
	CreationTime        *core.Time           `json:"creation_time,omitempty"`
	ID                  *core.ID             `json:"id,omitempty"`
	LeadFilterSettings  *[]map[string]string `json:"lead_filter_settings,omitempty"`
	LeadGenData         *LeadgenForm         `json:"lead_gen_data,omitempty"`
	MatchedFields       *[]map[string]string `json:"matched_fields,omitempty"`
	MatchedFieldsLabels *[]map[string]string `json:"matched_fields_labels,omitempty"`
	Resources           *[]map[string]string `json:"resources,omitempty"`
	ThirdPartyAppID     *core.ID             `json:"third_party_app_id,omitempty"`
}

var LeadGenDirectCRMIntegrationConfigFields = struct {
	AuthID              string
	CreationTime        string
	ID                  string
	LeadFilterSettings  string
	LeadGenData         string
	MatchedFields       string
	MatchedFieldsLabels string
	Resources           string
	ThirdPartyAppID     string
}{
	AuthID:              "auth_id",
	CreationTime:        "creation_time",
	ID:                  "id",
	LeadFilterSettings:  "lead_filter_settings",
	LeadGenData:         "lead_gen_data",
	MatchedFields:       "matched_fields",
	MatchedFieldsLabels: "matched_fields_labels",
	Resources:           "resources",
	ThirdPartyAppID:     "third_party_app_id",
}
