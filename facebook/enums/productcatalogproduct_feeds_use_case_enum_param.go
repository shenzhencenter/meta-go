package enums

type ProductcatalogproductFeedsUseCaseEnumParam string

const (
	ProductcatalogproductFeedsUseCaseEnumParamCreatorAsset ProductcatalogproductFeedsUseCaseEnumParam = "CREATOR_ASSET"
)

func (value ProductcatalogproductFeedsUseCaseEnumParam) String() string {
	return string(value)
}
