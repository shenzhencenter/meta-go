package enums

type ProductcatalogproductsVisibilityEnumParam string

const (
	ProductcatalogproductsVisibilityEnumParamPublished ProductcatalogproductsVisibilityEnumParam = "published"
	ProductcatalogproductsVisibilityEnumParamStaging   ProductcatalogproductsVisibilityEnumParam = "staging"
)

func (value ProductcatalogproductsVisibilityEnumParam) String() string {
	return string(value)
}
