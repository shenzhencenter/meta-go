package objects

type BrandedContentEligibleSponsorIDs struct {
	FbPage           *Page   `json:"fb_page,omitempty"`
	IgAccountV2      *IGUser `json:"ig_account_v2,omitempty"`
	IgApprovalNeeded *bool   `json:"ig_approval_needed,omitempty"`
}

var BrandedContentEligibleSponsorIDsFields = struct {
	FbPage           string
	IgAccountV2      string
	IgApprovalNeeded string
}{
	FbPage:           "fb_page",
	IgAccountV2:      "ig_account_v2",
	IgApprovalNeeded: "ig_approval_needed",
}
