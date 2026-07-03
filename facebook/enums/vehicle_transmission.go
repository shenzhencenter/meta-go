package enums

type VehicleTransmission string

const (
	VehicleTransmissionAutomatic VehicleTransmission = "AUTOMATIC"
	VehicleTransmissionManual    VehicleTransmission = "MANUAL"
	VehicleTransmissionNone      VehicleTransmission = "NONE"
	VehicleTransmissionOther     VehicleTransmission = "OTHER"
)

func (value VehicleTransmission) String() string {
	return string(value)
}
