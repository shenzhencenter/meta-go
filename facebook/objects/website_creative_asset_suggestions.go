package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type WebsiteCreativeAssetSuggestions struct {
	AdAccountID      *core.ID `json:"ad_account_id,omitempty"`
	ExtractionStatus *string  `json:"extraction_status,omitempty"`
	ID               *core.ID `json:"id,omitempty"`
	LinkURL          *string  `json:"link_url,omitempty"`
}

var WebsiteCreativeAssetSuggestionsFields = struct {
	AdAccountID      string
	ExtractionStatus string
	ID               string
	LinkURL          string
}{
	AdAccountID:      "ad_account_id",
	ExtractionStatus: "extraction_status",
	ID:               "id",
	LinkURL:          "link_url",
}
