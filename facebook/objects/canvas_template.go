package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CanvasTemplate struct {
	Channels              *[]map[string][]map[string]string    `json:"channels,omitempty"`
	Description           *string                              `json:"description,omitempty"`
	Document              *Canvas                              `json:"document,omitempty"`
	ID                    *core.ID                             `json:"id,omitempty"`
	IsMultiTabSupportable *bool                                `json:"is_multi_tab_supportable,omitempty"`
	IsNew                 *bool                                `json:"is_new,omitempty"`
	Name                  *string                              `json:"name,omitempty"`
	Objectives            *[]map[string]map[string]interface{} `json:"objectives,omitempty"`
	OwnerID               *User                                `json:"owner_id,omitempty"`
	RequiredCapabilities  *[]string                            `json:"required_capabilities,omitempty"`
	SnapshotPhoto         *Photo                               `json:"snapshot_photo,omitempty"`
	Status                *string                              `json:"status,omitempty"`
	SubVerticals          *[]string                            `json:"sub_verticals,omitempty"`
	Verticals             *[]map[string]string                 `json:"verticals,omitempty"`
}

var CanvasTemplateFields = struct {
	Channels              string
	Description           string
	Document              string
	ID                    string
	IsMultiTabSupportable string
	IsNew                 string
	Name                  string
	Objectives            string
	OwnerID               string
	RequiredCapabilities  string
	SnapshotPhoto         string
	Status                string
	SubVerticals          string
	Verticals             string
}{
	Channels:              "channels",
	Description:           "description",
	Document:              "document",
	ID:                    "id",
	IsMultiTabSupportable: "is_multi_tab_supportable",
	IsNew:                 "is_new",
	Name:                  "name",
	Objectives:            "objectives",
	OwnerID:               "owner_id",
	RequiredCapabilities:  "required_capabilities",
	SnapshotPhoto:         "snapshot_photo",
	Status:                "status",
	SubVerticals:          "sub_verticals",
	Verticals:             "verticals",
}
