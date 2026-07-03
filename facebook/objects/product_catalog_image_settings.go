package objects

type ProductCatalogImageSettings struct {
	CarouselAd *ProductCatalogImageSettingsOperation `json:"carousel_ad,omitempty"`
	SingleAd   *ProductCatalogImageSettingsOperation `json:"single_ad,omitempty"`
}

var ProductCatalogImageSettingsFields = struct {
	CarouselAd string
	SingleAd   string
}{
	CarouselAd: "carousel_ad",
	SingleAd:   "single_ad",
}
