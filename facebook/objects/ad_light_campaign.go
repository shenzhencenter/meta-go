package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdLightCampaign struct {
	CampaignID *core.ID `json:"campaign_id,omitempty"`
	ID         *core.ID `json:"id,omitempty"`
}
