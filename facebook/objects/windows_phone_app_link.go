package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type WindowsPhoneAppLink struct {
	AppID   *core.ID `json:"app_id,omitempty"`
	AppName *string  `json:"app_name,omitempty"`
	URL     *string  `json:"url,omitempty"`
}
