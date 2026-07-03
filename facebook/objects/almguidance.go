package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type ALMGuidance struct {
	AdAccountID          *core.ID                  `json:"ad_account_id,omitempty"`
	Guidances            *[]map[string]interface{} `json:"guidances,omitempty"`
	OpportunityScore     *float64                  `json:"opportunity_score,omitempty"`
	ParentAdvertiserID   *core.ID                  `json:"parent_advertiser_id,omitempty"`
	ParentAdvertiserName *string                   `json:"parent_advertiser_name,omitempty"`
}
