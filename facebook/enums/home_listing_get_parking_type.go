package enums

type HomelistinggetParkingType string

const (
	HomelistinggetParkingTypeEmptyValue HomelistinggetParkingType = "EMPTY_VALUE"
	HomelistinggetParkingTypeGarage     HomelistinggetParkingType = "GARAGE"
	HomelistinggetParkingTypeNone       HomelistinggetParkingType = "NONE"
	HomelistinggetParkingTypeOffStreet  HomelistinggetParkingType = "OFF_STREET"
	HomelistinggetParkingTypeOther      HomelistinggetParkingType = "OTHER"
	HomelistinggetParkingTypeStreet     HomelistinggetParkingType = "STREET"
)

func (value HomelistinggetParkingType) String() string {
	return string(value)
}
