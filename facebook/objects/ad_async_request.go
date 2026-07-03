package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAsyncRequest struct {
	AsyncRequestSet *AdAsyncRequestSet      `json:"async_request_set,omitempty"`
	CreatedTime     *core.Time              `json:"created_time,omitempty"`
	ID              *core.ID                `json:"id,omitempty"`
	Input           *map[string]interface{} `json:"input,omitempty"`
	Result          *map[string]interface{} `json:"result,omitempty"`
	ScopeObjectID   *core.ID                `json:"scope_object_id,omitempty"`
	Status          *string                 `json:"status,omitempty"`
	Type            *string                 `json:"type,omitempty"`
	UpdatedTime     *core.Time              `json:"updated_time,omitempty"`
}

var AdAsyncRequestFields = struct {
	AsyncRequestSet string
	CreatedTime     string
	ID              string
	Input           string
	Result          string
	ScopeObjectID   string
	Status          string
	Type            string
	UpdatedTime     string
}{
	AsyncRequestSet: "async_request_set",
	CreatedTime:     "created_time",
	ID:              "id",
	Input:           "input",
	Result:          "result",
	ScopeObjectID:   "scope_object_id",
	Status:          "status",
	Type:            "type",
	UpdatedTime:     "updated_time",
}
