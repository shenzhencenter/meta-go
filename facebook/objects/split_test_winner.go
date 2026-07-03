package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type SplitTestWinner struct {
	AdObjectLevel    *string               `json:"ad_object_level,omitempty"`
	Confidences      *[]map[string]float64 `json:"confidences,omitempty"`
	WinnerAdObjectID *core.ID              `json:"winner_ad_object_id,omitempty"`
}

var SplitTestWinnerFields = struct {
	AdObjectLevel    string
	Confidences      string
	WinnerAdObjectID string
}{
	AdObjectLevel:    "ad_object_level",
	Confidences:      "confidences",
	WinnerAdObjectID: "winner_ad_object_id",
}
