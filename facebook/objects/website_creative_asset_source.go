package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type WebsiteCreativeAssetSource struct {
	ID        *core.ID `json:"id,omitempty"`
	SourceURL *string  `json:"source_url,omitempty"`
}

var WebsiteCreativeAssetSourceFields = struct {
	ID        string
	SourceURL string
}{
	ID:        "id",
	SourceURL: "source_url",
}
