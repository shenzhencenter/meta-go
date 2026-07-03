package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAssetFeedSpecEvents struct {
	ID *core.ID `json:"id,omitempty"`
}

var AdAssetFeedSpecEventsFields = struct {
	ID string
}{
	ID: "id",
}
