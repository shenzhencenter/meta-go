package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type BrandedContentFBPromodeUser struct {
	DelegatePageForAdsOnlyID *core.ID `json:"delegate_page_for_ads_only_id,omitempty"`
	IsIabp                   *bool    `json:"is_iabp,omitempty"`
	IsManaged                *bool    `json:"is_managed,omitempty"`
	Name                     *string  `json:"name,omitempty"`
	ProfilePictureURL        *string  `json:"profile_picture_url,omitempty"`
}
