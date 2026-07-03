package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type Album struct {
	BackdatedTime            *core.Time              `json:"backdated_time,omitempty"`
	BackdatedTimeGranularity *string                 `json:"backdated_time_granularity,omitempty"`
	CanBackdate              *bool                   `json:"can_backdate,omitempty"`
	CanUpload                *bool                   `json:"can_upload,omitempty"`
	Count                    *uint64                 `json:"count,omitempty"`
	CoverPhoto               *Photo                  `json:"cover_photo,omitempty"`
	CreatedTime              *core.Time              `json:"created_time,omitempty"`
	Description              *string                 `json:"description,omitempty"`
	EditLink                 *string                 `json:"edit_link,omitempty"`
	Event                    *Event                  `json:"event,omitempty"`
	From                     *map[string]interface{} `json:"from,omitempty"`
	ID                       *core.ID                `json:"id,omitempty"`
	IsUserFacing             *bool                   `json:"is_user_facing,omitempty"`
	Link                     *string                 `json:"link,omitempty"`
	Location                 *string                 `json:"location,omitempty"`
	ModifiedMajor            *core.Time              `json:"modified_major,omitempty"`
	Name                     *string                 `json:"name,omitempty"`
	PhotoCount               *uint64                 `json:"photo_count,omitempty"`
	Place                    *Place                  `json:"place,omitempty"`
	Privacy                  *string                 `json:"privacy,omitempty"`
	Type                     *string                 `json:"type,omitempty"`
	UpdatedTime              *core.Time              `json:"updated_time,omitempty"`
	VideoCount               *uint64                 `json:"video_count,omitempty"`
}

var AlbumFields = struct {
	BackdatedTime            string
	BackdatedTimeGranularity string
	CanBackdate              string
	CanUpload                string
	Count                    string
	CoverPhoto               string
	CreatedTime              string
	Description              string
	EditLink                 string
	Event                    string
	From                     string
	ID                       string
	IsUserFacing             string
	Link                     string
	Location                 string
	ModifiedMajor            string
	Name                     string
	PhotoCount               string
	Place                    string
	Privacy                  string
	Type                     string
	UpdatedTime              string
	VideoCount               string
}{
	BackdatedTime:            "backdated_time",
	BackdatedTimeGranularity: "backdated_time_granularity",
	CanBackdate:              "can_backdate",
	CanUpload:                "can_upload",
	Count:                    "count",
	CoverPhoto:               "cover_photo",
	CreatedTime:              "created_time",
	Description:              "description",
	EditLink:                 "edit_link",
	Event:                    "event",
	From:                     "from",
	ID:                       "id",
	IsUserFacing:             "is_user_facing",
	Link:                     "link",
	Location:                 "location",
	ModifiedMajor:            "modified_major",
	Name:                     "name",
	PhotoCount:               "photo_count",
	Place:                    "place",
	Privacy:                  "privacy",
	Type:                     "type",
	UpdatedTime:              "updated_time",
	VideoCount:               "video_count",
}
