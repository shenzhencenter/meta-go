package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdCreativeWhatsAppChannelSpec struct {
	ChannelID  *core.ID `json:"channel_id,omitempty"`
	ChannelURL *string  `json:"channel_url,omitempty"`
}
