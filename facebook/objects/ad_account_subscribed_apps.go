package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdAccountSubscribedApps struct {
	AppID   *core.ID `json:"app_id,omitempty"`
	AppName *string  `json:"app_name,omitempty"`
}
