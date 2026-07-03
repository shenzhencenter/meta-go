package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BusinessRequest struct {
	Accessor       *Business  `json:"accessor,omitempty"`
	CreationTime   *core.Time `json:"creation_time,omitempty"`
	ID             *core.ID   `json:"id,omitempty"`
	ObjectID       *core.ID   `json:"object_id,omitempty"`
	ObjectType     *string    `json:"object_type,omitempty"`
	PermittedTasks *[]string  `json:"permitted_tasks,omitempty"`
	RequestStatus  *string    `json:"request_status,omitempty"`
	RequestType    *string    `json:"request_type,omitempty"`
	Requestor      *string    `json:"requestor,omitempty"`
}

var BusinessRequestFields = struct {
	Accessor       string
	CreationTime   string
	ID             string
	ObjectID       string
	ObjectType     string
	PermittedTasks string
	RequestStatus  string
	RequestType    string
	Requestor      string
}{
	Accessor:       "accessor",
	CreationTime:   "creation_time",
	ID:             "id",
	ObjectID:       "object_id",
	ObjectType:     "object_type",
	PermittedTasks: "permitted_tasks",
	RequestStatus:  "request_status",
	RequestType:    "request_type",
	Requestor:      "requestor",
}
