package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var PartnerIntegrationLinkedFields = struct {
	AdsPixel                     string
	Application                  string
	CompletedIntegrationTypes    string
	ExternalBusinessConnectionID string
	ExternalID                   string
	HasOauthToken                string
	ID                           string
	MbeAppID                     string
	MbeAssetID                   string
	MbeExternalBusinessID        string
	Name                         string
	OfflineConversionDataSet     string
	Page                         string
	Partner                      string
	ProductCatalog               string
	SetupStatus                  string
}{
	AdsPixel:                     "ads_pixel",
	Application:                  "application",
	CompletedIntegrationTypes:    "completed_integration_types",
	ExternalBusinessConnectionID: "external_business_connection_id",
	ExternalID:                   "external_id",
	HasOauthToken:                "has_oauth_token",
	ID:                           "id",
	MbeAppID:                     "mbe_app_id",
	MbeAssetID:                   "mbe_asset_id",
	MbeExternalBusinessID:        "mbe_external_business_id",
	Name:                         "name",
	OfflineConversionDataSet:     "offline_conversion_data_set",
	Page:                         "page",
	Partner:                      "partner",
	ProductCatalog:               "product_catalog",
	SetupStatus:                  "setup_status",
}
