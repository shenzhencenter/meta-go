package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type URL struct {
	Engagement           *map[string]interface{} `json:"engagement,omitempty"`
	ID                   *core.ID                `json:"id,omitempty"`
	OgObject             *map[string]interface{} `json:"og_object,omitempty"`
	OwnershipPermissions *map[string]interface{} `json:"ownership_permissions,omitempty"`
	Scopes               *map[string]interface{} `json:"scopes,omitempty"`
}
