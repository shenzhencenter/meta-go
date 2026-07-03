package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsPixelCustomAttributionSource struct {
	ID         *core.ID `json:"id,omitempty"`
	IsEligible *bool    `json:"is_eligible,omitempty"`
	Name       *string  `json:"name,omitempty"`
}
