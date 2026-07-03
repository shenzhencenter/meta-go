package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdDraft struct {
	AccountID       *core.ID                `json:"account_id,omitempty"`
	APIVersion      *string                 `json:"api_version,omitempty"`
	AsyncRequestSet *AdAsyncRequestSet      `json:"async_request_set,omitempty"`
	AuthorID        *core.ID                `json:"author_id,omitempty"`
	CreatedBy       *string                 `json:"created_by,omitempty"`
	DraftVersion    *string                 `json:"draft_version,omitempty"`
	ID              *core.ID                `json:"id,omitempty"`
	IsActive        *bool                   `json:"is_active,omitempty"`
	Name            *string                 `json:"name,omitempty"`
	OwnershipType   *string                 `json:"ownership_type,omitempty"`
	PublishStatus   *map[string]interface{} `json:"publish_status,omitempty"`
	State           *string                 `json:"state,omitempty"`
	Summary         *string                 `json:"summary,omitempty"`
	TimeCreated     *core.Time              `json:"time_created,omitempty"`
	TimeUpdated     *core.Time              `json:"time_updated,omitempty"`
}

var AdDraftFields = struct {
	AccountID       string
	APIVersion      string
	AsyncRequestSet string
	AuthorID        string
	CreatedBy       string
	DraftVersion    string
	ID              string
	IsActive        string
	Name            string
	OwnershipType   string
	PublishStatus   string
	State           string
	Summary         string
	TimeCreated     string
	TimeUpdated     string
}{
	AccountID:       "account_id",
	APIVersion:      "api_version",
	AsyncRequestSet: "async_request_set",
	AuthorID:        "author_id",
	CreatedBy:       "created_by",
	DraftVersion:    "draft_version",
	ID:              "id",
	IsActive:        "is_active",
	Name:            "name",
	OwnershipType:   "ownership_type",
	PublishStatus:   "publish_status",
	State:           "state",
	Summary:         "summary",
	TimeCreated:     "time_created",
	TimeUpdated:     "time_updated",
}
