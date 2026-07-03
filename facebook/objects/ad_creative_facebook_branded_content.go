package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdCreativeFacebookBrandedContent struct {
	SharedToSponsorStatus *string  `json:"shared_to_sponsor_status,omitempty"`
	SponsorPageID         *core.ID `json:"sponsor_page_id,omitempty"`
	SponsorRelationship   *string  `json:"sponsor_relationship,omitempty"`
}
