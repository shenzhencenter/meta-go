package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AREffect struct {
	CreationTime     *core.Time `json:"creation_time,omitempty"`
	ID               *core.ID   `json:"id,omitempty"`
	LastModifiedTime *core.Time `json:"last_modified_time,omitempty"`
	Name             *string    `json:"name,omitempty"`
	Status           *string    `json:"status,omitempty"`
	Surfaces         *[]string  `json:"surfaces,omitempty"`
}

var AREffectFields = struct {
	CreationTime     string
	ID               string
	LastModifiedTime string
	Name             string
	Status           string
	Surfaces         string
}{
	CreationTime:     "creation_time",
	ID:               "id",
	LastModifiedTime: "last_modified_time",
	Name:             "name",
	Status:           "status",
	Surfaces:         "surfaces",
}
