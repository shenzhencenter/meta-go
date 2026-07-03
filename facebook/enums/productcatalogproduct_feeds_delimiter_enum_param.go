package enums

type ProductcatalogproductFeedsDelimiterEnumParam string

const (
	ProductcatalogproductFeedsDelimiterEnumParamAutodetect ProductcatalogproductFeedsDelimiterEnumParam = "AUTODETECT"
	ProductcatalogproductFeedsDelimiterEnumParamBar        ProductcatalogproductFeedsDelimiterEnumParam = "BAR"
	ProductcatalogproductFeedsDelimiterEnumParamComma      ProductcatalogproductFeedsDelimiterEnumParam = "COMMA"
	ProductcatalogproductFeedsDelimiterEnumParamSemicolon  ProductcatalogproductFeedsDelimiterEnumParam = "SEMICOLON"
	ProductcatalogproductFeedsDelimiterEnumParamTab        ProductcatalogproductFeedsDelimiterEnumParam = "TAB"
	ProductcatalogproductFeedsDelimiterEnumParamTilde      ProductcatalogproductFeedsDelimiterEnumParam = "TILDE"
)

func (value ProductcatalogproductFeedsDelimiterEnumParam) String() string {
	return string(value)
}
