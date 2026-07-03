package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type WindowsAppLink struct {
	AppID             *core.ID `json:"app_id,omitempty"`
	AppName           *string  `json:"app_name,omitempty"`
	PackageFamilyName *string  `json:"package_family_name,omitempty"`
	URL               *string  `json:"url,omitempty"`
}
