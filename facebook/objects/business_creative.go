package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type BusinessCreative struct {
	CreationTime *time.Time `json:"creation_time,omitempty"`
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
