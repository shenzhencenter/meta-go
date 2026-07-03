package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ProductDeliveryPreference struct {
	AdObjectID              *core.ID  `json:"ad_object_id,omitempty"`
	ID                      *core.ID  `json:"id,omitempty"`
	ProductPriority         *string   `json:"product_priority,omitempty"`
	ProductPriorityCategory *[]string `json:"product_priority_category,omitempty"`
}
