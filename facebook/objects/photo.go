package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type Photo struct {
	Album                    *Album                  `json:"album,omitempty"`
	AltText                  *string                 `json:"alt_text,omitempty"`
	AltTextCustom            *string                 `json:"alt_text_custom,omitempty"`
	BackdatedTime            *core.Time              `json:"backdated_time,omitempty"`
	BackdatedTimeGranularity *string                 `json:"backdated_time_granularity,omitempty"`
	CanBackdate              *bool                   `json:"can_backdate,omitempty"`
	CanDelete                *bool                   `json:"can_delete,omitempty"`
	CanTag                   *bool                   `json:"can_tag,omitempty"`
	CreatedTime              *core.Time              `json:"created_time,omitempty"`
	Event                    *Event                  `json:"event,omitempty"`
	From                     *map[string]interface{} `json:"from,omitempty"`
	Height                   *uint64                 `json:"height,omitempty"`
	Icon                     *string                 `json:"icon,omitempty"`
	ID                       *core.ID                `json:"id,omitempty"`
	Images                   *[]PlatformImageSource  `json:"images,omitempty"`
	Link                     *string                 `json:"link,omitempty"`
	Name                     *string                 `json:"name,omitempty"`
	NameTags                 *[]EntityAtTextRange    `json:"name_tags,omitempty"`
	PageStoryID              *core.ID                `json:"page_story_id,omitempty"`
	Picture                  *string                 `json:"picture,omitempty"`
	Place                    *Place                  `json:"place,omitempty"`
	Position                 *uint64                 `json:"position,omitempty"`
	Source                   *string                 `json:"source,omitempty"`
	Target                   *Profile                `json:"target,omitempty"`
	UpdatedTime              *core.Time              `json:"updated_time,omitempty"`
	WebpImages               *[]PlatformImageSource  `json:"webp_images,omitempty"`
	Width                    *uint64                 `json:"width,omitempty"`
}

var PhotoFields = struct {
	Album                    string
	AltText                  string
	AltTextCustom            string
	BackdatedTime            string
	BackdatedTimeGranularity string
	CanBackdate              string
	CanDelete                string
	CanTag                   string
	CreatedTime              string
	Event                    string
	From                     string
	Height                   string
	Icon                     string
	ID                       string
	Images                   string
	Link                     string
	Name                     string
	NameTags                 string
	PageStoryID              string
	Picture                  string
	Place                    string
	Position                 string
	Source                   string
	Target                   string
	UpdatedTime              string
	WebpImages               string
	Width                    string
}{
	Album:                    "album",
	AltText:                  "alt_text",
	AltTextCustom:            "alt_text_custom",
	BackdatedTime:            "backdated_time",
	BackdatedTimeGranularity: "backdated_time_granularity",
	CanBackdate:              "can_backdate",
	CanDelete:                "can_delete",
	CanTag:                   "can_tag",
	CreatedTime:              "created_time",
	Event:                    "event",
	From:                     "from",
	Height:                   "height",
	Icon:                     "icon",
	ID:                       "id",
	Images:                   "images",
	Link:                     "link",
	Name:                     "name",
	NameTags:                 "name_tags",
	PageStoryID:              "page_story_id",
	Picture:                  "picture",
	Place:                    "place",
	Position:                 "position",
	Source:                   "source",
	Target:                   "target",
	UpdatedTime:              "updated_time",
	WebpImages:               "webp_images",
	Width:                    "width",
}
