package enums

type ProductcatalogproductsGenderEnumParam string

const (
	ProductcatalogproductsGenderEnumParamFemale ProductcatalogproductsGenderEnumParam = "female"
	ProductcatalogproductsGenderEnumParamMale   ProductcatalogproductsGenderEnumParam = "male"
	ProductcatalogproductsGenderEnumParamUnisex ProductcatalogproductsGenderEnumParam = "unisex"
)

func (value ProductcatalogproductsGenderEnumParam) String() string {
	return string(value)
}
