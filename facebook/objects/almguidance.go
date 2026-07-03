package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ALMGuidance struct {
	AdAccountID          *core.ID                  `json:"ad_account_id,omitempty"`
	Guidances            *[]map[string]interface{} `json:"guidances,omitempty"`
	OpportunityScore     *float64                  `json:"opportunity_score,omitempty"`
	ParentAdvertiserID   *core.ID                  `json:"parent_advertiser_id,omitempty"`
	ParentAdvertiserName *string                   `json:"parent_advertiser_name,omitempty"`
}

var ALMGuidanceFields = struct {
	AdAccountID          string
	Guidances            string
	OpportunityScore     string
	ParentAdvertiserID   string
	ParentAdvertiserName string
}{
	AdAccountID:          "ad_account_id",
	Guidances:            "guidances",
	OpportunityScore:     "opportunity_score",
	ParentAdvertiserID:   "parent_advertiser_id",
	ParentAdvertiserName: "parent_advertiser_name",
}
