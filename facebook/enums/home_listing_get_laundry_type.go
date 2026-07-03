package enums

type HomelistinggetLaundryType string

const (
	HomelistinggetLaundryTypeEmptyValue HomelistinggetLaundryType = "EMPTY_VALUE"
	HomelistinggetLaundryTypeInBuilding HomelistinggetLaundryType = "IN_BUILDING"
	HomelistinggetLaundryTypeInUnit     HomelistinggetLaundryType = "IN_UNIT"
	HomelistinggetLaundryTypeNone       HomelistinggetLaundryType = "NONE"
	HomelistinggetLaundryTypeOther      HomelistinggetLaundryType = "OTHER"
)

func (value HomelistinggetLaundryType) String() string {
	return string(value)
}
