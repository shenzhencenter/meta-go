package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type SocialWifiSite struct {
	ID *core.ID `json:"id,omitempty"`
}

var SocialWifiSiteFields = struct {
	ID string
}{
	ID: "id",
}
