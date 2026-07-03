package enums

type ProductcatalogproductsAgeGroupEnumParam string

const (
	ProductcatalogproductsAgeGroupEnumParamAdult   ProductcatalogproductsAgeGroupEnumParam = "adult"
	ProductcatalogproductsAgeGroupEnumParamAllAges ProductcatalogproductsAgeGroupEnumParam = "all ages"
	ProductcatalogproductsAgeGroupEnumParamInfant  ProductcatalogproductsAgeGroupEnumParam = "infant"
	ProductcatalogproductsAgeGroupEnumParamKids    ProductcatalogproductsAgeGroupEnumParam = "kids"
	ProductcatalogproductsAgeGroupEnumParamNewborn ProductcatalogproductsAgeGroupEnumParam = "newborn"
	ProductcatalogproductsAgeGroupEnumParamTeen    ProductcatalogproductsAgeGroupEnumParam = "teen"
	ProductcatalogproductsAgeGroupEnumParamToddler ProductcatalogproductsAgeGroupEnumParam = "toddler"
)

func (value ProductcatalogproductsAgeGroupEnumParam) String() string {
	return string(value)
}
