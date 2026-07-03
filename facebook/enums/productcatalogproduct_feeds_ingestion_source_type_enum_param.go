package enums

type ProductcatalogproductFeedsIngestionSourceTypeEnumParam string

const (
	ProductcatalogproductFeedsIngestionSourceTypeEnumParamPrimaryFeed       ProductcatalogproductFeedsIngestionSourceTypeEnumParam = "PRIMARY_FEED"
	ProductcatalogproductFeedsIngestionSourceTypeEnumParamSupplementaryFeed ProductcatalogproductFeedsIngestionSourceTypeEnumParam = "SUPPLEMENTARY_FEED"
)

func (value ProductcatalogproductFeedsIngestionSourceTypeEnumParam) String() string {
	return string(value)
}
