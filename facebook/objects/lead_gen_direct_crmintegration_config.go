package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type LeadGenDirectCRMIntegrationConfig struct {
	AuthID              *core.ID             `json:"auth_id,omitempty"`
	CreationTime        *time.Time           `json:"creation_time,omitempty"`
	ID                  *core.ID             `json:"id,omitempty"`
	LeadFilterSettings  *[]map[string]string `json:"lead_filter_settings,omitempty"`
	LeadGenData         *LeadgenForm         `json:"lead_gen_data,omitempty"`
	MatchedFields       *[]map[string]string `json:"matched_fields,omitempty"`
	MatchedFieldsLabels *[]map[string]string `json:"matched_fields_labels,omitempty"`
	Resources           *[]map[string]string `json:"resources,omitempty"`
	ThirdPartyAppID     *core.ID             `json:"third_party_app_id,omitempty"`
}
