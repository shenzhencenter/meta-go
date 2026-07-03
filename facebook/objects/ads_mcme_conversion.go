package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsMcmeConversion struct {
	CreationTime        *core.Time `json:"creation_time,omitempty"`
	Description         *string    `json:"description,omitempty"`
	ID                  *core.ID   `json:"id,omitempty"`
	IsArchived          *bool      `json:"is_archived,omitempty"`
	McmeConversionType  *string    `json:"mcme_conversion_type,omitempty"`
	Name                *string    `json:"name,omitempty"`
	OmnichannelObjectID *core.ID   `json:"omnichannel_object_id,omitempty"`
}

var AdsMcmeConversionFields = struct {
	CreationTime        string
	Description         string
	ID                  string
	IsArchived          string
	McmeConversionType  string
	Name                string
	OmnichannelObjectID string
}{
	CreationTime:        "creation_time",
	Description:         "description",
	ID:                  "id",
	IsArchived:          "is_archived",
	McmeConversionType:  "mcme_conversion_type",
	Name:                "name",
	OmnichannelObjectID: "omnichannel_object_id",
}
