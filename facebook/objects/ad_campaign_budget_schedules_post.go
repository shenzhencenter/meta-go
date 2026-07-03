package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCampaignBudgetSchedulesPost struct {
	ID *core.ID `json:"id,omitempty"`
}

var AdCampaignBudgetSchedulesPostFields = struct {
	ID string
}{
	ID: "id",
}
