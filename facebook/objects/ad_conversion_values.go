package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdConversionValues struct {
	AdgroupID  *core.ID       `json:"adgroup_id,omitempty"`
	CampaignID *core.ID       `json:"campaign_id,omitempty"`
	Values     *[]interface{} `json:"values,omitempty"`
}
