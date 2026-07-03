package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type WifiInformation struct {
	ID                *core.ID `json:"id,omitempty"`
	Name              *string  `json:"name,omitempty"`
	NetworkAccessType *string  `json:"network_access_type,omitempty"`
}
