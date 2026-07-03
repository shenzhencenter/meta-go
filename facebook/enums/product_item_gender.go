package enums

type ProductitemGenderEnum string

const (
	ProductitemGenderEnumFemale ProductitemGenderEnum = "female"
	ProductitemGenderEnumMale   ProductitemGenderEnum = "male"
	ProductitemGenderEnumUnisex ProductitemGenderEnum = "unisex"
)

func (value ProductitemGenderEnum) String() string {
	return string(value)
}
