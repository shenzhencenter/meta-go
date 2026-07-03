package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ALMEndAdvertiserInfo struct {
	EstimatedAdBudget    *int      `json:"estimated_ad_budget,omitempty"`
	ID                   *core.ID  `json:"id,omitempty"`
	ParentAdvertiserID   *core.ID  `json:"parent_advertiser_id,omitempty"`
	ParentAdvertiserName *string   `json:"parent_advertiser_name,omitempty"`
	Tag                  *[]string `json:"tag,omitempty"`
}

var ALMEndAdvertiserInfoFields = struct {
	EstimatedAdBudget    string
	ID                   string
	ParentAdvertiserID   string
	ParentAdvertiserName string
	Tag                  string
}{
	EstimatedAdBudget:    "estimated_ad_budget",
	ID:                   "id",
	ParentAdvertiserID:   "parent_advertiser_id",
	ParentAdvertiserName: "parent_advertiser_name",
	Tag:                  "tag",
}
