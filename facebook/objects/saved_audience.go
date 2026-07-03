package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
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
	TimeCreated                  *time.Time                    `json:"time_created,omitempty"`
	TimeUpdated                  *time.Time                    `json:"time_updated,omitempty"`
}
