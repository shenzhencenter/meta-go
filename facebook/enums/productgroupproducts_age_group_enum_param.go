package enums

type ProductgroupproductsAgeGroupEnumParam string

const (
	ProductgroupproductsAgeGroupEnumParamAdult   ProductgroupproductsAgeGroupEnumParam = "adult"
	ProductgroupproductsAgeGroupEnumParamAllAges ProductgroupproductsAgeGroupEnumParam = "all ages"
	ProductgroupproductsAgeGroupEnumParamInfant  ProductgroupproductsAgeGroupEnumParam = "infant"
	ProductgroupproductsAgeGroupEnumParamKids    ProductgroupproductsAgeGroupEnumParam = "kids"
	ProductgroupproductsAgeGroupEnumParamNewborn ProductgroupproductsAgeGroupEnumParam = "newborn"
	ProductgroupproductsAgeGroupEnumParamTeen    ProductgroupproductsAgeGroupEnumParam = "teen"
	ProductgroupproductsAgeGroupEnumParamToddler ProductgroupproductsAgeGroupEnumParam = "toddler"
)

func (value ProductgroupproductsAgeGroupEnumParam) String() string {
	return string(value)
}
