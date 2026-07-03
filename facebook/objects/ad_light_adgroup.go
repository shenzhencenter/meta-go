package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdLightAdgroup struct {
	AdsetID *core.ID `json:"adset_id,omitempty"`
	ID      *core.ID `json:"id,omitempty"`
}

var AdLightAdgroupFields = struct {
	AdsetID string
	ID      string
}{
	AdsetID: "adset_id",
	ID:      "id",
}
