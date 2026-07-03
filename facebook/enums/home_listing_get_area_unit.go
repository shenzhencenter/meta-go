package enums

type HomelistinggetAreaUnit string

const (
	HomelistinggetAreaUnitEmptyValue HomelistinggetAreaUnit = "EMPTY_VALUE"
	HomelistinggetAreaUnitSqft       HomelistinggetAreaUnit = "SQFT"
	HomelistinggetAreaUnitSqm        HomelistinggetAreaUnit = "SQM"
)

func (value HomelistinggetAreaUnit) String() string {
	return string(value)
}
