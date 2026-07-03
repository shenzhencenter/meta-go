package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type VideoList struct {
	CreationTime *core.Time              `json:"creation_time,omitempty"`
	Description  *string                 `json:"description,omitempty"`
	ID           *core.ID                `json:"id,omitempty"`
	LastModified *core.Time              `json:"last_modified,omitempty"`
	Owner        *map[string]interface{} `json:"owner,omitempty"`
	SeasonNumber *int                    `json:"season_number,omitempty"`
	Thumbnail    *string                 `json:"thumbnail,omitempty"`
	Title        *string                 `json:"title,omitempty"`
	VideosCount  *int                    `json:"videos_count,omitempty"`
}

var VideoListFields = struct {
	CreationTime string
	Description  string
	ID           string
	LastModified string
	Owner        string
	SeasonNumber string
	Thumbnail    string
	Title        string
	VideosCount  string
}{
	CreationTime: "creation_time",
	Description:  "description",
	ID:           "id",
	LastModified: "last_modified",
	Owner:        "owner",
	SeasonNumber: "season_number",
	Thumbnail:    "thumbnail",
	Title:        "title",
	VideosCount:  "videos_count",
}
