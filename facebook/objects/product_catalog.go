package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
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
	UserAccessExpireTime                     *time.Time                     `json:"user_access_expire_time,omitempty"`
	Vertical                                 *string                        `json:"vertical,omitempty"`
}
