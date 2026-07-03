package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type AdAsyncRequest struct {
	AsyncRequestSet *AdAsyncRequestSet      `json:"async_request_set,omitempty"`
	CreatedTime     *time.Time              `json:"created_time,omitempty"`
	ID              *core.ID                `json:"id,omitempty"`
	Input           *map[string]interface{} `json:"input,omitempty"`
	Result          *map[string]interface{} `json:"result,omitempty"`
	ScopeObjectID   *core.ID                `json:"scope_object_id,omitempty"`
	Status          *string                 `json:"status,omitempty"`
	Type            *string                 `json:"type,omitempty"`
	UpdatedTime     *time.Time              `json:"updated_time,omitempty"`
}
