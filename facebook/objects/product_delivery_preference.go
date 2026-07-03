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

var ProductDeliveryPreferenceFields = struct {
	AdObjectID              string
	ID                      string
	ProductPriority         string
	ProductPriorityCategory string
}{
	AdObjectID:              "ad_object_id",
	ID:                      "id",
	ProductPriority:         "product_priority",
	ProductPriorityCategory: "product_priority_category",
}
