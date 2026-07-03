package enums

type HomelistinggetHeatingType string

const (
	HomelistinggetHeatingTypeCentral    HomelistinggetHeatingType = "CENTRAL"
	HomelistinggetHeatingTypeElectric   HomelistinggetHeatingType = "ELECTRIC"
	HomelistinggetHeatingTypeEmptyValue HomelistinggetHeatingType = "EMPTY_VALUE"
	HomelistinggetHeatingTypeGas        HomelistinggetHeatingType = "GAS"
	HomelistinggetHeatingTypeNone       HomelistinggetHeatingType = "NONE"
	HomelistinggetHeatingTypeOther      HomelistinggetHeatingType = "OTHER"
	HomelistinggetHeatingTypeRadiator   HomelistinggetHeatingType = "RADIATOR"
)

func (value HomelistinggetHeatingType) String() string {
	return string(value)
}
