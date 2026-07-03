package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdConversions struct {
	AccountID  *core.ID       `json:"account_id,omitempty"`
	AdgroupID  *core.ID       `json:"adgroup_id,omitempty"`
	CampaignID *core.ID       `json:"campaign_id,omitempty"`
	Values     *[]interface{} `json:"values,omitempty"`
}

var AdConversionsFields = struct {
	AccountID  string
	AdgroupID  string
	CampaignID string
	Values     string
}{
	AccountID:  "account_id",
	AdgroupID:  "adgroup_id",
	CampaignID: "campaign_id",
	Values:     "values",
}
