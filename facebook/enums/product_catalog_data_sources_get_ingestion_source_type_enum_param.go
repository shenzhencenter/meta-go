package enums

type ProductcatalogdatasourcesgetIngestionSourceTypeEnumParam string

const (
	ProductcatalogdatasourcesgetIngestionSourceTypeEnumParamAll           ProductcatalogdatasourcesgetIngestionSourceTypeEnumParam = "ALL"
	ProductcatalogdatasourcesgetIngestionSourceTypeEnumParamPrimary       ProductcatalogdatasourcesgetIngestionSourceTypeEnumParam = "PRIMARY"
	ProductcatalogdatasourcesgetIngestionSourceTypeEnumParamSupplementary ProductcatalogdatasourcesgetIngestionSourceTypeEnumParam = "SUPPLEMENTARY"
)

func (value ProductcatalogdatasourcesgetIngestionSourceTypeEnumParam) String() string {
	return string(value)
}
