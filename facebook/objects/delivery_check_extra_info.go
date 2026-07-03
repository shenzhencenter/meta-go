package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type DeliveryCheckExtraInfo struct {
	AdgroupIds  *[]core.ID `json:"adgroup_ids,omitempty"`
	CampaignIds *[]core.ID `json:"campaign_ids,omitempty"`
	Countries   *[]string  `json:"countries,omitempty"`
}

var DeliveryCheckExtraInfoFields = struct {
	AdgroupIds  string
	CampaignIds string
	Countries   string
}{
	AdgroupIds:  "adgroup_ids",
	CampaignIds: "campaign_ids",
	Countries:   "countries",
}
