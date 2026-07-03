package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdLightCampaignGroup struct {
	ID *core.ID `json:"id,omitempty"`
}

var AdLightCampaignGroupFields = struct {
	ID string
}{
	ID: "id",
}
