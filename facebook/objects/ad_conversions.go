package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdConversions struct {
	AccountID  *core.ID       `json:"account_id,omitempty"`
	AdgroupID  *core.ID       `json:"adgroup_id,omitempty"`
	CampaignID *core.ID       `json:"campaign_id,omitempty"`
	Values     *[]interface{} `json:"values,omitempty"`
}
