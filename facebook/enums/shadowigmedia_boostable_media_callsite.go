package enums

type ShadowigmediaBoostableMediaCallsite string

const (
	ShadowigmediaBoostableMediaCallsiteAdsManagerL1EditorDynamicAdsWithExistingPost ShadowigmediaBoostableMediaCallsite = "ADS_MANAGER_L1_EDITOR_DYNAMIC_ADS_WITH_EXISTING_POST"
	ShadowigmediaBoostableMediaCallsitePaHubCatalogIngestionCreatorAsset            ShadowigmediaBoostableMediaCallsite = "PA_HUB_CATALOG_INGESTION_CREATOR_ASSET"
)

func (value ShadowigmediaBoostableMediaCallsite) String() string {
	return string(value)
}
