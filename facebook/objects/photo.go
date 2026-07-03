package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type Photo struct {
	Album                    *Album                  `json:"album,omitempty"`
	AltText                  *string                 `json:"alt_text,omitempty"`
	AltTextCustom            *string                 `json:"alt_text_custom,omitempty"`
	BackdatedTime            *time.Time              `json:"backdated_time,omitempty"`
	BackdatedTimeGranularity *string                 `json:"backdated_time_granularity,omitempty"`
	CanBackdate              *bool                   `json:"can_backdate,omitempty"`
	CanDelete                *bool                   `json:"can_delete,omitempty"`
	CanTag                   *bool                   `json:"can_tag,omitempty"`
	CreatedTime              *time.Time              `json:"created_time,omitempty"`
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
	UpdatedTime              *time.Time              `json:"updated_time,omitempty"`
	WebpImages               *[]PlatformImageSource  `json:"webp_images,omitempty"`
	Width                    *uint64                 `json:"width,omitempty"`
}
