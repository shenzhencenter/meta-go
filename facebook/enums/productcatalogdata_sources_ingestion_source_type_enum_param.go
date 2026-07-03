package enums

type ProductcatalogdataSourcesIngestionSourceTypeEnumParam string

const (
	ProductcatalogdataSourcesIngestionSourceTypeEnumParamAll           ProductcatalogdataSourcesIngestionSourceTypeEnumParam = "ALL"
	ProductcatalogdataSourcesIngestionSourceTypeEnumParamPrimary       ProductcatalogdataSourcesIngestionSourceTypeEnumParam = "PRIMARY"
	ProductcatalogdataSourcesIngestionSourceTypeEnumParamSupplementary ProductcatalogdataSourcesIngestionSourceTypeEnumParam = "SUPPLEMENTARY"
)

func (value ProductcatalogdataSourcesIngestionSourceTypeEnumParam) String() string {
	return string(value)
}
