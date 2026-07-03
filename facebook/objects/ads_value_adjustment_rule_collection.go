package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsValueAdjustmentRuleCollection struct {
	ID               *core.ID   `json:"id,omitempty"`
	IsDefaultSetting *bool      `json:"is_default_setting,omitempty"`
	LastAttachTime   *core.Time `json:"last_attach_time,omitempty"`
	Name             *string    `json:"name,omitempty"`
	ProductType      *string    `json:"product_type,omitempty"`
	Status           *string    `json:"status,omitempty"`
}

var AdsValueAdjustmentRuleCollectionFields = struct {
	ID               string
	IsDefaultSetting string
	LastAttachTime   string
	Name             string
	ProductType      string
	Status           string
}{
	ID:               "id",
	IsDefaultSetting: "is_default_setting",
	LastAttachTime:   "last_attach_time",
	Name:             "name",
	ProductType:      "product_type",
	Status:           "status",
}
