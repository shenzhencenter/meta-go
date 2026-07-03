package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BusinessVideo struct {
	Business        *Business `json:"business,omitempty"`
	ID              *core.ID  `json:"id,omitempty"`
	MediaLibraryURL *string   `json:"media_library_url,omitempty"`
	Name            *string   `json:"name,omitempty"`
	Video           *AdVideo  `json:"video,omitempty"`
}

var BusinessVideoFields = struct {
	Business        string
	ID              string
	MediaLibraryURL string
	Name            string
	Video           string
}{
	Business:        "business",
	ID:              "id",
	MediaLibraryURL: "media_library_url",
	Name:            "name",
	Video:           "video",
}
