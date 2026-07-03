package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BrandedContentFBPromodeUser struct {
	DelegatePageForAdsOnlyID *core.ID `json:"delegate_page_for_ads_only_id,omitempty"`
	IsIabp                   *bool    `json:"is_iabp,omitempty"`
	IsManaged                *bool    `json:"is_managed,omitempty"`
	Name                     *string  `json:"name,omitempty"`
	ProfilePictureURL        *string  `json:"profile_picture_url,omitempty"`
}

var BrandedContentFBPromodeUserFields = struct {
	DelegatePageForAdsOnlyID string
	IsIabp                   string
	IsManaged                string
	Name                     string
	ProfilePictureURL        string
}{
	DelegatePageForAdsOnlyID: "delegate_page_for_ads_only_id",
	IsIabp:                   "is_iabp",
	IsManaged:                "is_managed",
	Name:                     "name",
	ProfilePictureURL:        "profile_picture_url",
}
