package enums

type VehicleStateOfVehicle string

const (
	VehicleStateOfVehicleCpo  VehicleStateOfVehicle = "CPO"
	VehicleStateOfVehicleNew  VehicleStateOfVehicle = "NEW"
	VehicleStateOfVehicleUsed VehicleStateOfVehicle = "USED"
)

func (value VehicleStateOfVehicle) String() string {
	return string(value)
}
