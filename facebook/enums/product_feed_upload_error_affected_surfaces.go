package enums

type ProductfeeduploaderrorAffectedSurfaces string

const (
	ProductfeeduploaderrorAffectedSurfacesDynamicAds    ProductfeeduploaderrorAffectedSurfaces = "Dynamic Ads"
	ProductfeeduploaderrorAffectedSurfacesMarketplace   ProductfeeduploaderrorAffectedSurfaces = "Marketplace"
	ProductfeeduploaderrorAffectedSurfacesUsMarketplace ProductfeeduploaderrorAffectedSurfaces = "US Marketplace"
)

func (value ProductfeeduploaderrorAffectedSurfaces) String() string {
	return string(value)
}
