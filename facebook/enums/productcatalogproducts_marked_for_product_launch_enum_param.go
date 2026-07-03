package enums

type ProductcatalogproductsMarkedForProductLaunchEnumParam string

const (
	ProductcatalogproductsMarkedForProductLaunchEnumParamDefault   ProductcatalogproductsMarkedForProductLaunchEnumParam = "default"
	ProductcatalogproductsMarkedForProductLaunchEnumParamMarked    ProductcatalogproductsMarkedForProductLaunchEnumParam = "marked"
	ProductcatalogproductsMarkedForProductLaunchEnumParamNotMarked ProductcatalogproductsMarkedForProductLaunchEnumParam = "not_marked"
)

func (value ProductcatalogproductsMarkedForProductLaunchEnumParam) String() string {
	return string(value)
}
