package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCreativeFacebookBrandedContent struct {
	SharedToSponsorStatus *string  `json:"shared_to_sponsor_status,omitempty"`
	SponsorPageID         *core.ID `json:"sponsor_page_id,omitempty"`
	SponsorRelationship   *string  `json:"sponsor_relationship,omitempty"`
}

var AdCreativeFacebookBrandedContentFields = struct {
	SharedToSponsorStatus string
	SponsorPageID         string
	SponsorRelationship   string
}{
	SharedToSponsorStatus: "shared_to_sponsor_status",
	SponsorPageID:         "sponsor_page_id",
	SponsorRelationship:   "sponsor_relationship",
}
