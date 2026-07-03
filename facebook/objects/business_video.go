package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type BusinessVideo struct {
	Business        *Business `json:"business,omitempty"`
	ID              *core.ID  `json:"id,omitempty"`
	MediaLibraryURL *string   `json:"media_library_url,omitempty"`
	Name            *string   `json:"name,omitempty"`
	Video           *AdVideo  `json:"video,omitempty"`
}
