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

var WindowsAppLinkFields = struct {
	AppID             string
	AppName           string
	PackageFamilyName string
	URL               string
}{
	AppID:             "app_id",
	AppName:           "app_name",
	PackageFamilyName: "package_family_name",
	URL:               "url",
}
