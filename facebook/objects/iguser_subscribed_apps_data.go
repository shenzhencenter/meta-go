package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type IGUserSubscribedAppsData struct {
	AppID            *core.ID  `json:"app_id,omitempty"`
	SubscribedFields *[]string `json:"subscribed_fields,omitempty"`
}
