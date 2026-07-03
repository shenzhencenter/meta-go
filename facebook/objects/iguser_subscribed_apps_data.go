package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type IGUserSubscribedAppsData struct {
	AppID            *core.ID  `json:"app_id,omitempty"`
	SubscribedFields *[]string `json:"subscribed_fields,omitempty"`
}
