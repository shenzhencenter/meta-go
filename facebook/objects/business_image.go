package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type BusinessImage struct {
	Business        *Business  `json:"business,omitempty"`
	CreationTime    *time.Time `json:"creation_time,omitempty"`
	Hash            *string    `json:"hash,omitempty"`
	Height          *int       `json:"height,omitempty"`
	ID              *core.ID   `json:"id,omitempty"`
	MediaLibraryURL *string    `json:"media_library_url,omitempty"`
	Name            *string    `json:"name,omitempty"`
	URL             *string    `json:"url,omitempty"`
	URLX128         *string    `json:"url_128,omitempty"`
	Width           *int       `json:"width,omitempty"`
}
