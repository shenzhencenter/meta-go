package enums

type ProductgroupproductsGenderEnumParam string

const (
	ProductgroupproductsGenderEnumParamFemale ProductgroupproductsGenderEnumParam = "female"
	ProductgroupproductsGenderEnumParamMale   ProductgroupproductsGenderEnumParam = "male"
	ProductgroupproductsGenderEnumParamUnisex ProductgroupproductsGenderEnumParam = "unisex"
)

func (value ProductgroupproductsGenderEnumParam) String() string {
	return string(value)
}
