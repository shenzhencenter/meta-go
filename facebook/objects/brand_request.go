package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type BrandRequest struct {
	AdCountries        *[]string                 `json:"ad_countries,omitempty"`
	AdditionalContacts *[]string                 `json:"additional_contacts,omitempty"`
	ApprovalLevel      *uint64                   `json:"approval_level,omitempty"`
	Cells              *[]map[string]interface{} `json:"cells,omitempty"`
	Countries          *[]string                 `json:"countries,omitempty"`
	DenyReason         *string                   `json:"deny_reason,omitempty"`
	EndTime            *time.Time                `json:"end_time,omitempty"`
	EstimatedReach     *uint64                   `json:"estimated_reach,omitempty"`
	ID                 *core.ID                  `json:"id,omitempty"`
	IsMulticell        *bool                     `json:"is_multicell,omitempty"`
	Locale             *string                   `json:"locale,omitempty"`
	MaxAge             *uint64                   `json:"max_age,omitempty"`
	MinAge             *uint64                   `json:"min_age,omitempty"`
	Questions          *[]map[string]interface{} `json:"questions,omitempty"`
	Region             *string                   `json:"region,omitempty"`
	RequestStatus      *string                   `json:"request_status,omitempty"`
	ReviewDate         *time.Time                `json:"review_date,omitempty"`
	StartTime          *time.Time                `json:"start_time,omitempty"`
	Status             *string                   `json:"status,omitempty"`
	SubmitDate         *time.Time                `json:"submit_date,omitempty"`
	TotalBudget        *uint64                   `json:"total_budget,omitempty"`
}
