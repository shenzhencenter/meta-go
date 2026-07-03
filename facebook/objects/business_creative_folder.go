package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type BusinessCreativeFolder struct {
	Business                   *Business            `json:"business,omitempty"`
	CreationTime               *time.Time           `json:"creation_time,omitempty"`
	CreativeInsightPermissions *[]map[string]string `json:"creative_insight_permissions,omitempty"`
	Description                *string              `json:"description,omitempty"`
	ID                         *core.ID             `json:"id,omitempty"`
	MediaLibraryURL            *string              `json:"media_library_url,omitempty"`
	Name                       *string              `json:"name,omitempty"`
	OwnerBusiness              *Business            `json:"owner_business,omitempty"`
}
