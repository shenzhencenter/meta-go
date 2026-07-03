package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type WindowsPhoneAppLink struct {
	AppID   *core.ID `json:"app_id,omitempty"`
	AppName *string  `json:"app_name,omitempty"`
	URL     *string  `json:"url,omitempty"`
}

var WindowsPhoneAppLinkFields = struct {
	AppID   string
	AppName string
	URL     string
}{
	AppID:   "app_id",
	AppName: "app_name",
	URL:     "url",
}
