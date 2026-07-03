package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdAccountAAACompatibleAdObjects struct {
	AdgroupIds       *[]core.ID `json:"adgroup_ids,omitempty"`
	CampaignGroupIds *[]core.ID `json:"campaign_group_ids,omitempty"`
	CampaignIds      *[]core.ID `json:"campaign_ids,omitempty"`
}
