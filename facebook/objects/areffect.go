package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type AREffect struct {
	CreationTime     *time.Time `json:"creation_time,omitempty"`
	ID               *core.ID   `json:"id,omitempty"`
	LastModifiedTime *time.Time `json:"last_modified_time,omitempty"`
	Name             *string    `json:"name,omitempty"`
	Status           *string    `json:"status,omitempty"`
	Surfaces         *[]string  `json:"surfaces,omitempty"`
}
