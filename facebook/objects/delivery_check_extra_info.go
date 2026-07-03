package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type DeliveryCheckExtraInfo struct {
	AdgroupIds  *[]core.ID `json:"adgroup_ids,omitempty"`
	CampaignIds *[]core.ID `json:"campaign_ids,omitempty"`
	Countries   *[]string  `json:"countries,omitempty"`
}
