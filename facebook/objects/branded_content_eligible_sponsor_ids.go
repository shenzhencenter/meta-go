package objects

type BrandedContentEligibleSponsorIDs struct {
	FbPage           *Page   `json:"fb_page,omitempty"`
	IgAccountV2      *IGUser `json:"ig_account_v2,omitempty"`
	IgApprovalNeeded *bool   `json:"ig_approval_needed,omitempty"`
}
