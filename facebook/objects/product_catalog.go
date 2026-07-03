package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ProductCatalog struct {
	AdAccountToCollaborativeAdsShareSettings *CollaborativeAdsShareSettings `json:"ad_account_to_collaborative_ads_share_settings,omitempty"`
	AgencyCollaborativeAdsShareSettings      *CollaborativeAdsShareSettings `json:"agency_collaborative_ads_share_settings,omitempty"`
	Business                                 *Business                      `json:"business,omitempty"`
	CatalogStore                             *StoreCatalogSettings          `json:"catalog_store,omitempty"`
	CommerceMerchantSettings                 *CommerceMerchantSettings      `json:"commerce_merchant_settings,omitempty"`
	CreatorUser                              *User                          `json:"creator_user,omitempty"`
	DaDisplaySettings                        *ProductCatalogImageSettings   `json:"da_display_settings,omitempty"`
	DefaultImageURL                          *string                        `json:"default_image_url,omitempty"`
	FallbackImageURL                         *[]string                      `json:"fallback_image_url,omitempty"`
	FeedCount                                *int                           `json:"feed_count,omitempty"`
	ID                                       *core.ID                       `json:"id,omitempty"`
	IsCatalogSegment                         *bool                          `json:"is_catalog_segment,omitempty"`
	IsLocalCatalog                           *bool                          `json:"is_local_catalog,omitempty"`
	Name                                     *string                        `json:"name,omitempty"`
	OwnerBusiness                            *Business                      `json:"owner_business,omitempty"`
	ProductCount                             *int                           `json:"product_count,omitempty"`
	ShowAssignPermissions                    *bool                          `json:"show_assign_permissions,omitempty"`
	StoreCatalogSettings                     *StoreCatalogSettings          `json:"store_catalog_settings,omitempty"`
	UserAccessExpireTime                     *core.Time                     `json:"user_access_expire_time,omitempty"`
	Vertical                                 *string                        `json:"vertical,omitempty"`
}

var ProductCatalogFields = struct {
	AdAccountToCollaborativeAdsShareSettings string
	AgencyCollaborativeAdsShareSettings      string
	Business                                 string
	CatalogStore                             string
	CommerceMerchantSettings                 string
	CreatorUser                              string
	DaDisplaySettings                        string
	DefaultImageURL                          string
	FallbackImageURL                         string
	FeedCount                                string
	ID                                       string
	IsCatalogSegment                         string
	IsLocalCatalog                           string
	Name                                     string
	OwnerBusiness                            string
	ProductCount                             string
	ShowAssignPermissions                    string
	StoreCatalogSettings                     string
	UserAccessExpireTime                     string
	Vertical                                 string
}{
	AdAccountToCollaborativeAdsShareSettings: "ad_account_to_collaborative_ads_share_settings",
	AgencyCollaborativeAdsShareSettings:      "agency_collaborative_ads_share_settings",
	Business:                                 "business",
	CatalogStore:                             "catalog_store",
	CommerceMerchantSettings:                 "commerce_merchant_settings",
	CreatorUser:                              "creator_user",
	DaDisplaySettings:                        "da_display_settings",
	DefaultImageURL:                          "default_image_url",
	FallbackImageURL:                         "fallback_image_url",
	FeedCount:                                "feed_count",
	ID:                                       "id",
	IsCatalogSegment:                         "is_catalog_segment",
	IsLocalCatalog:                           "is_local_catalog",
	Name:                                     "name",
	OwnerBusiness:                            "owner_business",
	ProductCount:                             "product_count",
	ShowAssignPermissions:                    "show_assign_permissions",
	StoreCatalogSettings:                     "store_catalog_settings",
	UserAccessExpireTime:                     "user_access_expire_time",
	Vertical:                                 "vertical",
}
