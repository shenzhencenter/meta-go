package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type AdsPivotRules struct {
	CreationTime *time.Time                `json:"creation_time,omitempty"`
	Creator      *Profile                  `json:"creator,omitempty"`
	Description  *string                   `json:"description,omitempty"`
	ID           *core.ID                  `json:"id,omitempty"`
	Name         *string                   `json:"name,omitempty"`
	Permission   *string                   `json:"permission,omitempty"`
	Rules        *[]map[string]interface{} `json:"rules,omitempty"`
	Scope        *string                   `json:"scope,omitempty"`
	UpdateBy     *Profile                  `json:"update_by,omitempty"`
	UpdateTime   *time.Time                `json:"update_time,omitempty"`
}
