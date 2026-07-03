package enums

type ProductcatalogagenciesPermittedRolesEnumParam string

const (
	ProductcatalogagenciesPermittedRolesEnumParamAdmin      ProductcatalogagenciesPermittedRolesEnumParam = "ADMIN"
	ProductcatalogagenciesPermittedRolesEnumParamAdvertiser ProductcatalogagenciesPermittedRolesEnumParam = "ADVERTISER"
)

func (value ProductcatalogagenciesPermittedRolesEnumParam) String() string {
	return string(value)
}
