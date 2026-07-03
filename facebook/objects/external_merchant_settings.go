package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type ExternalMerchantSettings struct {
	ConnectWoo       *string  `json:"connect_woo,omitempty"`
	ExternalPlatform *string  `json:"external_platform,omitempty"`
	ID               *core.ID `json:"id,omitempty"`
}
