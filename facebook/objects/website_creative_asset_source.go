package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type WebsiteCreativeAssetSource struct {
	ID        *core.ID `json:"id,omitempty"`
	SourceURL *string  `json:"source_url,omitempty"`
}
