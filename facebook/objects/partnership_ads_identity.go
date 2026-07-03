package objects

type PartnershipAdsIdentity struct {
	IsRecommended       *bool                     `json:"is_recommended,omitempty"`
	IsSaved             *bool                     `json:"is_saved,omitempty"`
	PostTypes           *[]string                 `json:"post_types,omitempty"`
	SecondaryIdentities *[]map[string]interface{} `json:"secondary_identities,omitempty"`
}

var PartnershipAdsIdentityFields = struct {
	IsRecommended       string
	IsSaved             string
	PostTypes           string
	SecondaryIdentities string
}{
	IsRecommended:       "is_recommended",
	IsSaved:             "is_saved",
	PostTypes:           "post_types",
	SecondaryIdentities: "secondary_identities",
}
