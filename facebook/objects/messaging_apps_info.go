package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type MessagingAppsInfo struct {
	CtdSupportOnlyForIgApp          *bool    `json:"ctd_support_only_for_ig_app,omitempty"`
	HasInstagramMessagingPermission *bool    `json:"has_instagram_messaging_permission,omitempty"`
	HasMessengerMessagingPermission *bool    `json:"has_messenger_messaging_permission,omitempty"`
	ID                              *core.ID `json:"id,omitempty"`
	Name                            *string  `json:"name,omitempty"`
}
