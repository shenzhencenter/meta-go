package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
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
