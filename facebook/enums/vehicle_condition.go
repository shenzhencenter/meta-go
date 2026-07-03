package enums

type VehicleCondition string

const (
	VehicleConditionExcellent VehicleCondition = "EXCELLENT"
	VehicleConditionFair      VehicleCondition = "FAIR"
	VehicleConditionGood      VehicleCondition = "GOOD"
	VehicleConditionNone      VehicleCondition = "NONE"
	VehicleConditionOther     VehicleCondition = "OTHER"
	VehicleConditionPoor      VehicleCondition = "POOR"
	VehicleConditionVeryGood  VehicleCondition = "VERY_GOOD"
)

func (value VehicleCondition) String() string {
	return string(value)
}
