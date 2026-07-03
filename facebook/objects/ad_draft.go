package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
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
	TimeCreated     *time.Time              `json:"time_created,omitempty"`
	TimeUpdated     *time.Time              `json:"time_updated,omitempty"`
}
