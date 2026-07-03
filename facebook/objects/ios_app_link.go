package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type IosAppLink struct {
	AppName    *string  `json:"app_name,omitempty"`
	AppStoreID *core.ID `json:"app_store_id,omitempty"`
	URL        *string  `json:"url,omitempty"`
}
