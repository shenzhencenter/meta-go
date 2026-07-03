package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdCreativeBrandedContentAdsPartners struct {
	FbPageID           *core.ID `json:"fb_page_id,omitempty"`
	HasCreateAdsAccess *bool    `json:"has_create_ads_access,omitempty"`
	IdentityType       *string  `json:"identity_type,omitempty"`
	IgAssetID          *core.ID `json:"ig_asset_id,omitempty"`
	IgUserID           *core.ID `json:"ig_user_id,omitempty"`
}
