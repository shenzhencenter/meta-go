package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CanvasAdSettings struct {
	IsCanvasCollectionEligible *bool    `json:"is_canvas_collection_eligible,omitempty"`
	LeadFormCreatedTime        *uint64  `json:"lead_form_created_time,omitempty"`
	LeadFormName               *string  `json:"lead_form_name,omitempty"`
	LeadGenFormID              *core.ID `json:"lead_gen_form_id,omitempty"`
	LeadsCount                 *int     `json:"leads_count,omitempty"`
	ProductSetID               *core.ID `json:"product_set_id,omitempty"`
	UseRetailerItemIds         *bool    `json:"use_retailer_item_ids,omitempty"`
}

var CanvasAdSettingsFields = struct {
	IsCanvasCollectionEligible string
	LeadFormCreatedTime        string
	LeadFormName               string
	LeadGenFormID              string
	LeadsCount                 string
	ProductSetID               string
	UseRetailerItemIds         string
}{
	IsCanvasCollectionEligible: "is_canvas_collection_eligible",
	LeadFormCreatedTime:        "lead_form_created_time",
	LeadFormName:               "lead_form_name",
	LeadGenFormID:              "lead_gen_form_id",
	LeadsCount:                 "leads_count",
	ProductSetID:               "product_set_id",
	UseRetailerItemIds:         "use_retailer_item_ids",
}
