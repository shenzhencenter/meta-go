package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type AdExportPreset struct {
	CreatedTime *time.Time `json:"created_time,omitempty"`
	Fields      *[]string  `json:"fields,omitempty"`
	ID          *core.ID   `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Owner       *User      `json:"owner,omitempty"`
	UpdatedTime *time.Time `json:"updated_time,omitempty"`
}
