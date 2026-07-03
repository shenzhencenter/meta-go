package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ExternalMerchantSettings struct {
	ConnectWoo       *string  `json:"connect_woo,omitempty"`
	ExternalPlatform *string  `json:"external_platform,omitempty"`
	ID               *core.ID `json:"id,omitempty"`
}

var ExternalMerchantSettingsFields = struct {
	ConnectWoo       string
	ExternalPlatform string
	ID               string
}{
	ConnectWoo:       "connect_woo",
	ExternalPlatform: "external_platform",
	ID:               "id",
}
