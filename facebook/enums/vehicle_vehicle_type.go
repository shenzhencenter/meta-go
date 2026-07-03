package enums

type VehicleVehicleType string

const (
	VehicleVehicleTypeBoat       VehicleVehicleType = "BOAT"
	VehicleVehicleTypeCarTruck   VehicleVehicleType = "CAR_TRUCK"
	VehicleVehicleTypeCommercial VehicleVehicleType = "COMMERCIAL"
	VehicleVehicleTypeMotorcycle VehicleVehicleType = "MOTORCYCLE"
	VehicleVehicleTypeOther      VehicleVehicleType = "OTHER"
	VehicleVehicleTypePowersport VehicleVehicleType = "POWERSPORT"
	VehicleVehicleTypeRvCamper   VehicleVehicleType = "RV_CAMPER"
	VehicleVehicleTypeTrailer    VehicleVehicleType = "TRAILER"
)

func (value VehicleVehicleType) String() string {
	return string(value)
}
