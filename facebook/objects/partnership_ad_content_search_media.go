package objects

type PartnershipAdContentSearchMedia struct {
	FbAdCodeSponsorCount                     *int                         `json:"fb_ad_code_sponsor_count,omitempty"`
	FbAdCodeSponsors                         *[]FBPageAndInstagramAccount `json:"fb_ad_code_sponsors,omitempty"`
	IgAdCodeSponsorCount                     *int                         `json:"ig_ad_code_sponsor_count,omitempty"`
	IgAdCodeSponsors                         *[]FBPageAndInstagramAccount `json:"ig_ad_code_sponsors,omitempty"`
	IgMedia                                  *IGMedia                     `json:"ig_media,omitempty"`
	IgMediaHasProductTags                    *bool                        `json:"ig_media_has_product_tags,omitempty"`
	IsAdCodeEligibleForBoostingByTwoSponsors *bool                        `json:"is_ad_code_eligible_for_boosting_by_two_sponsors,omitempty"`
	IsAdCodeEntry                            *bool                        `json:"is_ad_code_entry,omitempty"`
}
