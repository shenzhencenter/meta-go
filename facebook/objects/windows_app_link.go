package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type WindowsAppLink struct {
	AppID             *core.ID `json:"app_id,omitempty"`
	AppName           *string  `json:"app_name,omitempty"`
	PackageFamilyName *string  `json:"package_family_name,omitempty"`
	URL               *string  `json:"url,omitempty"`
}
