package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type SavedAudience struct {
	Account                      *AdAccount                    `json:"account,omitempty"`
	ApproximateCountLowerBound   *int                          `json:"approximate_count_lower_bound,omitempty"`
	ApproximateCountUpperBound   *int                          `json:"approximate_count_upper_bound,omitempty"`
	DeleteTime                   *int                          `json:"delete_time,omitempty"`
	Description                  *string                       `json:"description,omitempty"`
	ID                           *core.ID                      `json:"id,omitempty"`
	Name                         *string                       `json:"name,omitempty"`
	OperationStatus              *CustomAudienceStatus         `json:"operation_status,omitempty"`
	OwnerBusiness                *Business                     `json:"owner_business,omitempty"`
	PageDeletionMarkedDeleteTime *int                          `json:"page_deletion_marked_delete_time,omitempty"`
	PermissionForActions         *AudiencePermissionForActions `json:"permission_for_actions,omitempty"`
	RunStatus                    *string                       `json:"run_status,omitempty"`
	SentenceLines                *[]interface{}                `json:"sentence_lines,omitempty"`
	Targeting                    *Targeting                    `json:"targeting,omitempty"`
	TimeCreated                  *core.Time                    `json:"time_created,omitempty"`
	TimeUpdated                  *core.Time                    `json:"time_updated,omitempty"`
}

var SavedAudienceFields = struct {
	Account                      string
	ApproximateCountLowerBound   string
	ApproximateCountUpperBound   string
	DeleteTime                   string
	Description                  string
	ID                           string
	Name                         string
	OperationStatus              string
	OwnerBusiness                string
	PageDeletionMarkedDeleteTime string
	PermissionForActions         string
	RunStatus                    string
	SentenceLines                string
	Targeting                    string
	TimeCreated                  string
	TimeUpdated                  string
}{
	Account:                      "account",
	ApproximateCountLowerBound:   "approximate_count_lower_bound",
	ApproximateCountUpperBound:   "approximate_count_upper_bound",
	DeleteTime:                   "delete_time",
	Description:                  "description",
	ID:                           "id",
	Name:                         "name",
	OperationStatus:              "operation_status",
	OwnerBusiness:                "owner_business",
	PageDeletionMarkedDeleteTime: "page_deletion_marked_delete_time",
	PermissionForActions:         "permission_for_actions",
	RunStatus:                    "run_status",
	SentenceLines:                "sentence_lines",
	Targeting:                    "targeting",
	TimeCreated:                  "time_created",
	TimeUpdated:                  "time_updated",
}
