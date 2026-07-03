package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type VideoList struct {
	CreationTime *time.Time              `json:"creation_time,omitempty"`
	Description  *string                 `json:"description,omitempty"`
	ID           *core.ID                `json:"id,omitempty"`
	LastModified *time.Time              `json:"last_modified,omitempty"`
	Owner        *map[string]interface{} `json:"owner,omitempty"`
	SeasonNumber *int                    `json:"season_number,omitempty"`
	Thumbnail    *string                 `json:"thumbnail,omitempty"`
	Title        *string                 `json:"title,omitempty"`
	VideosCount  *int                    `json:"videos_count,omitempty"`
}
