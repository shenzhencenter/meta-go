package enums

type ProductitemAgeGroup string

const (
	ProductitemAgeGroupAdult   ProductitemAgeGroup = "adult"
	ProductitemAgeGroupAllAges ProductitemAgeGroup = "all ages"
	ProductitemAgeGroupInfant  ProductitemAgeGroup = "infant"
	ProductitemAgeGroupKids    ProductitemAgeGroup = "kids"
	ProductitemAgeGroupNewborn ProductitemAgeGroup = "newborn"
	ProductitemAgeGroupTeen    ProductitemAgeGroup = "teen"
	ProductitemAgeGroupToddler ProductitemAgeGroup = "toddler"
)

func (value ProductitemAgeGroup) String() string {
	return string(value)
}
