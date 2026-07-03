package enums

type ProductcatalogproductFeedsEncodingEnumParam string

const (
	ProductcatalogproductFeedsEncodingEnumParamAutodetect ProductcatalogproductFeedsEncodingEnumParam = "AUTODETECT"
	ProductcatalogproductFeedsEncodingEnumParamLatin1     ProductcatalogproductFeedsEncodingEnumParam = "LATIN1"
	ProductcatalogproductFeedsEncodingEnumParamUtf16be    ProductcatalogproductFeedsEncodingEnumParam = "UTF16BE"
	ProductcatalogproductFeedsEncodingEnumParamUtf16le    ProductcatalogproductFeedsEncodingEnumParam = "UTF16LE"
	ProductcatalogproductFeedsEncodingEnumParamUtf32be    ProductcatalogproductFeedsEncodingEnumParam = "UTF32BE"
	ProductcatalogproductFeedsEncodingEnumParamUtf32le    ProductcatalogproductFeedsEncodingEnumParam = "UTF32LE"
	ProductcatalogproductFeedsEncodingEnumParamUTF8       ProductcatalogproductFeedsEncodingEnumParam = "UTF8"
)

func (value ProductcatalogproductFeedsEncodingEnumParam) String() string {
	return string(value)
}
