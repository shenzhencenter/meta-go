package enums

type VehicleAvailability string

const (
	VehicleAvailabilityAvailable    VehicleAvailability = "AVAILABLE"
	VehicleAvailabilityNotAvailable VehicleAvailability = "NOT_AVAILABLE"
	VehicleAvailabilityPending      VehicleAvailability = "PENDING"
	VehicleAvailabilityUnknown      VehicleAvailability = "UNKNOWN"
)

func (value VehicleAvailability) String() string {
	return string(value)
}
