package enums

type ProductitemGender string

const (
	ProductitemGenderFemale ProductitemGender = "female"
	ProductitemGenderMale   ProductitemGender = "male"
	ProductitemGenderUnisex ProductitemGender = "unisex"
)

func (value ProductitemGender) String() string {
	return string(value)
}
