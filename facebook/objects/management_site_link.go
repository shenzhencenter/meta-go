package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ManagementSiteLink struct {
	AdAccountID   *core.ID `json:"ad_account_id,omitempty"`
	ID            *core.ID `json:"id,omitempty"`
	LinkDomain    *string  `json:"link_domain,omitempty"`
	LinkHash      *string  `json:"link_hash,omitempty"`
	LinkImageHash *string  `json:"link_image_hash,omitempty"`
	LinkImageURL  *string  `json:"link_image_url,omitempty"`
	LinkTitle     *string  `json:"link_title,omitempty"`
	LinkType      *string  `json:"link_type,omitempty"`
	LinkURL       *string  `json:"link_url,omitempty"`
}

var ManagementSiteLinkFields = struct {
	AdAccountID   string
	ID            string
	LinkDomain    string
	LinkHash      string
	LinkImageHash string
	LinkImageURL  string
	LinkTitle     string
	LinkType      string
	LinkURL       string
}{
	AdAccountID:   "ad_account_id",
	ID:            "id",
	LinkDomain:    "link_domain",
	LinkHash:      "link_hash",
	LinkImageHash: "link_image_hash",
	LinkImageURL:  "link_image_url",
	LinkTitle:     "link_title",
	LinkType:      "link_type",
	LinkURL:       "link_url",
}
