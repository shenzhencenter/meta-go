package objects

type ProductCatalogImageSettings struct {
	CarouselAd *ProductCatalogImageSettingsOperation `json:"carousel_ad,omitempty"`
	SingleAd   *ProductCatalogImageSettingsOperation `json:"single_ad,omitempty"`
}
