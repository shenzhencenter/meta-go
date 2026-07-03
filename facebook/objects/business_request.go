package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type BusinessRequest struct {
	Accessor       *Business  `json:"accessor,omitempty"`
	CreationTime   *time.Time `json:"creation_time,omitempty"`
	ID             *core.ID   `json:"id,omitempty"`
	ObjectID       *core.ID   `json:"object_id,omitempty"`
	ObjectType     *string    `json:"object_type,omitempty"`
	PermittedTasks *[]string  `json:"permitted_tasks,omitempty"`
	RequestStatus  *string    `json:"request_status,omitempty"`
	RequestType    *string    `json:"request_type,omitempty"`
	Requestor      *string    `json:"requestor,omitempty"`
}
