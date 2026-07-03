package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BusinessCreativeFolder struct {
	Business                   *Business            `json:"business,omitempty"`
	CreationTime               *core.Time           `json:"creation_time,omitempty"`
	CreativeInsightPermissions *[]map[string]string `json:"creative_insight_permissions,omitempty"`
	Description                *string              `json:"description,omitempty"`
	ID                         *core.ID             `json:"id,omitempty"`
	MediaLibraryURL            *string              `json:"media_library_url,omitempty"`
	Name                       *string              `json:"name,omitempty"`
	OwnerBusiness              *Business            `json:"owner_business,omitempty"`
}

var BusinessCreativeFolderFields = struct {
	Business                   string
	CreationTime               string
	CreativeInsightPermissions string
	Description                string
	ID                         string
	MediaLibraryURL            string
	Name                       string
	OwnerBusiness              string
}{
	Business:                   "business",
	CreationTime:               "creation_time",
	CreativeInsightPermissions: "creative_insight_permissions",
	Description:                "description",
	ID:                         "id",
	MediaLibraryURL:            "media_library_url",
	Name:                       "name",
	OwnerBusiness:              "owner_business",
}
