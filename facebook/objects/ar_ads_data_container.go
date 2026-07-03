package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type ArAdsDataContainer struct {
	CameraFacingOverride *string                   `json:"camera_facing_override,omitempty"`
	CreationTime         *time.Time                `json:"creation_time,omitempty"`
	Effect               *[]map[string]interface{} `json:"effect,omitempty"`
	ID                   *core.ID                  `json:"id,omitempty"`
	IsPublished          *bool                     `json:"is_published,omitempty"`
	LastModifiedTime     *time.Time                `json:"last_modified_time,omitempty"`
	Name                 *string                   `json:"name,omitempty"`
}
