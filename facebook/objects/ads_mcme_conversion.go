package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type AdsMcmeConversion struct {
	CreationTime        *time.Time `json:"creation_time,omitempty"`
	Description         *string    `json:"description,omitempty"`
	ID                  *core.ID   `json:"id,omitempty"`
	IsArchived          *bool      `json:"is_archived,omitempty"`
	McmeConversionType  *string    `json:"mcme_conversion_type,omitempty"`
	Name                *string    `json:"name,omitempty"`
	OmnichannelObjectID *core.ID   `json:"omnichannel_object_id,omitempty"`
}
