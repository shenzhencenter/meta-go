package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type WebsiteCreativeInfo struct {
	ID        *core.ID  `json:"id,omitempty"`
	ImageUrls *[]string `json:"image_urls,omitempty"`
	LinkURL   *string   `json:"link_url,omitempty"`
}
