package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BusinessImage struct {
	Business        *Business  `json:"business,omitempty"`
	CreationTime    *core.Time `json:"creation_time,omitempty"`
	Hash            *string    `json:"hash,omitempty"`
	Height          *int       `json:"height,omitempty"`
	ID              *core.ID   `json:"id,omitempty"`
	MediaLibraryURL *string    `json:"media_library_url,omitempty"`
	Name            *string    `json:"name,omitempty"`
	URL             *string    `json:"url,omitempty"`
	URLX128         *string    `json:"url_128,omitempty"`
	Width           *int       `json:"width,omitempty"`
}

var BusinessImageFields = struct {
	Business        string
	CreationTime    string
	Hash            string
	Height          string
	ID              string
	MediaLibraryURL string
	Name            string
	URL             string
	URLX128         string
	Width           string
}{
	Business:        "business",
	CreationTime:    "creation_time",
	Hash:            "hash",
	Height:          "height",
	ID:              "id",
	MediaLibraryURL: "media_library_url",
	Name:            "name",
	URL:             "url",
	URLX128:         "url_128",
	Width:           "width",
}
