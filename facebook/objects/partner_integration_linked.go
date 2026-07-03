package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type PartnerIntegrationLinked struct {
	AdsPixel                     *AdsPixel                 `json:"ads_pixel,omitempty"`
	Application                  *Application              `json:"application,omitempty"`
	CompletedIntegrationTypes    *[]string                 `json:"completed_integration_types,omitempty"`
	ExternalBusinessConnectionID *core.ID                  `json:"external_business_connection_id,omitempty"`
	ExternalID                   *core.ID                  `json:"external_id,omitempty"`
	HasOauthToken                *bool                     `json:"has_oauth_token,omitempty"`
	ID                           *core.ID                  `json:"id,omitempty"`
	MbeAppID                     *core.ID                  `json:"mbe_app_id,omitempty"`
	MbeAssetID                   *core.ID                  `json:"mbe_asset_id,omitempty"`
	MbeExternalBusinessID        *core.ID                  `json:"mbe_external_business_id,omitempty"`
	Name                         *string                   `json:"name,omitempty"`
	OfflineConversionDataSet     *OfflineConversionDataSet `json:"offline_conversion_data_set,omitempty"`
	Page                         *Page                     `json:"page,omitempty"`
	Partner                      *string                   `json:"partner,omitempty"`
	ProductCatalog               *ProductCatalog           `json:"product_catalog,omitempty"`
	SetupStatus                  *string                   `json:"setup_status,omitempty"`
}
