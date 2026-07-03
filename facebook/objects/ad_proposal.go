package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdProposal struct {
	AdProposalTypeName   *string    `json:"ad_proposal_type_name,omitempty"`
	Adaccount            *AdAccount `json:"adaccount,omitempty"`
	CreationTime         *core.Time `json:"creation_time,omitempty"`
	Creator              *User      `json:"creator,omitempty"`
	DeliveryInterface    *string    `json:"delivery_interface,omitempty"`
	ExpirationTime       *core.Time `json:"expiration_time,omitempty"`
	HasConflict          *bool      `json:"has_conflict,omitempty"`
	ID                   *core.ID   `json:"id,omitempty"`
	KpiMetric            *string    `json:"kpi_metric,omitempty"`
	Message              *string    `json:"message,omitempty"`
	Name                 *string    `json:"name,omitempty"`
	ProposalDtsTemplate  *string    `json:"proposal_dts_template,omitempty"`
	ProposalTemplateName *string    `json:"proposal_template_name,omitempty"`
	Recommendation       *string    `json:"recommendation,omitempty"`
	ReviewTime           *core.Time `json:"review_time,omitempty"`
	ReviewedBy           *User      `json:"reviewed_by,omitempty"`
	SendTime             *core.Time `json:"send_time,omitempty"`
	Status               *string    `json:"status,omitempty"`
	UseTesting           *bool      `json:"use_testing,omitempty"`
}

var AdProposalFields = struct {
	AdProposalTypeName   string
	Adaccount            string
	CreationTime         string
	Creator              string
	DeliveryInterface    string
	ExpirationTime       string
	HasConflict          string
	ID                   string
	KpiMetric            string
	Message              string
	Name                 string
	ProposalDtsTemplate  string
	ProposalTemplateName string
	Recommendation       string
	ReviewTime           string
	ReviewedBy           string
	SendTime             string
	Status               string
	UseTesting           string
}{
	AdProposalTypeName:   "ad_proposal_type_name",
	Adaccount:            "adaccount",
	CreationTime:         "creation_time",
	Creator:              "creator",
	DeliveryInterface:    "delivery_interface",
	ExpirationTime:       "expiration_time",
	HasConflict:          "has_conflict",
	ID:                   "id",
	KpiMetric:            "kpi_metric",
	Message:              "message",
	Name:                 "name",
	ProposalDtsTemplate:  "proposal_dts_template",
	ProposalTemplateName: "proposal_template_name",
	Recommendation:       "recommendation",
	ReviewTime:           "review_time",
	ReviewedBy:           "reviewed_by",
	SendTime:             "send_time",
	Status:               "status",
	UseTesting:           "use_testing",
}
