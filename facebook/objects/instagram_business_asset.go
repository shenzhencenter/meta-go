package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type InstagramBusinessAsset struct {
	ID         *core.ID `json:"id,omitempty"`
	IgUserID   *core.ID `json:"ig_user_id,omitempty"`
	IgUsername *string  `json:"ig_username,omitempty"`
}

var InstagramBusinessAssetFields = struct {
	ID         string
	IgUserID   string
	IgUsername string
}{
	ID:         "id",
	IgUserID:   "ig_user_id",
	IgUsername: "ig_username",
}
