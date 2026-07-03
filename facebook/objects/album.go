package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type Album struct {
	BackdatedTime            *time.Time              `json:"backdated_time,omitempty"`
	BackdatedTimeGranularity *string                 `json:"backdated_time_granularity,omitempty"`
	CanBackdate              *bool                   `json:"can_backdate,omitempty"`
	CanUpload                *bool                   `json:"can_upload,omitempty"`
	Count                    *uint64                 `json:"count,omitempty"`
	CoverPhoto               *Photo                  `json:"cover_photo,omitempty"`
	CreatedTime              *time.Time              `json:"created_time,omitempty"`
	Description              *string                 `json:"description,omitempty"`
	EditLink                 *string                 `json:"edit_link,omitempty"`
	Event                    *Event                  `json:"event,omitempty"`
	From                     *map[string]interface{} `json:"from,omitempty"`
	ID                       *core.ID                `json:"id,omitempty"`
	IsUserFacing             *bool                   `json:"is_user_facing,omitempty"`
	Link                     *string                 `json:"link,omitempty"`
	Location                 *string                 `json:"location,omitempty"`
	ModifiedMajor            *time.Time              `json:"modified_major,omitempty"`
	Name                     *string                 `json:"name,omitempty"`
	PhotoCount               *uint64                 `json:"photo_count,omitempty"`
	Place                    *Place                  `json:"place,omitempty"`
	Privacy                  *string                 `json:"privacy,omitempty"`
	Type                     *string                 `json:"type,omitempty"`
	UpdatedTime              *time.Time              `json:"updated_time,omitempty"`
	VideoCount               *uint64                 `json:"video_count,omitempty"`
}
