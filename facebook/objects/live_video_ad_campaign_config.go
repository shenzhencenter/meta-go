package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type LiveVideoAdCampaignConfig struct {
	ID              *core.ID `json:"id,omitempty"`
	LiveVideoAdType *string  `json:"live_video_ad_type,omitempty"`
}
