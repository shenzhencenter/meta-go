package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BrandRequest struct {
	AdCountries        *[]string                 `json:"ad_countries,omitempty"`
	AdditionalContacts *[]string                 `json:"additional_contacts,omitempty"`
	ApprovalLevel      *uint64                   `json:"approval_level,omitempty"`
	Cells              *[]map[string]interface{} `json:"cells,omitempty"`
	Countries          *[]string                 `json:"countries,omitempty"`
	DenyReason         *string                   `json:"deny_reason,omitempty"`
	EndTime            *core.Time                `json:"end_time,omitempty"`
	EstimatedReach     *uint64                   `json:"estimated_reach,omitempty"`
	ID                 *core.ID                  `json:"id,omitempty"`
	IsMulticell        *bool                     `json:"is_multicell,omitempty"`
	Locale             *string                   `json:"locale,omitempty"`
	MaxAge             *uint64                   `json:"max_age,omitempty"`
	MinAge             *uint64                   `json:"min_age,omitempty"`
	Questions          *[]map[string]interface{} `json:"questions,omitempty"`
	Region             *string                   `json:"region,omitempty"`
	RequestStatus      *string                   `json:"request_status,omitempty"`
	ReviewDate         *core.Time                `json:"review_date,omitempty"`
	StartTime          *core.Time                `json:"start_time,omitempty"`
	Status             *string                   `json:"status,omitempty"`
	SubmitDate         *core.Time                `json:"submit_date,omitempty"`
	TotalBudget        *uint64                   `json:"total_budget,omitempty"`
}

var BrandRequestFields = struct {
	AdCountries        string
	AdditionalContacts string
	ApprovalLevel      string
	Cells              string
	Countries          string
	DenyReason         string
	EndTime            string
	EstimatedReach     string
	ID                 string
	IsMulticell        string
	Locale             string
	MaxAge             string
	MinAge             string
	Questions          string
	Region             string
	RequestStatus      string
	ReviewDate         string
	StartTime          string
	Status             string
	SubmitDate         string
	TotalBudget        string
}{
	AdCountries:        "ad_countries",
	AdditionalContacts: "additional_contacts",
	ApprovalLevel:      "approval_level",
	Cells:              "cells",
	Countries:          "countries",
	DenyReason:         "deny_reason",
	EndTime:            "end_time",
	EstimatedReach:     "estimated_reach",
	ID:                 "id",
	IsMulticell:        "is_multicell",
	Locale:             "locale",
	MaxAge:             "max_age",
	MinAge:             "min_age",
	Questions:          "questions",
	Region:             "region",
	RequestStatus:      "request_status",
	ReviewDate:         "review_date",
	StartTime:          "start_time",
	Status:             "status",
	SubmitDate:         "submit_date",
	TotalBudget:        "total_budget",
}
