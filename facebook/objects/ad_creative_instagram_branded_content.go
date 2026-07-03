package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCreativeInstagramBrandedContent struct {
	SponsorID *core.ID `json:"sponsor_id,omitempty"`
}

var AdCreativeInstagramBrandedContentFields = struct {
	SponsorID string
}{
	SponsorID: "sponsor_id",
}
