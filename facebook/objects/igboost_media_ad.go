package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type IGBoostMediaAd struct {
	AdID     *core.ID `json:"ad_id,omitempty"`
	AdStatus *string  `json:"ad_status,omitempty"`
}
