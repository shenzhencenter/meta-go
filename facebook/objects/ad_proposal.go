package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type AdProposal struct {
	AdProposalTypeName   *string    `json:"ad_proposal_type_name,omitempty"`
	Adaccount            *AdAccount `json:"adaccount,omitempty"`
	CreationTime         *time.Time `json:"creation_time,omitempty"`
	Creator              *User      `json:"creator,omitempty"`
	DeliveryInterface    *string    `json:"delivery_interface,omitempty"`
	ExpirationTime       *time.Time `json:"expiration_time,omitempty"`
	HasConflict          *bool      `json:"has_conflict,omitempty"`
	ID                   *core.ID   `json:"id,omitempty"`
	KpiMetric            *string    `json:"kpi_metric,omitempty"`
	Message              *string    `json:"message,omitempty"`
	Name                 *string    `json:"name,omitempty"`
	ProposalDtsTemplate  *string    `json:"proposal_dts_template,omitempty"`
	ProposalTemplateName *string    `json:"proposal_template_name,omitempty"`
	Recommendation       *string    `json:"recommendation,omitempty"`
	ReviewTime           *time.Time `json:"review_time,omitempty"`
	ReviewedBy           *User      `json:"reviewed_by,omitempty"`
	SendTime             *time.Time `json:"send_time,omitempty"`
	Status               *string    `json:"status,omitempty"`
	UseTesting           *bool      `json:"use_testing,omitempty"`
}
