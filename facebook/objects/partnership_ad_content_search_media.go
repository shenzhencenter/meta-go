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

var PartnershipAdContentSearchMediaFields = struct {
	FbAdCodeSponsorCount                     string
	FbAdCodeSponsors                         string
	IgAdCodeSponsorCount                     string
	IgAdCodeSponsors                         string
	IgMedia                                  string
	IgMediaHasProductTags                    string
	IsAdCodeEligibleForBoostingByTwoSponsors string
	IsAdCodeEntry                            string
}{
	FbAdCodeSponsorCount:                     "fb_ad_code_sponsor_count",
	FbAdCodeSponsors:                         "fb_ad_code_sponsors",
	IgAdCodeSponsorCount:                     "ig_ad_code_sponsor_count",
	IgAdCodeSponsors:                         "ig_ad_code_sponsors",
	IgMedia:                                  "ig_media",
	IgMediaHasProductTags:                    "ig_media_has_product_tags",
	IsAdCodeEligibleForBoostingByTwoSponsors: "is_ad_code_eligible_for_boosting_by_two_sponsors",
	IsAdCodeEntry:                            "is_ad_code_entry",
}
