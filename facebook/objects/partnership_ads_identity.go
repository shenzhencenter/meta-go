package objects

type PartnershipAdsIdentity struct {
	IsRecommended       *bool                     `json:"is_recommended,omitempty"`
	IsSaved             *bool                     `json:"is_saved,omitempty"`
	PostTypes           *[]string                 `json:"post_types,omitempty"`
	SecondaryIdentities *[]map[string]interface{} `json:"secondary_identities,omitempty"`
}
