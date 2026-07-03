package enums

type HomelistinggetFurnishType string

const (
	HomelistinggetFurnishTypeEmptyValue    HomelistinggetFurnishType = "EMPTY_VALUE"
	HomelistinggetFurnishTypeFurnished     HomelistinggetFurnishType = "FURNISHED"
	HomelistinggetFurnishTypeSemifurnished HomelistinggetFurnishType = "SEMIFURNISHED"
	HomelistinggetFurnishTypeUnfurnished   HomelistinggetFurnishType = "UNFURNISHED"
)

func (value HomelistinggetFurnishType) String() string {
	return string(value)
}
