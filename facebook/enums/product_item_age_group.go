package enums

type ProductitemAgeGroupEnum string

const (
	ProductitemAgeGroupEnumAdult   ProductitemAgeGroupEnum = "adult"
	ProductitemAgeGroupEnumAllAges ProductitemAgeGroupEnum = "all ages"
	ProductitemAgeGroupEnumInfant  ProductitemAgeGroupEnum = "infant"
	ProductitemAgeGroupEnumKids    ProductitemAgeGroupEnum = "kids"
	ProductitemAgeGroupEnumNewborn ProductitemAgeGroupEnum = "newborn"
	ProductitemAgeGroupEnumTeen    ProductitemAgeGroupEnum = "teen"
	ProductitemAgeGroupEnumToddler ProductitemAgeGroupEnum = "toddler"
)

func (value ProductitemAgeGroupEnum) String() string {
	return string(value)
}
