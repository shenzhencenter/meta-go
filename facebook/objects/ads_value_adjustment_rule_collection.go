package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type AdsValueAdjustmentRuleCollection struct {
	ID               *core.ID   `json:"id,omitempty"`
	IsDefaultSetting *bool      `json:"is_default_setting,omitempty"`
	LastAttachTime   *time.Time `json:"last_attach_time,omitempty"`
	Name             *string    `json:"name,omitempty"`
	ProductType      *string    `json:"product_type,omitempty"`
	Status           *string    `json:"status,omitempty"`
}
