package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ArAdsDataContainer struct {
	CameraFacingOverride *string                   `json:"camera_facing_override,omitempty"`
	CreationTime         *core.Time                `json:"creation_time,omitempty"`
	Effect               *[]map[string]interface{} `json:"effect,omitempty"`
	ID                   *core.ID                  `json:"id,omitempty"`
	IsPublished          *bool                     `json:"is_published,omitempty"`
	LastModifiedTime     *core.Time                `json:"last_modified_time,omitempty"`
	Name                 *string                   `json:"name,omitempty"`
}

var ArAdsDataContainerFields = struct {
	CameraFacingOverride string
	CreationTime         string
	Effect               string
	ID                   string
	IsPublished          string
	LastModifiedTime     string
	Name                 string
}{
	CameraFacingOverride: "camera_facing_override",
	CreationTime:         "creation_time",
	Effect:               "effect",
	ID:                   "id",
	IsPublished:          "is_published",
	LastModifiedTime:     "last_modified_time",
	Name:                 "name",
}
