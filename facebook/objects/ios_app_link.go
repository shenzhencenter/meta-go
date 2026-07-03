package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type IosAppLink struct {
	AppName    *string  `json:"app_name,omitempty"`
	AppStoreID *core.ID `json:"app_store_id,omitempty"`
	URL        *string  `json:"url,omitempty"`
}

var IosAppLinkFields = struct {
	AppName    string
	AppStoreID string
	URL        string
}{
	AppName:    "app_name",
	AppStoreID: "app_store_id",
	URL:        "url",
}
