package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BusinessCreative struct {
	CreationTime *core.Time `json:"creation_time,omitempty"`
	Duration     *int       `json:"duration,omitempty"`
	Hash         *string    `json:"hash,omitempty"`
	Height       *int       `json:"height,omitempty"`
	ID           *core.ID   `json:"id,omitempty"`
	Name         *string    `json:"name,omitempty"`
	Thumbnail    *string    `json:"thumbnail,omitempty"`
	Type         *string    `json:"type,omitempty"`
	URL          *string    `json:"url,omitempty"`
	VideoID      *core.ID   `json:"video_id,omitempty"`
	Width        *int       `json:"width,omitempty"`
}

var BusinessCreativeFields = struct {
	CreationTime string
	Duration     string
	Hash         string
	Height       string
	ID           string
	Name         string
	Thumbnail    string
	Type         string
	URL          string
	VideoID      string
	Width        string
}{
	CreationTime: "creation_time",
	Duration:     "duration",
	Hash:         "hash",
	Height:       "height",
	ID:           "id",
	Name:         "name",
	Thumbnail:    "thumbnail",
	Type:         "type",
	URL:          "url",
	VideoID:      "video_id",
	Width:        "width",
}
