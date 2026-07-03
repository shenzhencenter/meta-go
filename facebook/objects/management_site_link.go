package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
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
