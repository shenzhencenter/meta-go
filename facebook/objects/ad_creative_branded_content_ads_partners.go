package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCreativeBrandedContentAdsPartners struct {
	FbPageID           *core.ID `json:"fb_page_id,omitempty"`
	HasCreateAdsAccess *bool    `json:"has_create_ads_access,omitempty"`
	IdentityType       *string  `json:"identity_type,omitempty"`
	IgAssetID          *core.ID `json:"ig_asset_id,omitempty"`
	IgUserID           *core.ID `json:"ig_user_id,omitempty"`
}

var AdCreativeBrandedContentAdsPartnersFields = struct {
	FbPageID           string
	HasCreateAdsAccess string
	IdentityType       string
	IgAssetID          string
	IgUserID           string
}{
	FbPageID:           "fb_page_id",
	HasCreateAdsAccess: "has_create_ads_access",
	IdentityType:       "identity_type",
	IgAssetID:          "ig_asset_id",
	IgUserID:           "ig_user_id",
}
