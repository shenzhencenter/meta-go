package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type VideoGroup struct {
	CreatedTime     *string    `json:"created_time,omitempty"`
	Description     *string    `json:"description,omitempty"`
	DisableReason   *string    `json:"disable_reason,omitempty"`
	ID              *core.ID   `json:"id,omitempty"`
	IgProfileIds    *[]core.ID `json:"ig_profile_ids,omitempty"`
	IsDisabled      *bool      `json:"is_disabled,omitempty"`
	IsFbVideoGroup  *bool      `json:"is_fb_video_group,omitempty"`
	LastUsedTime    *string    `json:"last_used_time,omitempty"`
	Length          *float64   `json:"length,omitempty"`
	Name            *string    `json:"name,omitempty"`
	PageID          *core.ID   `json:"page_id,omitempty"`
	PageIds         *[]core.ID `json:"page_ids,omitempty"`
	Picture         *string    `json:"picture,omitempty"`
	Placements      *[]string  `json:"placements,omitempty"`
	VideoGroupTypes *[]string  `json:"video_group_types,omitempty"`
	Videos          *[]string  `json:"videos,omitempty"`
	Views           *uint64    `json:"views,omitempty"`
}

var VideoGroupFields = struct {
	CreatedTime     string
	Description     string
	DisableReason   string
	ID              string
	IgProfileIds    string
	IsDisabled      string
	IsFbVideoGroup  string
	LastUsedTime    string
	Length          string
	Name            string
	PageID          string
	PageIds         string
	Picture         string
	Placements      string
	VideoGroupTypes string
	Videos          string
	Views           string
}{
	CreatedTime:     "created_time",
	Description:     "description",
	DisableReason:   "disable_reason",
	ID:              "id",
	IgProfileIds:    "ig_profile_ids",
	IsDisabled:      "is_disabled",
	IsFbVideoGroup:  "is_fb_video_group",
	LastUsedTime:    "last_used_time",
	Length:          "length",
	Name:            "name",
	PageID:          "page_id",
	PageIds:         "page_ids",
	Picture:         "picture",
	Placements:      "placements",
	VideoGroupTypes: "video_group_types",
	Videos:          "videos",
	Views:           "views",
}
