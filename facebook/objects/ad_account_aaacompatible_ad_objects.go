package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAccountAAACompatibleAdObjects struct {
	AdgroupIds       *[]core.ID `json:"adgroup_ids,omitempty"`
	CampaignGroupIds *[]core.ID `json:"campaign_group_ids,omitempty"`
	CampaignIds      *[]core.ID `json:"campaign_ids,omitempty"`
}

var AdAccountAAACompatibleAdObjectsFields = struct {
	AdgroupIds       string
	CampaignGroupIds string
	CampaignIds      string
}{
	AdgroupIds:       "adgroup_ids",
	CampaignGroupIds: "campaign_group_ids",
	CampaignIds:      "campaign_ids",
}
