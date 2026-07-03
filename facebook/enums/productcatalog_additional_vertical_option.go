package enums

type ProductcatalogAdditionalVerticalOption string

const (
	ProductcatalogAdditionalVerticalOptionLocalDaCatalog ProductcatalogAdditionalVerticalOption = "LOCAL_DA_CATALOG"
	ProductcatalogAdditionalVerticalOptionLocalProducts  ProductcatalogAdditionalVerticalOption = "LOCAL_PRODUCTS"
)

func (value ProductcatalogAdditionalVerticalOption) String() string {
	return string(value)
}
