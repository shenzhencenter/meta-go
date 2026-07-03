package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CatalogWebsiteSettings struct {
	ID               *core.ID `json:"id,omitempty"`
	IsAllowedToCrawl *bool    `json:"is_allowed_to_crawl,omitempty"`
}

var CatalogWebsiteSettingsFields = struct {
	ID               string
	IsAllowedToCrawl string
}{
	ID:               "id",
	IsAllowedToCrawl: "is_allowed_to_crawl",
}
