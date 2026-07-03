package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdLightAdgroup struct {
	AdsetID *core.ID `json:"adset_id,omitempty"`
	ID      *core.ID `json:"id,omitempty"`
}
