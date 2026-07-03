package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdCampaignGroupStructureTree struct {
	Children *[]AdCampaignGroupStructureTree `json:"children,omitempty"`
	ID       *core.ID                        `json:"id,omitempty"`
	Name     *string                         `json:"name,omitempty"`
	Status   *string                         `json:"status,omitempty"`
}
