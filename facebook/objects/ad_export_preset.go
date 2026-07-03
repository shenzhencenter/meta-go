package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdExportPreset struct {
	CreatedTime *core.Time `json:"created_time,omitempty"`
	Fields      *[]string  `json:"fields,omitempty"`
	ID          *core.ID   `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Owner       *User      `json:"owner,omitempty"`
	UpdatedTime *core.Time `json:"updated_time,omitempty"`
}

var AdExportPresetFields = struct {
	CreatedTime string
	Fields      string
	ID          string
	Name        string
	Owner       string
	UpdatedTime string
}{
	CreatedTime: "created_time",
	Fields:      "fields",
	ID:          "id",
	Name:        "name",
	Owner:       "owner",
	UpdatedTime: "updated_time",
}
