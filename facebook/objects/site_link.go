package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type SiteLink struct {
	ID            *core.ID `json:"id,omitempty"`
	LinkImageHash *string  `json:"link_image_hash,omitempty"`
	LinkTitle     *string  `json:"link_title,omitempty"`
	LinkType      *string  `json:"link_type,omitempty"`
	LinkURL       *string  `json:"link_url,omitempty"`
}

var SiteLinkFields = struct {
	ID            string
	LinkImageHash string
	LinkTitle     string
	LinkType      string
	LinkURL       string
}{
	ID:            "id",
	LinkImageHash: "link_image_hash",
	LinkTitle:     "link_title",
	LinkType:      "link_type",
	LinkURL:       "link_url",
}
