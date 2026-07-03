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

var MessagingAppsInfoFields = struct {
	CtdSupportOnlyForIgApp          string
	HasInstagramMessagingPermission string
	HasMessengerMessagingPermission string
	ID                              string
	Name                            string
}{
	CtdSupportOnlyForIgApp:          "ctd_support_only_for_ig_app",
	HasInstagramMessagingPermission: "has_instagram_messaging_permission",
	HasMessengerMessagingPermission: "has_messenger_messaging_permission",
	ID:                              "id",
	Name:                            "name",
}
